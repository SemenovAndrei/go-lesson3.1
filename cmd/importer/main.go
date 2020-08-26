package main

import (
	"encoding/csv"
	"github.com/i-hit/go-lesson3.1.git/pkg/transaction"
	"io"
	"log"
	"os"
)

func main() {
	if err := execute("export.csv"); err != nil {
		os.Exit(1)
	}
}

func execute(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return err
	}

	svc := transaction.NewService()

	for value := range records {
		if _, err = svc.Register(transaction.MapRowToTransaction(records[value])); err != nil {
			log.Println(err)
			return err
		}
	}

	return err
}
