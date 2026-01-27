package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	debug     = flag.Bool("debug", false, "Enable debugging output")
	dryrun    = flag.Bool("dryrun", false, "Run in dry mode, with no db commits")
	database  = flag.String("db", "cad:cad@/cad", "MySQL CAD backup database")
	backupdir = flag.String("backupdir", "backup", "Read from backup directory")
)

func main() {
	flag.Parse()

	// db initialization
	var l logger.Interface
	l = logger.Default.LogMode(logger.Warn)
	if *debug {
		l = logger.Default
	}

	var db *gorm.DB
	var err error

	if !*dryrun {
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: *database,
		}), &gorm.Config{
			DefaultTransactionTimeout: time.Second * 10,
			Logger:                    l,
		})
		if err != nil {
			log.Printf("ERR: gorm.Open: %s", err.Error())
			panic(err)
		}
		err = db.AutoMigrate(
			&IncidentObj{},
			&NarrativeObj{},
			&UnitObj{},
			&UnitLogObj{},
			&CallLogObj{},
			&CallObj{},
		)
		if err != nil {
			log.Printf("ERR: db.AutoMigrate: %s", err.Error())
			panic(err)
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

		var status CADCall
		err = json.Unmarshal(contents, &status)
		if err != nil {
			log.Printf("ERROR: GetStatus(): %s", err.Error())
			continue
		}
		/*
			for k := range status.Incidents {
				status.Incidents[k].CallID = int64(status.Call.CallID)
			}

			for k := range status.Units {
				status.Units[k].CallID = int64(status.Call.CallID)
			}

			for k := range status.Narratives {
				status.Narratives[k].CallID = int64(status.Call.CallID)
			}

			for k := range status.Logs {
				status.Logs[k].CallID = int64(status.Call.CallID)
			}

			for k := range status.UnitLogs {
				status.UnitLogs[k].CallID = int64(status.Call.CallID)
			}
		*/
		if *debug {
			log.Printf("DEBUG: CADCall : %#v", status)
		}

		if !*dryrun {
			log.Printf("Importing call record for %s", status.Call.IncidentNumber)
			tx := db.Create(&(status.Call))
			if tx.Error != nil {
				log.Printf("ERROR: %s", tx.Error)
			}
			{
				log.Printf("Importing %d incident records", len(status.Incidents))
				for _, v := range status.Incidents {
					tx := db.Create(&v)
					if tx.Error != nil {
						log.Printf("ERROR: %s", tx.Error)
					}
				}
			}
			{
				log.Printf("Importing %d narratives records", len(status.Narratives))
				for _, v := range status.Narratives {
					tx := db.Create(&v)
					if tx.Error != nil {
						log.Printf("ERROR: %s", tx.Error)
					}
				}
			}
			{
				log.Printf("Importing %d unit records", len(status.Units))
				for _, v := range status.Units {
					tx := db.Create(&v)
					if tx.Error != nil {
						log.Printf("ERROR: %s", tx.Error)
					}
				}
			}
			{
				log.Printf("Importing %d unit log records", len(status.UnitLogs))
				for _, v := range status.UnitLogs {
					tx := db.Create(&v)
					if tx.Error != nil {
						log.Printf("ERROR: %s", tx.Error)
					}
				}
			}
			{
				log.Printf("Importing %d log records", len(status.Logs))
				for _, v := range status.Logs {
					tx := db.Create(&v)
					if tx.Error != nil {
						log.Printf("ERROR: %s", tx.Error)
					}
				}
			}

			/*
				tx := db.Create(&status)
				if tx.Error != nil {
					log.Printf("ERROR: %s", tx.Error)
				}
			*/
		}

		//return // DEBUG: TODO: FIXME: XXX
	}
}
