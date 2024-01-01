package main


import (
	"flag"
	"os"
	"log"

	"github.com/xwxb/srss/platform"
	"github.com/xwxb/srss/storage"
)

var OptList = make([]string, 0)

func opt(envVar, defaultValue string) string {
	OptList = append(OptList, envVar)
	value := os.Getenv(envVar)
	if value != "" {
		return value
	}
	return defaultValue
}

func main() {

	var addr, dbPath string

	flag.StringVar(&addr, "addr", opt("YARR_ADDR", "127.0.0.1:7070"), "address to run server on")
	flag.StringVar(&dbPath, "db", opt("YARR_DB", ""), "storage file `path`")
	flag.Parse()


	store, err := storage.New(dbPath)
	if err != nil {
		log.Fatal("Failed to initialise database: ", err)
	}

	srv := server.NewServer(store, addr)


	platform.Start()
}
