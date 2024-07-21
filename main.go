package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const DEPLOYS_FILE = "deploys.json"

type Deploy struct {
	ID          uuid.UUID
	Application string
	Environment string
	Team        string
	Squad       string
	Actor       string
	Commit      string
	Release     string
	Url         string
	Urgency     string
	Priority    string
	Category    string
	Description string
	Status      string
	State       string
	StartedAt   time.Time
	FinishedAt  time.Time
	Duration    string
}

type DeployFile struct {
	Deploys []*Deploy
}

func NewDeploy(
	application string,
	environment string,
	team string,
	squad string,
	actor string,
	commit string,
	release string,
	url string,
	urgency string,
	priority string,
	category string,
	description string,
	status string,
	state string,
	started_at time.Time,
	finished_at time.Time,
) *Deploy {
	return &Deploy{
		ID:          uuid.New(),
		Application: application,
		Environment: environment,
		Team:        team,
		Squad:       squad,
		Actor:       actor,
		Commit:      commit,
		Release:     release,
		Url:         url,
		Urgency:     urgency,
		Priority:    priority,
		Category:    category,
		Description: description,
		Status:      status,
		State:       state,
		StartedAt:   started_at,
		FinishedAt:  finished_at,
		Duration:    finished_at.Sub(started_at).String(),
	}
}

type DeployRepository interface {
	Create() error
	GetById(id uuid.UUID) (*Deploy, error)
	List() ([]*Deploy, error)
	Update(id uuid.UUID) (*Deploy, error)
	Delete(id uuid.UUID) error
}

func (deploy *Deploy) Create() error {
	_, err := os.Stat(DEPLOYS_FILE)
	var deploy_list []*Deploy

	if err == nil {
		file_content, err := os.ReadFile(DEPLOYS_FILE)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
		}

		if err := json.Unmarshal(file_content, &deploy_list); err != nil {
			log.Fatalf("Fail to decode file: %v", err)
		}
	}

	deploy_list = append(deploy_list, deploy)
	deploy_file_content, err := json.MarshalIndent(&deploy_list, "", "  ")

	if err != nil {
		log.Fatalf("Fail to encode: %v", err)
		return err
	}
	if err := os.WriteFile(DEPLOYS_FILE, deploy_file_content, 0644); err != nil {
		log.Fatalf("Fail to write deploy_file: %v", err)
		return err
	}

	return nil
}

func main() {
	deploy := NewDeploy(
		"app-test",
		"dev",
		"devops",
		"devops",
		"contact@aleroxac.io",
		"01234567abcdef",
		"v1.0.100",
		"https://jenkins.aleroxac.io/app-test/#100",
		"low",
		"low",
		"feature",
		"app-test deploy 100",
		"finished",
		"success",
		time.Now(),
		time.Now().Add(120*time.Second),
	)

	err := deploy.Create()
	if err != nil {
		log.Fatalf("Fail to create deploy: %v", err)
		return
	}
}
