package main

import (
	"github.com/i-hit/go-lesson3.1.git/pkg/transaction"
	"io"
	"log"
	"os"
)

func main() {
	if err := execute("export.json"); err != nil {
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

	svc := transaction.NewService()

	err = svc.ImportJson(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
