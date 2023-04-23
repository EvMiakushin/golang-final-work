package harvest

import (
	"golang-final-work/data"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func MailBase(MailFile string) []EmailData {
	var mail EmailData
	var newMail []EmailData
	file, err := os.ReadFile(MailFile)
	if err != nil {
		log.Println(err)
	}
	fileLines := strings.Split(string(file), "\n")

	for _, b := range fileLines {
		fields := strings.Split(b, ";")
		if len(fields) != 3 {
			continue
		}
		_, ok := data.Country[fields[0]]
		if !ok {
			continue
		}
		if _, ok := data.ESP[fields[1]]; !ok {
			continue
		}
		mail.Country = fields[0]
		mail.Provider = fields[1]
		mail.DeliveryTime, err = strconv.Atoi(fields[2])
		if err != nil {
			log.Println(err)
		}
		newMail = append(newMail, mail)

	}

	return newMail
}

func SortingMail(mail []EmailData) map[string][][]EmailData {
	var dataMail map[string][][]EmailData
	dataMail = make(map[string][][]EmailData, 0)
	ctr := make(map[string]int)

	for _, v := range mail {
		ctr[v.Country]++
	}

	for ind, _ := range ctr {
		dataMail[ind] = [][]EmailData{
			fastProv(mail, ind),
			slowProv(mail, ind),
		}
	}

	return dataMail
}

func fastProv(mail []EmailData, ind string) []EmailData {

	fast := make([]EmailData, 0)
	for _, v := range mail {
		if ind == v.Country {
			fast = append(fast, v)
		}
	}
	sort.Slice(fast, func(a, z int) bool {
		return fast[a].DeliveryTime < fast[z].DeliveryTime
	})
	return fast[:3]
}

func slowProv(mail []EmailData, ind string) []EmailData {

	slow := make([]EmailData, 0)
	for _, v := range mail {
		if ind == v.Country {
			slow = append(slow, v)
		}
	}
	sort.Slice(slow, func(a, z int) bool {
		return slow[a].DeliveryTime > slow[z].DeliveryTime
	})
	return slow[:3]
}
