package helper

import (
	"github.com/leekchan/accounting"
)

// Accounting function for parsing decimal value to format rupiah
func Accounting(value float64) string {
	ac := accounting.Accounting{Symbol: "Rp ", Precision: 2, Thousand: ".", Decimal: ","}
	valueRupiah := ac.FormatMoney(value)
	return valueRupiah
}
