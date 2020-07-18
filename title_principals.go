package main

import (
	"github.com/lib/pq"
	"strconv"
	"strings"
)

func titlePrincipals(columns []string) ([]interface{}, error) {
	var err error

	var ordering interface{}
	ordering, err = strconv.ParseInt(columns[1], 10, 32)
	if err != nil {
		// Check for null value
		if columns[1] != `\N` {
			return nil, err
		}
	}

	var characters []string
	characters = strings.Split(columns[5], ",")

	return []interface{}{
		columns[0],                 // tconst
		ordering,                   // columns[1]
		columns[2],                 // nconst
		columns[3],                 // category
		columns[4],                 // job
		pq.StringArray(characters), // columns[5]
	}, nil
}
