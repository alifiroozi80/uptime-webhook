package alert

import (
	"encoding/json"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	log "github.com/sirupsen/logrus"
)

// CreateSquadcastIncident creates an incident in squadcast
func (r *Repository) CreateSquadcastIncident(alert models.Alert) error {
	var urls []string

	for _, tag := range alert.Data.Service.Tags {
		url, ok := r.config.Notifier.Squadcast.Teams[strings.ToLower(tag)]
		if !ok {
			log.Errorf("[SQUADCAST] Team not found for tag: %s", tag)
			continue
		}

		urls = append(urls, url)
	}

	payload := formatSquadcastMessage(alert)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range urls {
		r.CallSquadcast(url, body)
	}

	return nil
}
