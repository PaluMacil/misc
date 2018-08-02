package main

import (
	"fmt"
	"log"
	"os"
	"path"
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
	{
		ID:            4,
		Name:          "Alien Kraus",
		Tall:          true,
		NextHug:       time.Now(),
		FavoriteFoods: []string{"hummus", "slime", "gilded harps"},
		FavoriteColor: "gray",
	},
}

func main() {
	const dataDir = "data"
	const indexName = "friends"
	var indexFile = path.Join(dataDir, "indexes", indexName)
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
	for i, f := range friends {
		log.Println("indexing", i+1, "of", len(friends), "friends:", f.Name)
		id := strconv.Itoa(f.ID)
		err := index.Index(id, f)
		if err != nil {
			log.Fatalln("indexing", f.Name+":", err)
		}
	}

	// search for exact text
	query := bleve.NewTermQuery("hotdog")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Println("FAILURE! Search 1:", err)
	}
	fmt.Println("Results for hotdog...")
	fmt.Println(searchResults)
	// search for text starting with...
	query2 := bleve.NewPrefixQuery("hum")
	search2 := bleve.NewSearchRequest(query2)
	searchResults2, err := index.Search(search2)
	if err != nil {
		log.Println("FAILURE! Search 2:", err)
	}
	fmt.Println("Results for prefix 'hum'...")
	fmt.Println(searchResults2)
}
