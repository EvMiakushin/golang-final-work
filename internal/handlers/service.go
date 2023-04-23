package handlers

import (
	"golang-final-work/internal/config"
	"golang-final-work/internal/harvest"
	"sync"
	"time"
)

type Service struct {
	lastRun time.Time
	cache   harvest.ResultT
	mutex   sync.Mutex
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetCacheData() harvest.ResultT {
	var (
		sms   [][]harvest.SMSData
		mms   []harvest.MMSData
		vicl  []harvest.VoiceCallData
		mail  []harvest.EmailData
		bill  harvest.BillingData
		supp  []harvest.SupportData
		incid []harvest.IncidentData
	)
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if time.Since(s.lastRun) < 5*time.Second {
		return s.cache
	}
	conf := config.GetConfig()

	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		defer wg.Done()
		sms = harvest.ReadAndSortSms(conf.SmsFile)
	}()
	go func() {
		defer wg.Done()
		mms = harvest.MmsBase(conf.MmsUrl)
	}()
	go func() {
		defer wg.Done()
		vicl = harvest.ViclBase(conf.ViclFile)
	}()
	go func() {
		defer wg.Done()
		mail = harvest.MailBase(conf.MailFile)
	}()
	go func() {
		defer wg.Done()
		bill = harvest.BillingBase(conf.BillFile)
	}()
	go func() {
		defer wg.Done()
		supp = harvest.SupportBase(conf.SupportUrl)
	}()
	go func() {
		defer wg.Done()
		incid = harvest.IncidentBase(conf.IncidentUrl)
	}()

	wg.Wait()

	result := harvest.GetResultData(sms, mms, vicl, mail, bill, supp, incid)
	s.cache = s.checkCacheData(result)
	s.lastRun = time.Now()

	return result
}

func (s *Service) checkCacheData(result harvest.ResultT) harvest.ResultT {
	if result.Data.SMS == nil {
		result.Data.SMS = s.cache.Data.SMS
	}
	if result.Data.MMS == nil {
		result.Data.MMS = s.cache.Data.MMS
	}
	if result.Data.VoiceCall == nil {
		result.Data.VoiceCall = s.cache.Data.VoiceCall
	}
	if result.Data.Email == nil {
		result.Data.Email = s.cache.Data.Email
	}
	if &result.Data.Billing == nil {
		result.Data.Billing = s.cache.Data.Billing
	}
	if result.Data.Support == nil {
		result.Data.Support = s.cache.Data.Support
	}
	if result.Data.Incidents == nil {
		result.Data.Incidents = s.cache.Data.Incidents
	}
	return result
}
