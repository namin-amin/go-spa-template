package persistance

import (
	"log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

type DataBase struct {
	db *dbx.DB
}

func (db *DataBase)  GetDb() *dbx.DB  {
	return db.db
}

func NewDb(con string) *DataBase {
	idb, err := dbx.Open("postres",con)

	if err != nil {
		log.Fatalln("could not connect yo db")
	}
	return &DataBase{
		db: idb,
	 }
}