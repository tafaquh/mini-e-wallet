package dto

import (
	"fmt"
)

type CardType int

const (
	Credit CardType = 1
	Debit CardType = 2
)

func (e CardType) String() string {
    switch e {
		case Credit:
			return "credit"
		case Debit:
			return "debit"
		default:
			return fmt.Sprintf("%d", int(e))
    }
}