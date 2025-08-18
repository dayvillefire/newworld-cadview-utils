package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/dayvillefire/newworld-cadview-agent/agent"
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
		}), &gorm.Config{Logger: l})
		if err != nil {
			log.Printf("ERR: gorm.Open: %s", err.Error())
			panic(err)
		}
		err = db.AutoMigrate(
			&agent.CADCall{},
			&agent.IncidentObj{},
			&agent.NarrativeObj{},
			&agent.UnitObj{},
			&agent.UnitLogObj{},
			&agent.CallLogObj{},
			&agent.CallObj{},
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

		var status agent.CADCall
		err = json.Unmarshal(contents, &status)
		if err != nil {
			log.Printf("ERROR: GetStatus(): %s", err.Error())
			continue
		}

		if *debug {
			log.Printf("DEBUG: CADCall : %#v", status)
		}

		if !*dryrun {
			tx := db.Create(&status)
			if tx.Error != nil {
				log.Printf("ERROR: %s", tx.Error)
			}
		}

		//return // DEBUG: TODO: FIXME: XXX
	}
}
