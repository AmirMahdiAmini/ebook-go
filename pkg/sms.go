package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	SMS interface {
		SendSMS() error
	}
	sms struct {
		MessageText string   `json:"messageText"`
		Mobiles     []string `json:"mobiles"`
		LineNumber  string   `json:"lineNumber"`
	}
)

func NewSMS(message string, phone string) SMS {
	return &sms{
		MessageText: message,
		Mobiles:     []string{phone},
		LineNumber:  os.Getenv("SMS_LINE_NUMBER"),
	}
}
func (c *sms) SendSMS() error {
	postBody, _ := json.Marshal(map[string]interface{}{
		"messageText": c.MessageText,
		"mobiles":     c.Mobiles,
		"lineNumber":  c.LineNumber,
	})
	req, err := http.NewRequest("POST", "https://api.sms.ir/v1/send/bulk", bytes.NewBuffer(postBody))
	if err != nil {
		return errors.New("failed to send a sms #22")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", os.Getenv("SMS_SECRET_KEY"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("failed to send a sms #23")
	}
	defer resp.Body.Close()

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		return errors.New("failed to send a sms #24")
	}
	return nil
}
