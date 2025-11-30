package service

import (
	"main/internal/client"
)

type TaskService struct {
	Client *client.Client
}

func (t *TaskService) CheckURL(links []string) map[string]string {
	linksStates := make(map[string]string)

	for _, link := range links {
		linksStates[link] = t.Client.Get(link)
	}

	return linksStates
}
