package main

import (
	"github.com/lib/pq"
	"strconv"
	"strings"
)

func titleBasics(columns []string) ([]interface{}, error) {
	var err error
	var is_adult interface{}
	ia, err := strconv.ParseInt(columns[4], 10, 32)
	if err != nil {
		// Check for null value
		if columns[4] != `\N` {
			return nil, err
		}
	}
	if ia == 0 {
		is_adult = false
	} else {
		is_adult = true
	}

	var start_year, end_year, runtime_minutes interface{}
	start_year, err = strconv.ParseInt(columns[5], 10, 32)
	if err != nil {
		// Check for null value
		if columns[5] != `\N` {
			return nil, err
		}
	}
	end_year, err = strconv.ParseInt(columns[6], 10, 32)
	if err != nil {
		// Check for null value
		if columns[6] != `\N` {
			return nil, err
		}
	}
	runtime_minutes, err = strconv.ParseInt(columns[7], 10, 32)
	if err != nil {
		// Check for null value
		if columns[7] != `\N` {
			return nil, err
		}
	}

	var genres []string
	genres = strings.Split(columns[8], ",")

	return []interface{}{
		columns[0],             // tconst
		columns[1],             // title_type
		columns[2],             // primary_title
		columns[3],             // original_title
		is_adult,               // columns[4]
		start_year,             // columns[5]
		end_year,               // columns[6]
		runtime_minutes,        // columns[7]
		pq.StringArray(genres), // columns[8]
	}, nil
}
