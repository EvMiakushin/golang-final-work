package harvest

import (
	"log"
	"math"
	"os"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
	DecimalBilling uint8
}

func BillingBase(BillFile string) BillingData {
	var bill BillingData
	var flag uint8
	buff := make([]bool, 8)
	var rez uint8
	file, err := os.ReadFile(BillFile)
	if err != nil {
		log.Println(err)
	}

	for j := 0; j < len(file); j++ {
		if file[j] == 49 {
			buff[len(file)-j-1] = true
			i := float64(len(file) - j - 1)
			rez += uint8(math.Pow(2, i))
		} else {
			buff[len(file)-j-1] = false
		}
		flag = flag << 1
	}
	bill.CreateCustomer = buff[0]
	bill.Purchase = buff[1]
	bill.Payout = buff[2]
	bill.Recurring = buff[3]
	bill.FraudControl = buff[4]
	bill.CheckoutPage = buff[5]
	bill.DecimalBilling = rez

	return bill
}
