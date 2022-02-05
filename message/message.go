package message

import (
	"fmt"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewMonthlyTotalMessage(monthlyTotal map[string]int) *linebot.TextMessage {
	message := "【お支払い金額】"
	for month, total := range monthlyTotal {
		message += fmt.Sprintf("\n%s %s円", month, strconv.Itoa(total))
	}

	return linebot.NewTextMessage(message)
}
