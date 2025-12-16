package service

import (
	"encoding/json"
	"fmt"
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
			// writeUrl(link, state)
		}()
	}

	wg.Wait()

	return linksStates
}

func WriteJson() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("error reading config %v\n", err)
		return
	}
	storage_path := cfg.StoragePath

	inner_json, err := os.ReadFile(storage_path)
	if err != nil {
		log.Printf("created storage-file %s\n", storage_path)
		os.Create(storage_path)
	}

	var data map[string]interface{}
	err = json.Unmarshal(inner_json, &data)
	if err != nil {
		log.Printf("error unmarshalling %v\n", err)
		return
	}
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

	type Links struct {
		Requested_links [][]string
	}



	// файл пустой -> заполняем с нуля
	if len(data) == 2 || len(data) == 0 {
		// links := Links{
		// 	Requested_links: [][]string{{url, state}},
		// }
		links := map[string]interface
		j_file, err := json.MarshalIndent(links, "", " ")
		if err != nil {
			log.Printf("error parsing json %v\n", err)
			return
		}
		os.WriteFile(storage_path, j_file, 0644)
	} else {
		var j_file map[string]interface{}
		data, err = os.ReadFile(storage_path)
		if err != nil {
			log.Printf("error reading file %v\n", err)
			return
		}

		err = json.Unmarshal(data, &j_file)
		if err != nil {
			log.Printf("error unmarshalling file %v\n", err)
			return
		}

		fmt.Println(j_file["Requested_links"])
	}

	// link_to_add := Links{
	// 	requested_links: [][]string{
	// 		{url, state},
	// 	},
	// }
}
