package day0

import (
	"rsc.io/quote"
	"utils"
)

func QuoteMessage() string {
	utils.LogDay(0)
	message := quote.Go()
	return message
}
