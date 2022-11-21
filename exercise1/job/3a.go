package job

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// SELECT MIN(t.title) AS movie_title
// FROM keyword AS k,
//      movie_info AS mi,
//      movie_keyword AS mk,
//      title AS t
// WHERE k.keyword LIKE '%sequel%'
//   AND mi.info IN ('Sweden',
//                   'Norway',
//                   'Germany',
//                   'Denmark',
//                   'Swedish',
//                   'Denish',
//                   'Norwegian',
//                   'German')
//   AND t.production_year > 2005
//   AND t.id = mi.movie_id
//   AND t.id = mk.movie_id
//   AND mk.movie_id = mi.movie_id
//   AND k.id = mk.keyword_id;

func Job3a(keyword, movieInfo, movieKeyword, title dataframe.DataFrame) {
	keyword = keyword.Filter(dataframe.F{
		Colname:    "keyword",
		Comparator: series.CompFunc,
		Comparando: like("%sequel%"),
	})

	movieInfo = movieInfo.Filter(dataframe.F{
		Colname:    "info",
		Comparator: series.In,
		Comparando: []string{"Sweden", "Norway", "Germany", "Denmark", "Swedish", "Denish", "Norwegian", "German"},
	})

	title = title.Filter(dataframe.F{
		Colname:    "production_year",
		Comparator: series.Greater,
		Comparando: 2005,
	})

	result := keyword.InnerJoin(movieKeyword, "keyword_id")
	result = result.InnerJoin(movieInfo, "movie_id")
	result = result.InnerJoin(title, "movie_id")

	sorted := result.Arrange(dataframe.Sort("title")).Select([]string{"title"}).Subset([]int{0})

	_ = sorted.SetNames("movie_title")

	fmt.Println(sorted)
}
