package transaction

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"sync"
	"time"
)

type Transaction struct {
	Id      string
	From    string
	To      string
	Amount  int64
	Created int64
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

func (s *Service) Export(writer io.Writer) error {
	s.mu.Lock()
	if len(s.transactions) == 0 {
		s.mu.Unlock()
		return nil
	}

	records := make([][]string, len(s.transactions))
	for _, t := range s.transactions {
		record := []string{
			t.Id,
			t.From,
			t.To,
			strconv.Itoa(int(t.Amount)),
			strconv.FormatInt(t.Created, 10),
		}
		records = append(records, record)
	}
	s.mu.Unlock()

	w := csv.NewWriter(writer)
	return w.WriteAll(records)
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
