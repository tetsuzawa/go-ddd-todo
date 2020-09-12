package main

import (
	"flag"
	"log"
	"os"

	"github.com/tetsuzawa/godddtodo/interfaces/api/server"
)

func main() {
	var databaseDatasource string
	var port int
	flag.StringVar(&databaseDatasource, "databaseDatasource", "./data.db", "Should looks like ./data.db")
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s := server.NewServer()
	if err := s.Init(databaseDatasource); err != nil {
		log.Fatal(err)
	}
	s.Run(port)
}
