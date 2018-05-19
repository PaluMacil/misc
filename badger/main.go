package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger"
)

var usage = `
This is a simple demo. Valid usage of the badger command:
	add [key] [value]
	get [key]
	rm [key]
	ls
`

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println(usage)
		return
	}

	opts := badger.DefaultOptions
	opts.Dir = "data"
	opts.ValueDir = "data"
	db, err := badger.Open(opts)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	switch os.Args[1] {
	case "add":
		if len(os.Args) != 4 {
			fmt.Println(usage)
			return
		}
		err := db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte(os.Args[2]), []byte(os.Args[3]))
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Panicln(err)
		}
	case "get":
		if len(os.Args) != 3 {
			fmt.Println(usage)
			return
		}
		err := db.View(func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(os.Args[2]))
			if err != nil {
				return err
			}
			value, err := item.Value()
			if err != nil {
				return err
			}
			fmt.Printf("Value: %s", value)
			return nil
		})
		if err != nil {
			log.Panicln(err)
		}
	case "rm":
		if len(os.Args) != 3 {
			fmt.Println(usage)
			return
		}
		err := db.Update(func(txn *badger.Txn) error {
			err = txn.Delete([]byte(os.Args[2]))
			if err != nil {
				return err
			}
			return txn.Commit(nil)
		})
		if err != nil {
			log.Panicln(err)
		}
	case "ls":
		if len(os.Args) != 2 {
			fmt.Println(usage)
			return
		}
		err = db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			opts.PrefetchSize = 10
			it := txn.NewIterator(opts)
			defer it.Close()
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				k := item.Key()
				v, err := item.Value()
				if err != nil {
					return err
				}
				fmt.Printf("key=%s, value=%s\n", k, v)
			}
			return nil
		})
		if err != nil {
			log.Panicln(err)
		}
	default:
		fmt.Println(usage)
		return
	}
}
