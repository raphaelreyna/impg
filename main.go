package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"os"
	"sync"
)

type extractor func([]string) ([]interface{}, error)

type table struct {
	name      string
	file      string
	extractor extractor
}

var CreateTablesSQL = ""

func upload(ctx context.Context, db *sql.DB, t *table, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()
	err := uploadTable(ctx, db, t.name, t.file, t.extractor)
	if err != nil {
		cancel()
		fmt.Println(err)
		return
	}
}

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
	ctx, cancelCtx := context.WithCancel(context.Background())

	tables := []*table{
		&table{
			name:      "name_basic",
			file:      "name.basic.tsv",
			extractor: nameBasics,
		},
		&table{
			name:      "title_akas",
			file:      "title.akas.tsv",
			extractor: titleAkas,
		},
		&table{
			name:      "title_basics",
			file:      "title.basics.tsv",
			extractor: titleBasics,
		},
		&table{
			name:      "title_crew",
			file:      "title.crew.tsv",
			extractor: titleCrew,
		},
		&table{
			name:      "title_episode",
			file:      "title.episode.tsv",
			extractor: titleEpisode,
		},
		&table{
			name:      "title_principals",
			file:      "title.principals.tsv",
			extractor: titlePrincipals,
		},
		&table{
			name:      "title_ratings",
			file:      "title.ratings.tsv",
			extractor: titleRatings,
		},
	}

	for _, t := range tables {
		wg.Add(1)
		go upload(ctx, db, t, wg, cancelCtx)
	}

	wg.Wait()

	retVal = 0
}
