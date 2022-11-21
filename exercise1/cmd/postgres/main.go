package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://localhost:5433/imdbload")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	matches, err := filepath.Glob("modules/join-order-benchmark/*.sql")

	if err != nil {
		log.Fatal(err)
	}

	job := make([]string, 0)

	r := regexp.MustCompile(`[^\d\n]*(\d+)([a-z])\.sql`)

	for _, match := range matches {
		if r.MatchString(match) {
			job = append(job, match)
		}
	}

	sort.Slice(job, func(i, j int) bool {
		mi := r.FindStringSubmatch(job[i])
		mj := r.FindStringSubmatch(job[j])

		ii, _ := strconv.Atoi(mi[1])
		ij, _ := strconv.Atoi(mj[1])

		si, sj := mi[2], mj[2]

		if ii == ij {
			return si < sj
		}

		return ii < ij
	})

	for _, name := range job {
		query, err := os.ReadFile(name)

		if err != nil {
			log.Fatal(err)
		}

		{
			start := time.Now()

			rows, err := db.Query(context.Background(), string(query))

			elapsed := time.Since(start)

			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				fmt.Println(name, elapsed)
			}
		}
	}
}
