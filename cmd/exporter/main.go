package main

import (
	"github.com/i-hit/go-lesson3.1.git/pkg/transaction"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := execute("export.xml"); err != nil {
		os.Exit(1)
	}
}


func execute(filename string) error {
	var err error
	var file *os.File
	if file, err = os.Create(filename); err != nil {
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

	for i := 1; i < 100; i++ {
		id := "0000" + strconv.Itoa(i)
		id = id[len(id) - 4:]
		if _, err = svc.Register(id, "0001", "0002", 10_000_00); err != nil {
			log.Println(err)
			return err
		}
	}

	if err = svc.ExportXml(filename); err != nil {
		log.Println(err)
		return err
	}

	return err
}
