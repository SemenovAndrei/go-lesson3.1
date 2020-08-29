package transaction

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Transaction struct {
	XMLName string `xml:"transaction"`
	Id      string `json:"id" xml:"id"`
	From    string `json:"from" xml:"from"`
	To      string `json:"to" xml:"to"`
	Amount  int64 `json:"amount" xml:"amount"`
	Created int64 `json:"created" xml:"created"`
}

type Transactions struct {
	XMLName string `xml:"transactions"`
	Transactions []*Transaction
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

func (s *Service) ExportXml(file string) error {
	var t Transactions
	t.Transactions = s.transactions

	encoded, err := xml.MarshalIndent(t, "", " ")
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


func (s *Service) ImportJson(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &s.transactions)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}
