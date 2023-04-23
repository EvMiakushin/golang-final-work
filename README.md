#### Оглавление:
____
1. [Дипломный проект](https://github.com/EvMiakushin/golang-final-work/#дипломный-проект).
2. [Системы](https://github.com/EvMiakushin/golang-final-work/#системы).
3. [Сбор данных](https://github.com/EvMiakushin/golang-final-work/#сбор-данных).
4. [Отчет о состоянии систем](https://github.com/EvMiakushin/golang-final-work/#отчет-о-состоянии-систем).
____

## Дипломный проект
____

Это небольшой сетевой сервис, который принимает запросы по сети и возвращает данные о состоянии систем.
____

### Системы:
- [SMS](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/sms.go):
```GO
type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
} 
```

- [MMS](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/mms.go):
```GO
type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}
```

- [Voice Call](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/vicl.go):
```GO
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
```

- [Support](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/support.go):
```GO
type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}
```

- [Email](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/mail.go):
```GO
type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}
```

- [Incidents](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/incident.go):
```GO
type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}
```

- [Billings](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/billing.go):
```GO
type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
        DecimalBilling uint8
}
```
____

### Сбор данных:
- [Result](https://github.com/EvMiakushin/golang-final-work/-/blob/main/internal/harvest/result.go):
```GO
type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}
```
____

### Отчет о состоянии систем:
- [Отчет](https://github.com/EvMiakushin/golang-final-work/-/blob/main/2023_03_15.png )

____
Для работы димпломного проекта нужен симулятор https://github.com/EvMiakushin/golang-final-work/-/tree/main/sim
