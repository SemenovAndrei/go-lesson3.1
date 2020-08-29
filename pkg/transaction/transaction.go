package transaction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

type Transaction struct {
	Id      string `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Amount  int64 `json:"amount"`
	Created int64 `json:"created"`
}

type Service struct {
	mu           sync.Mutex
	transactions []*Transaction
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(id, from, to string, amount int64) (string, error) {
	t := &Transaction{
		Id:      id,
		From:    from,
		To:      to,
		Amount:  amount,
		Created: time.Now().Unix(),
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions = append(s.transactions, t)

	return t.Id, nil
}

func (s *Service) ExportJson(file string) error {

	encoded, err := json.MarshalIndent(s.transactions, "", " ")
	if err != nil {
		log.Println(err)
		return nil
	}

	err = ioutil.WriteFile(file, encoded, 0777)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}


func MapRowToTransaction(row []string) (id, from, to string, amount int64) {
	a, err := strconv.Atoi(row[3])
	if err != nil {
		log.Println(err)
		return
	}
	amount = int64(a)

	return row[0], row[1], row[2], amount
}
