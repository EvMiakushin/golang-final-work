package harvest

import "fmt"

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

var (
	ResultS ResultSetT
	Result  ResultT
)

func GetResultData(sms [][]SMSData, mms []MMSData, vicl []VoiceCallData, mail []EmailData, bill BillingData, supp []SupportData, incid []IncidentData) ResultT {
	ResultS.SMS = sms
	ResultS.MMS = SortingMms(mms)
	ResultS.VoiceCall = vicl
	ResultS.Email = SortingMail(mail)
	ResultS.Billing = bill
	ResultS.Support = SortingSupp(supp)
	ResultS.Incidents = incid

	if ResultS.SMS != nil &&
		ResultS.MMS != nil &&
		ResultS.VoiceCall != nil &&
		ResultS.Email != nil &&
		ResultS.Billing == bill &&
		ResultS.Incidents != nil {
		Result.Status = true
		Result.Data = ResultS

	} else {
		Result.Status = false
		Result.Error = fmt.Sprintf("Not all data has been collected...")
	}

	return Result
}
