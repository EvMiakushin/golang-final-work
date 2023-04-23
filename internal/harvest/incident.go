package harvest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}

func IncidentBase(IncidentUrl string) []IncidentData {

	var newIncid []IncidentData
	res, _ := http.Get(IncidentUrl)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read data. error:%s\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		if err := json.Unmarshal(body, &newIncid); err != nil {
			log.Printf("Failed to read data. error:%s\n", err)
		}
	} else if res.StatusCode == 500 {
		log.Println("Error ", newIncid)
	}
	sort.Slice(newIncid, func(a, z int) bool {
		return newIncid[a].Status < newIncid[z].Status
	})
	return newIncid
}
