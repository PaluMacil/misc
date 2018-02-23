package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
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
		switch {
		case r.URL.Path == "/":
			files.ExecuteTemplate(w, "layout", people) //TODO: check error
		case isFile(r.URL):
			fs := http.FileServer(http.Dir("web"))
			fs.ServeHTTP(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})

	log.Println("Now serving on port 3033")
	err = http.ListenAndServe(":3033", nil)
	if err != nil {
		panic(err)
	}
}

func isFile(url *url.URL) bool {
	pathParts := strings.Split(url.Path, `/`)
	lastPart := pathParts[len(pathParts)-1]
	return strings.Contains(lastPart, ".")
}
