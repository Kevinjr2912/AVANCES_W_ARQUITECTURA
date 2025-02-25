package application

import (
	"encoding/json"
	"github/pull_request_webhook/domain/value_objects"
	"log"
)

func ProcessPullRequestEvent(rawData []byte) int {
	var eventPayload value_objects.PullRequestEventPayload

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403
	}

	if eventPayload.Action == "closed" {
		log.Printf("Evento Pull Request recibido: Rama Base=%s, Rama Default='%s',User='%s', Repositorio='%s'", eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref, eventPayload.PullRequest.User.Name, eventPayload.Repository.FullName)
	}
	
	return 200
}

// Mostrar hacia dónde se hace el Pull Request
// Desde dónde se hace el pull request
// Usuario(username) que realizó el pull request
// Nombre del repositorio en donde se realizó