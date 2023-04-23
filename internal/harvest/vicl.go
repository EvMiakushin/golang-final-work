package harvest

import (
	"golang-final-work/data"
	"log"
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func ViclBase(viclFile string) (newVicl []VoiceCallData) {
	var vicl VoiceCallData
	file, err := os.ReadFile(viclFile)
	if err != nil {
		log.Println(err)
	}
	fileLines := strings.Split(string(file), "\n")

	for _, b := range fileLines {
		fields := strings.Split(b, ";")
		if len(fields) != 8 {
			continue
		}
		_, ok := data.Country[fields[0]]
		if !ok {
			continue
		}
		if _, ok := data.ViCL[fields[3]]; !ok {
			continue
		}
		vicl.Country = fields[0]
		vicl.Bandwidth = fields[1]
		vicl.ResponseTime = fields[2]
		vicl.Provider = fields[3]
		constab, err := strconv.ParseFloat(fields[4], 32)
		if err != nil {
			log.Println(err)
		}
		vicl.ConnectionStability = float32(constab)
		vicl.TTFB, err = strconv.Atoi(fields[5])
		if err != nil {
			log.Println(err)
		}
		vicl.VoicePurity, err = strconv.Atoi(fields[6])
		if err != nil {
			log.Println(err)
		}
		vicl.MedianOfCallsTime, err = strconv.Atoi(fields[7])
		if err != nil {
			log.Println(err)
		}

		newVicl = append(newVicl, vicl)

	}
	return
}
