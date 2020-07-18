package main

import (
	"github.com/lib/pq"
	"strconv"
	"strings"
)

func titleAkas(columns []string) ([]interface{}, error) {
	var err error
	var ordering, isOriginalTitle interface{}
	ordering, err = strconv.ParseInt(columns[1], 10, 32)
	if err != nil {
		// Check for null value
		if columns[1] != `\N` {
			return nil, err
		}
	}
	iot, err := strconv.ParseInt(columns[7], 10, 32)
	if err != nil {
		// Check for null value
		if columns[7] != `\N` {
			return nil, err
		}
	}
	if iot == 0 {
		isOriginalTitle = false
	} else {
		isOriginalTitle = true
	}

	var types, attributes []string
	types = strings.Split(columns[5], ",")
	attributes = strings.Split(columns[6], ",")

	return []interface{}{
		columns[0],                 // title_id
		ordering,                   // columns[1]
		columns[2],                 // title
		columns[3],                 // region
		columns[4],                 // language
		pq.StringArray(types),      // columns[5]
		pq.StringArray(attributes), // columns[6]
		isOriginalTitle,            // columns[7]
	}, nil
}
