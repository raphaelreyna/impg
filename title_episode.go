package main

import (
	"strconv"
)

func titleEpisode(columns []string) ([]interface{}, error) {
	var err error
	var season_number, episode_number interface{}
	season_number, err = strconv.ParseInt(columns[2], 10, 32)
	if err != nil {
		// Check for null value
		if columns[2] != `\N` {
			return nil, err
		}
	}
	episode_number, err = strconv.ParseInt(columns[3], 10, 32)
	if err != nil {
		// Check for null value
		if columns[3] != `\N` {
			return nil, err
		}
	}

	return []interface{}{
		columns[0],     // tconst
		columns[1],     // parent_tconst
		season_number,  // columns[2]
		episode_number, // columns[3]
	}, nil
}
