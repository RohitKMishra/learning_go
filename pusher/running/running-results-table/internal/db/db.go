package db

type Record struct {
	Name string  `json:"name"`
	Time float32 `json:"time"`
}

func NewRecord(name string, time float32) Record {
	return Record{name, time}
}

type Database struct {
	contents []Record
}

func New() Database {
	contents := make([]Record, 0)
	return Database{contents}
}

func (database *Database) AddRecord(r Record) {
	database.contents = append(database.contents, r)
}

func (database *Database) GetRecords() []Record {
	return database.contents
}

// Here we are creating a new type called Record that represents the data that we store and a new struct called Database that rrepresents the actual database we areusing
// We then create some methods on the database type to add record and to get list of records.

// Next we  create our web server. For this we are going to create a new directory called internal/webapp under our work area, and a new file called webapp.go in this directory as follow.

// RUNNER LEADERBOARD
