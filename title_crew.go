package main

import (
	"github.com/lib/pq"
	"strings"
)

func titleCrew(columns []string) ([]interface{}, error) {
	var directors, writers []string
	directors = strings.Split(columns[1], ",")
	writers = strings.Split(columns[2], ",")

	return []interface{}{
		columns[0],                // tconst
		pq.StringArray(directors), // columns[1]
		pq.StringArray(writers),   // columns[2]
	}, nil
}
