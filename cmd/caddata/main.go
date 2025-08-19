package main

import (
	"flag"
)

var (
	action     = flag.String("action", "", "Action: ingest|stats")
	backupdir  = flag.String("backupdir", "backup", "Read from backup directory")
	filestore  = flag.String("filestore", "cad.json", "Store for JSON serialized data")
	add        = flag.Bool("add", false, "Add to filestore on ingest")
	ori        = flag.String("ori", "04042", "ORI/FDID")
	skipprefix = flag.String("skip-prefix", "STA,RES,A,FM", "Skip prefixes for unit analyses")
	reqsuffix  = flag.String("require-suffix", "", "Require certain suffix")
)

func main() {
	flag.Parse()

	switch *action {
	case "ingest":
		ingest()
	case "stats":
		stats()
	default:
		flag.PrintDefaults()
	}
}
