package main

import (
	"workspace/pusher/running/running-results-table/internal/db"
	"workspace/pusher/running/running-results-table/internal/notifier"
	"workspace/pusher/running/running-results-table/internal/webapp"
)

func main() {
	database := db.New()
	notifierClient := notifier.New(&database)
	webapp.StartServer(&database, &notifierClient)
}
