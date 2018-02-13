package touch

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

var RepoURL string
var SinceString string
var User string
var Password string
var Name string
var Email string

const tempDir = "tmp"

// Do touch command
func Do() {
	var since time.Time
	if SinceString == "" {
		since = time.Now()
	} else {
		var err error
		since, err = time.Parse("060102", SinceString) // YYMMDD
		if err != nil {
			log.Fatalln("Could not parse \"", SinceString, "\":", err)
		}
	}

	authData := &http.BasicAuth{User, Password}

	repo, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:  RepoURL,
		Auth: authData,
	})
	if err != nil {
		log.Printf("Tried user %s, password %s./n", User, Password)
		log.Fatalln("Could not clone repo:", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Fatalln("Could not get working tree:", err)
	}

	// Start with today, while the day is after the "since" date, touch file and go back one day
	// Add second to ensure a blank touch since still runs
	for d := time.Now().Add(time.Second * time.Duration(1)); d.After(since); d = d.AddDate(0, 0, -1) {
		log.Println("Processing touch on", d)

		filename := randTmpFilename()

		err = ioutil.WriteFile(path.Join(tempDir, filename), []byte("dummy file"), 0644)
		if err != nil {
			log.Fatalln("Could not write file", filename, ":", err)
		}

		_, err = w.Add(filename)
		if err != nil {
			log.Fatalln("Could not add dummy file:", err)
		}
		sig := &object.Signature{
			Name:  Name,
			Email: Email,
			When:  d,
		}
		commitHash, err := w.Commit("add dummy file", &git.CommitOptions{Author: sig})
		if err != nil {
			log.Fatalln("Could not commit add of", filename, ":", err)
		}
		log.Println("Commit success:", commitHash)

		err = repo.Push(&git.PushOptions{Auth: authData})
		if err != nil {
			log.Fatalln("Could not push repo to", RepoURL, "because of:", err)
		}
	}

	DeleteTemp()
	log.Println("Tasks complete. You might consider pushing the removal of the temp files as well.")
}

func DeleteTemp() {
	tempFiles, err := filepath.Glob(tempDir + "/*.tmp")
	if err != nil {
		log.Fatalln("Couldn't find temp files to delete.")
	}
	for _, f := range tempFiles {
		os.Remove(f)
	}
}
