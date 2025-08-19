package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dayvillefire/newworld-cadview-agent/agent"
)

func readcalls() (calls []agent.CADCall, err error) {
	log.Printf("Importing %s", *filestore)
	contents, err := os.ReadFile(*filestore)
	if err != nil {
		log.Printf("ERROR: Reading file %s: %s", *filestore, err.Error())
		return
	}

	err = json.Unmarshal(contents, &calls)
	if err != nil {
		log.Printf("ERROR: GetStatus(): %s", err.Error())
		return
	}

	return
}
