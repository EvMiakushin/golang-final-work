package harvest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func SupportBase(SupportUrl string) []SupportData {
	var newSupp []SupportData
	res, _ := http.Get(SupportUrl)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read data. error:%s\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		if err := json.Unmarshal(body, &newSupp); err != nil {
			log.Printf("Failed to read data. error:%s\n", err)
		}
	} else if res.StatusCode == 500 {
		log.Println("Error ", newSupp)
	}
	return newSupp
}

func SortingSupp(supp []SupportData) (Support []int) {

	Support = make([]int, 0)
	var indLoad int
	loadness := 0
	for _, v := range supp {
		loadness = loadness + v.ActiveTickets
	}
	timewait := 60 / 18 * loadness
	switch {
	case loadness < 9:
		indLoad = 1
	case loadness >= 9 && loadness <= 16:
		indLoad = 2
	default:
		indLoad = 3
	}
	Support = []int{
		indLoad,
		timewait,
	}
	return
}
