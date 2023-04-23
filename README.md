#### Оглавление:
____
1. [Дипломный проект](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/#дипломный-проект).
2. [Системы](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/#системы).
3. [Сбор данных](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/#сбор-данных).
4. [Отчет о состоянии систем](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/#отчет-о-состоянии-систем).
____

## Дипломный проект
____

Это небольшой сетевой сервис, который принимает запросы по сети и возвращает данные о состоянии систем.
____

### Системы:
- [SMS](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/sms.go):
```GO
type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
} 
```

- [MMS](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/mms.go):
```GO
type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}
```

- [Voice Call](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/vicl.go):
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

- [Support](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/support.go):
```GO
type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}
```

- [Email](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/mail.go):
```GO
type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}
```

- [Incidents](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/incident.go):
```GO
type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}
```

- [Billings](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/billing.go):
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
- [Result](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/internal/harvest/result.go):
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
- [Отчет](https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/blob/master/2023_03_15.png )

____
Для работы димпломного проекта нужен симулятор https://gitlab.skillbox.ru/evgenii_miakushin/go-final-dpo/-/tree/master/sim
