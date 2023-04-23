package harvest

import (
	"golang-final-work/data"
	"log"
	"os"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func ReadAndSortSms(SmsFile string) [][]SMSData {
	file, err := os.ReadFile(SmsFile)
	if err != nil {
		log.Println(err)
	}
	fileLines := strings.Split(string(file), "\n")

	var newSms []SMSData
	for _, b := range fileLines {
		fields := strings.Split(b, ";")
		if len(fields) != 4 {
			continue
		}
		_, ok := data.Country[fields[0]]
		if !ok {
			continue
		}
		if _, ok := data.Providers[fields[3]]; !ok {
			continue
		}
		sms := SMSData{
			Country:      fields[0],
			Bandwidth:    fields[1],
			ResponseTime: fields[2],
			Provider:     fields[3],
		}
		newSms = append(newSms, sms)
	}

	smsCopy := make([]SMSData, len(newSms))
	for i := 0; i < len(newSms); i++ {
		newSms[i].Country = data.Country[newSms[i].Country]
	}
	sort.Slice(newSms, func(i, j int) bool {
		return newSms[i].Provider < newSms[j].Provider
	})

	copy(smsCopy, newSms)
	sort.Slice(smsCopy, func(i, j int) bool {
		return smsCopy[i].Country < smsCopy[j].Country
	})

	dataSms := [][]SMSData{
		newSms,
		smsCopy,
	}

	return dataSms
}
