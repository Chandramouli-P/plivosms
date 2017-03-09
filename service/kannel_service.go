package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/yazver/gsmmodem/pdu"
)

type Coding int

const (
	GSM7 Coding = iota
	UTF8
	UCS2
)

// IKannel interface for SMS services
type IKannel interface {
	SendSMS(string, string, string) (string, error)
}

// KannelService struct about Kannel service
type KannelService struct {
	url      string
	username string
	password string
	coding   Coding

	requestClient *http.Client
}

// NewKannelService returns a pointer to struct
func NewKannelService(url, username, password string, coding Coding) *KannelService {
	return &KannelService{
		url:      url,
		username: username,
		password: password,
		coding:   coding,

		requestClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// SendSMS does request to Kannel service
func (ks *KannelService) SendSMS(from, to, text string) (string, error) {
	switch ks.coding {
	case GSM7:
		text = string(pdu.Encode7Bit(text))
	case UCS2:
		text = string(pdu.EncodeUcs2(text))
	}

	params := url.Values{}
	params.Set("username", ks.username)
	params.Set("password", ks.password)
	params.Set("smsc", "SMPPSim")
	params.Set("from", from)
	params.Set("to", to)
	params.Set("text", text)
	params.Set("coding", fmt.Sprintf("%d", ks.coding))

	req, err := http.NewRequest("GET", ks.url+"/cgi-bin/sendsms", nil)
	if err != nil {
		return "", err
	}
	req.URL.RawQuery = params.Encode()

	resp, err := ks.requestClient.Do(req)
	if err != nil {
		return "", err
	} else if resp != nil {
		defer resp.Body.Close()
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
