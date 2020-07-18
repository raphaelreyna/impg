package main

import (
	"strconv"
)

func titleRatings(columns []string) ([]interface{}, error) {
	var err error
	var average_rating interface{}
	average_rating, err = strconv.ParseFloat(columns[1], 32)
	if err != nil {
		// Check for null value
		if columns[2] != `\N` {
			return nil, err
		}
	}

	var num_votes int64
	num_votes, err = strconv.ParseInt(columns[2], 10, 32)
	if err != nil {
		// Check for null value
		if columns[3] != `\N` {
			return nil, err
		}
	}

	return []interface{}{
		columns[0],     // tconst
		average_rating, // columns[1]
		num_votes,      // columns[2]
	}, nil
}
