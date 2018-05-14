package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tinee/doit/http"

	"github.com/Tinee/doit/mongo"
)

func main() {
	mclient := mongo.NewClient(mongo.ClientInfo{
		DBName: "doit",
		Addr:   "localhost",
	})
	defer mclient.Close()

	if err := mclient.Open(); err != nil {
		log.Fatalln(err)
	}

	s := http.NewServer(3000, mclient)

	defer s.Close()
	if err := s.Open(); err != nil {
		log.Fatalln(err)
	}

	csig := make(chan os.Signal)
	signal.Notify(csig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Termination Signal %v", <-csig)
}
