package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/lib/pq"
	"os"
	"os/signal"
	"sync"
)

// CreateTablesSQL is set while building impg
var CreateTablesSQL = ""

type tableJob struct {
	name      string
	file      string
	extractor extractor
}

var tableJobs = []*tableJob{
	&tableJob{
		name:      "name_basics",
		file:      "name.basics.tsv",
		extractor: nameBasics,
	},
	&tableJob{
		name:      "title_akas",
		file:      "title.akas.tsv",
		extractor: titleAkas,
	},
	&tableJob{
		name:      "title_basics",
		file:      "title.basics.tsv",
		extractor: titleBasics,
	},
	&tableJob{
		name:      "title_crew",
		file:      "title.crew.tsv",
		extractor: titleCrew,
	},
	&tableJob{
		name:      "title_episode",
		file:      "title.episode.tsv",
		extractor: titleEpisode,
	},
	&tableJob{
		name:      "title_principals",
		file:      "title.principals.tsv",
		extractor: titlePrincipals,
	},
	&tableJob{
		name:      "title_ratings",
		file:      "title.ratings.tsv",
		extractor: titleRatings,
	},
}

// extractor functions are used to convert each row in a file into an appropriate empty interface slice.
type extractor func([]string) ([]interface{}, error)

func main() {
	retVal := 1
	defer func() {
		os.Exit(retVal)
	}()

	// Make sure SQL was added during build
	if CreateTablesSQL == "" {
		fmt.Printf("%s was not properly built: missing SQL for creating tables\n", os.Args[0])
		return
	}

	// Setup flags
	var (
		host     string
		port     int
		user     string
		password string
		sslMode  bool
	)
	flag.StringVar(&host, "host", "localhost", `Postgres server host address`)
	flag.IntVar(&port, "port", 5432, `Postgres server port`)
	flag.StringVar(&user, "user", "postgres", `User to connect to the Postgres server as`)
	flag.StringVar(&password, "pass", "postgres", `Password to connect to the Postgres server`)
	flag.BoolVar(&sslMode, "ssl", false, `Use ssl-mode when connecting to the Postgres server`)
	flag.Parse()

	// Create Postgres connection string from flag values
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=",
		host, port, user, password)
	if sslMode {
		connStr += "enable"
	} else {
		connStr += "disable"
	}

	// Connect to database
	connector, err := pq.NewConnector(connStr)
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

	// Listen for os signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	ctx, cancelCtx := context.WithCancel(context.Background())
	go func() {
		<-sigChan
		cancelCtx()
	}()

	// Start jobs
	wg := &sync.WaitGroup{}
	for _, tt := range tableJobs {
		wg.Add(1)
		go func(t *tableJob) {
			defer wg.Done()
			err := uploadTable(ctx, db, t.name, t.file, t.extractor)
			if err != nil {
				cancelCtx()
				fmt.Println(err)
				return
			}
		}(tt)
	}

	wg.Wait()

	retVal = 0

	fmt.Println("Done")
}
