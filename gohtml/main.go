package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Person is example data to be passed into the template
type Person struct {
	Name          string
	Age           int
	FavoriteColor string
}

func main() {
	files, err := template.ParseGlob("web/*.gohtml")
	if err != nil {
		log.Fatalln("Couldn't find template files.")
	}

	people := []Person{
		Person{Name: "Sam", Age: 4, FavoriteColor: "Blue"},
		Person{Name: "Jerry", Age: 34, FavoriteColor: "Red"},
		Person{Name: "Georgina", Age: 23, FavoriteColor: "Green"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("web"))
		pathParts := strings.Split(r.URL.Path, `/`)
		lastPart := pathParts[len(pathParts)-1]
		if strings.Contains(lastPart, ".") {
			fs.ServeHTTP(w, r)
		} else if r.URL.Path == "/" { //TODO: simplify logic and make production safe (even examples shouldn't use something dangerous)
			files.ExecuteTemplate(w, "layout", people)
		}
	})

	fmt.Println("Now serving on port 3033")
	err = http.ListenAndServe(":3033", nil)
	if err != nil {
		panic(err)
	}
}
