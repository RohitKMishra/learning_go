package notifier

import (
	"workspace/pusher/running/running-results-table/internal/db"

	"github.com/pusher/pusher-http-go"
)

type Notifier struct {
	notifyChannel chan<- bool
}

func notifier(database *db.Database, notifyChannel <-chan bool) {
	client := pusher.Client{
		AppID:   "902466",
		Key:     "843bfb1afcfc81ff0586",
		Secret:  "cc9fde8f246ef6b74f84",
		Cluster: "ap2",
		Secure:  true,
	}

	for {
		<-notifyChannel
		data := map[string][]db.Record{"results": database.GetRecords()}
		client.Trigger("results", "results", data)
	}
}

func New(database *db.Database) Notifier {
	notifyChannel := make(chan bool)
	go notifier(database, notifyChannel)
	return Notifier{
		notifyChannel,
	}
}

func (notifier *Notifier) Notify() {
	notifier.notifyChannel <- true
}
