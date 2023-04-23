package harvest

import (
	"encoding/json"
	"golang-final-work/data"
	"io"
	"log"
	"net/http"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func MmsBase(MmsUrl string) []MMSData {
	var newMms []MMSData

	res, _ := http.Get(MmsUrl)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read data. error:%s\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		if err := json.Unmarshal(body, &newMms); err != nil {
			log.Printf("Failed to read data. error:%s\n", err)
		}

		ind := 0
		var c, p bool
		for i := 0; i < len(newMms); i++ {
			c, p = true, true
			if _, ok := data.Country[newMms[i].Country]; !ok {
				c = false
			}
			if _, ok := data.Providers[newMms[i].Provider]; !ok {
				p = false
			}
			if c != true || p != true {
				ind++
				newMms[i] = newMms[len(newMms)-1]
				newMms = newMms[:len(newMms)-1]
				continue
			}
		}

	} else if res.StatusCode == 500 {
		log.Println("Error ", newMms)
	}
	return newMms
}

func SortingMms(mms []MMSData) [][]MMSData {
	var dataMms [][]MMSData

	mmsCopy := make([]MMSData, len(mms))
	for i, v := range mms {
		mms[i].Country = data.Country[v.Country]
	}
	sort.Slice(mms, func(a, z int) bool {
		return mms[a].Provider < mms[z].Provider
	})
	copy(mmsCopy, mms)
	sort.Slice(mmsCopy, func(a, z int) bool {
		return mmsCopy[a].Country < mmsCopy[z].Country
	})
	dataMms = [][]MMSData{
		mms,     //first slice
		mmsCopy, //second slice
	}
	return dataMms
}
