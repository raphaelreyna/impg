package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"os"
	"sync"
)

type extractor func([]string) ([]interface{}, error)

var CreateTablesSQL = ""

func main() {
	retVal := 1
	defer func() {
		os.Exit(retVal)
	}()

	if CreateTablesSQL == "" {
		fmt.Printf("%s was not properly built: missing SQL for creating tables\n", os.Args[0])
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("no connection string given")
		return
	}

	connector, err := pq.NewConnector(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	db := sql.OpenDB(connector)
	defer db.Close()

	//Make sure tables exist and are empty
	_, err = db.Exec(CreateTablesSQL)
	if err != nil {
		fmt.Printf("error creating tables: %s\n", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)
	// name_basics
	go func() {
		defer wg.Done()
		err = uploadTable(db, "name_basics", "name.basics.tsv", nameBasics)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	// title_akas
	go func() {
		defer wg.Done()
		err = uploadTable(db, "title_akas", "title.akas.tsv", titleAkas)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	// title_basics
	go func() {
		defer wg.Done()
		err = uploadTable(db, "title_basics", "title.basics.tsv", titleBasics)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	// title_crew
	// title_episode
	// title_principals
	// title_ratings

	wg.Wait()

	retVal = 0
}
