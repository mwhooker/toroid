package github

import (
	"encoding/json"
	"log"
)

type Repository struct {
	Name string
	Url  string
}

type Github struct {
	Repository Repository
	Ref        string
}

func ParsePayload(payload string) (gh *Github, err error) {
	if err = json.Unmarshal([]byte(payload), &gh); err != nil {
		log.Println(err)
	}
	log.Println(gh)
	return
}
