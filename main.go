package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/karuppiah7890/pg-query-killer/pkg/config"
	"github.com/karuppiah7890/pg-query-killer/pkg/postgres"
)

func checkSignal(signals chan os.Signal, done chan bool) {
	<-signals
	done <- true
	os.Exit(0)
}

func main() {
	done := make(chan bool, 1)
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt)

	go checkSignal(signals, done)

	c, err := config.NewConfigFromEnvVars()
	if err != nil {
		log.Fatalf("error occurred while getting configuration from environment variables: %v", err)
	}

	client := postgres.NewClient(c.GetPostgresUri())

	// Connect to PostgreSQL

	// Get List of long running queries - say which are running for say 1 minute or 2 minutes long or 5 minutes long

	longRunningQueries, err := client.GetListOfLongRunningQueries()

	if err != nil {
		log.Fatalf("error occurred while getting long running queries: %v", err)
	}

	fmt.Printf("%+v", longRunningQueries)

	// Log the query and kill the queries that don't have wait event or wait type

	// have dry run to just print queries and not kill them. use flags? hmm. Like --dry-run

	// Ensure that the log file is log rotated
}
