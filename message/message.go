package message

import (
	"fmt"
	"strconv"
)

func NewMonthlyTotalMessage(monthlyTotal map[string]int) string {
	message := "【お支払い予定金額】"
	if len(monthlyTotal) == 0 {
		message += "\nなし"
	}

	for month, total := range monthlyTotal {
		message += fmt.Sprintf("\n%s %s円", month, strconv.Itoa(total))
	}

	return message
}
