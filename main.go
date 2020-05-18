package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jLemmings/goMKdb/models"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
)
import _ "github.com/lib/pq"

//go:generate sqlboiler --wipe psql

func main() {
	db, err := sql.Open("postgres", "dbname=postgres host=localhost user=postgres password=postgres sslmode=disable")

	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	fmt.Println("Connected to DB")


	u := models.User{Name: "test"}

	err = u.Insert(context.Background(), db, boil.Infer())

	fmt.Println("User ID:", u.ID)

	got, err := models.FindUser(context.Background(), db, u.ID)
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	fmt.Println("Found User ID", got.ID)
}