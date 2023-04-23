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

/*
// With strings.Split
func SmsBase(SmsFile string) []SMSData {
	mutex.Lock()
	defer mutex.Unlock()
	var sms SMSData
	var newSms []SMSData

	file, err := os.ReadFile(SmsFile)
	if err != nil {
		log.Println(err)
	}
	fileLines := strings.Split(string(file), "\n")

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
		sms.Country = fields[0]
		sms.Bandwidth = fields[1]
		sms.ResponseTime = fields[2]
		sms.Provider = fields[3]

		newSms = append(newSms, sms)

	}

	return newSms
}

// Not strings.Split
func SmsBase1(SmsFile string) []SMSData {
	var sms SMSData
	var newSms []SMSData
	file, err := os.Open(SmsFile)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = ';'

	for {
		fields, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		_, ok := data.Country[fields[0]]
		if !ok {
			continue
		}
		if _, ok := data.Providers[fields[3]]; !ok {
			continue
		}
		sms.Country = fields[0]
		sms.Bandwidth = fields[1]
		sms.ResponseTime = fields[2]
		sms.Provider = fields[3]
		mutex.Lock()
		newSms = append(newSms, sms)
		mutex.Unlock()
	}

	return newSms
}

func SortingSms(sms []SMSData) [][]SMSData {
	var dataSms [][]SMSData
	smsCopy := make([]SMSData, len(sms))
	for i, v := range sms {
		sms[i].Country = data.Country[v.Country]
	}
	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider < sms[j].Provider
	})
	copy(smsCopy, sms)
	sort.Slice(smsCopy, func(i, j int) bool {
		return smsCopy[i].Country < smsCopy[j].Country
	})
	dataSms = [][]SMSData{
		sms,
		smsCopy,
	}
	return dataSms
}
*/
