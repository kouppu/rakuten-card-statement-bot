package rakuten

import (
	"strconv"
	"time"
)

type Payment struct {
	Title    string
	Price    int
	UsedAt   time.Time
	DueMonth string
}

type Statement struct {
	Payments []*Payment
}

func NewStatement(statementCsvRecords [][]string) (*Statement, error) {
	var payments []*Payment

	for index, record := range statementCsvRecords {
		if index == 0 {
			continue
		}
		if len(record[0]) == 0 {
			// 備考列があるケースがある
			continue
		}

		price, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, err
		}

		usedAt, err := time.Parse("2006/01/02", record[0])
		if err != nil {
			return nil, err
		}

		payment := newPayment(record[1], price, usedAt, record[7])
		payments = append(payments, payment)

	}

	return &Statement{
		Payments: payments,
	}, nil

}

func GetMonthlyTotal(s *Statement) map[string]int {
	m := make(map[string]int)
	for _, p := range s.Payments {
		m[p.DueMonth] += p.Price
	}

	return m
}

func newPayment(title string, price int, usedAt time.Time, dueMonth string) *Payment {
	return &Payment{
		Title:    title,
		Price:    price,
		UsedAt:   usedAt,
		DueMonth: dueMonth,
	}
}
