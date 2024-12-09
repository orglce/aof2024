package day0

import (
	"rsc.io/quote"
	"testing"
)

func TestQuoteMessage(t *testing.T) {
	expected := quote.Go()
	actual := QuoteMessage()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
