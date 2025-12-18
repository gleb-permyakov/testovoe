package service

import (
	"encoding/json"
	"log"
	"main/internal/client"
	"main/internal/config"
	"os"
	"sync"
)

type TaskService struct {
	Client *client.Client
}

func (t *TaskService) CheckURL(links []string) map[string]string {
	linksStates := make(map[string]string)

	wg := &sync.WaitGroup{}

	for _, link := range links {
		wg.Add(1)
		go func() {
			defer wg.Done()
			state := t.Client.Get(link)
			linksStates[link] = state
			WriteUrl(link, state)
		}()
	}

	wg.Wait()

	return linksStates
}

func WriteUrl(url, state string) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("error getting config %v\n", err)
		return
	}
	storage_path := cfg.StoragePath

	data, err := os.ReadFile(storage_path)
	if err != nil {
		log.Printf("created storage-file %s\n", storage_path)
		file, _ := os.Create(storage_path)
		file.Close()
	}

	type JsonFile struct {
		Requested_links [][]string
	}

	if len(data) == 2 || len(data) == 0 {
		var links JsonFile
		links.Requested_links = [][]string{{url, state}}
		j_file, err := json.MarshalIndent(links, "", " ")
		if err != nil {
			log.Printf("error parsing json %v\n", err)
			return
		}
		os.WriteFile(storage_path, j_file, 0644)
	} else {
		var links JsonFile
		j_file, err := os.ReadFile(storage_path)
		if err != nil {
			log.Printf("error reading file %v\n", err)
			return
		}
		err = json.Unmarshal(j_file, &links)
		if err != nil {
			log.Printf("error unmarshalling file %v\n", err)
			return
		}

		link := []string{url, state}
		links.Requested_links = append(links.Requested_links, link)

		j_file, err = json.MarshalIndent(links, "", " ")
		if err != nil {
			log.Printf("error parsing json %v\n", err)
			return
		}
		os.WriteFile(storage_path, j_file, 0644)
	}
}
