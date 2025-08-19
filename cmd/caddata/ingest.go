package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dayvillefire/newworld-cadview-agent/agent"
)

func ingest() {
	cadcalls := []agent.CADCall{}

	// Ingest first
	if *add {
		cc, err := readcalls()
		if err != nil {
			cadcalls = append(cadcalls, cc...)
		}
	}

	dirents, err := os.ReadDir(*backupdir)
	if err != nil {
		log.Printf("ERR: os.ReadDir: %s", err.Error())
		panic(err)
	}
	for _, dirent := range dirents {
		if dirent.IsDir() {
			continue
		}
		log.Printf("Processing %s", dirent.Name())
		fullPath := *backupdir + string(os.PathSeparator) + dirent.Name()
		contents, err := os.ReadFile(fullPath)
		if err != nil {
			log.Printf("ERROR: Reading file %s: %s", dirent.Name(), err.Error())
			continue
		}

		var status agent.CADCall
		err = json.Unmarshal(contents, &status)
		if err != nil {
			log.Printf("ERROR: GetStatus(): %s", err.Error())
			continue
		}

		cadcalls = append(cadcalls, status)

	}

	// Serialize

	var data []byte
	data, err = json.Marshal(cadcalls)
	if err != nil {
		log.Printf("ERR: %s", err.Error())
		return
	}

	log.Printf("INFO: Writing %d calls to %s (%d bytes)", len(cadcalls), *filestore, len(data))
	err = os.WriteFile(*filestore, data, 0644)
	if err != nil {
		log.Printf("ERR: %s", err.Error())
		return
	}
}
