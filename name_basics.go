package main

import (
	"github.com/lib/pq"
	"strconv"
	"strings"
)

func nameBasics(columns []string) ([]interface{}, error) {
	var birthYear, deathYear interface{}
	birthYear, err := strconv.ParseInt(columns[2], 10, 32)
	if err != nil {
		// Check for null value
		if columns[2] != `\N` {
			return nil, err
		}
	}
	deathYear, err = strconv.ParseInt(columns[3], 10, 32)
	if err != nil {
		// Check for null value
		if columns[3] != `\N` {
			return nil, err
		}
	}

	var primaryProfession, knownForTitles []string
	primaryProfession = strings.Split(columns[4], ",")
	knownForTitles = strings.Split(columns[5], ",")

	return []interface{}{
		columns[0],                        // nconst
		columns[1],                        // primary_name
		birthYear,                         // columns[2]
		deathYear,                         // columns[3]
		pq.StringArray(primaryProfession), // columns[4]
		pq.StringArray(knownForTitles),    // columns[5]
	}, nil
}
