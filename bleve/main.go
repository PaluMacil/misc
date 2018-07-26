package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/blevesearch/bleve"
)

type Friends struct {
	ID            int
	Name          string
	Tall          bool
	NextHug       time.Time
	FavoriteFoods []string
	FavoriteColor string
}

var friends = []Friends{
	{
		ID:            1,
		Name:          "Sally",
		Tall:          true,
		NextHug:       time.Now(),
		FavoriteFoods: []string{"green beans", "hotdog", "banana"},
		FavoriteColor: "pink",
	},
	{
		ID:            2,
		Name:          "Joe",
		Tall:          false,
		NextHug:       time.Now(),
		FavoriteFoods: []string{"squash", "lemon"},
		FavoriteColor: "blue",
	},
	{
		ID:            3,
		Name:          "Ben",
		Tall:          true,
		NextHug:       time.Now(),
		FavoriteFoods: []string{"humans", "zebra", "carrot"},
		FavoriteColor: "green",
	},
}

func main() {
	const indexFile = "example.bleve"
	var (
		err   error
		index bleve.Index
	)
	// open a new index if not exists; could check for err of
	// bleve.ErrorIndexPathDoesNotExist after open instead
	if _, err = os.Stat(indexFile); os.IsNotExist(err) {
		// path/to/whatever does not exist
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexFile, mapping)
		if err != nil {
			log.Fatalln(err)
		}
	} else if err != nil {
		log.Fatalln("checking for index:", err)
	} else {
		index, err = bleve.Open(indexFile)
		if err != nil {
			log.Fatalln("opening index:", err)
		}
	}

	// index some data
	for _, f := range friends {
		id := strconv.Itoa(f.ID)
		err := index.Index(id, f)
		if err != nil {
			log.Fatalln("indexing", f.Name+":", err)
		}
	}

	// search for some text
	query := bleve.NewMatchQuery("hotdog")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatalln("search 1:", err)
	}
	fmt.Println("Error:", err)
	fmt.Println("Results for hotdag...")
	fmt.Println(searchResults)
	fmt.Println(searchResults.Hits[0].Fields)
	// search for text starting with...
	query2 := bleve.NewPrefixQuery("hum")
	search2 := bleve.NewSearchRequest(query2)
	searchResults2, err := index.Search(search2)
	if err != nil {
		log.Fatalln("search 2:", err)
	}
	fmt.Println("Error:", err)
	fmt.Println("Results for prefix 'hum'...")
	fmt.Println(searchResults2)
	fmt.Println(searchResults2.Hits[0].Fields)
}
