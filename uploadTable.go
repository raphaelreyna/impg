package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/lib/pq"
	"os"
	"strings"
)

func uploadTable(ctx context.Context, db *sql.DB, tableName, fileName string, extract extractor) error {
	er := func(e error) error {
		return fmt.Errorf("%s error: %s", tableName, e.Error())
	}

	// Open tsv file
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return er(err)
	}

	scanner := bufio.NewScanner(file)
	var line string
	first := true
	//counter := 0

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return er(err)
	}
	var stmt *sql.Stmt

	for scanner.Scan() {
		line = scanner.Text()
		columns := strings.Split(line, "\t")
		if first {
			// Column names need to be in snake_case
			cols := []string{}
			for _, column := range columns {
				cols = append(cols, strcase.ToSnake(column))
			}

			// Create the statement used to upload a row
			stmtString := pq.CopyIn(tableName, cols...)
			stmt, err = tx.PrepareContext(ctx, stmtString)
			if err != nil {
				return er(err)
			}

			first = false
			continue
		}

		// Extract data for row using correct data types
		vals, err := extract(columns)
		if err != nil {
			return er(err)
		}

		_, err = stmt.ExecContext(ctx, vals...)
		if err != nil {
			return er(err)
		}

		/*
			if counter > 500 {
				break
			}
			counter += 1
		*/
	}

	if _, err = stmt.ExecContext(ctx); err != nil {
		return er(err)
	}

	if err = stmt.Close(); err != nil {
		return er(err)
	}

	return tx.Commit()
}
