package main

import (
	"2022ws-dbs-ex1/job"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
	"time"
)

var keyword dataframe.DataFrame
var movieInfo dataframe.DataFrame
var movieKeyword dataframe.DataFrame
var title dataframe.DataFrame

func main() {
	now := time.Now()

	job.Job3a(keyword, movieInfo, movieKeyword, title)

	fmt.Printf("3a: %s\n\n", time.Since(now))
	now = time.Now()

	job.Job3b(keyword, movieInfo, movieKeyword, title)

	fmt.Printf("3b: %s\n\n", time.Since(now))
	now = time.Now()

	job.Job3c(keyword, movieInfo, movieKeyword, title)

	fmt.Printf("3c: %s\n\n", time.Since(now))
}

func init() {
	keyword = loadCSV("data/csv/keyword.csv")
	movieInfo = loadCSV("data/csv/movie_info.csv")
	movieKeyword = loadCSV("data/csv/movie_keyword.csv")
	title = loadCSV("data/csv/title.csv")

	// https://github.com/go-gota/gota/issues/122

	_ = renameID(keyword, "keyword")
	_ = renameID(movieInfo, "movie_info")
	_ = renameID(movieKeyword, "movie_keyword")
	_ = renameID(title, "movie")

	fmt.Println()
}

func loadCSV(name string) dataframe.DataFrame {
	fmt.Println("Loading", name)

	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	return dataframe.ReadCSV(f)
}

func renameID(df dataframe.DataFrame, prefix string) error {
	return df.SetNames(append([]string{prefix + "_id"}, df.Names()[1:]...)...)
}
