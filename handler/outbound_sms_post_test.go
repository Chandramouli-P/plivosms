package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"testproject/handler"
	"testproject/service"
)

type caseData struct {
	Body   string
	Result handler.Response
}

func TestKannelLookupUnknown(t *testing.T) {
	rr := httptest.NewRecorder()
	hndl := handler.Handler{
		Kannel: service.NewKannelService("http://unknown", "chandramouli", "Test12", service.GSM7),
	}
	hFunc := http.HandlerFunc(hndl.OutboundSmsPost)

	testData := caseData{
		Body: `{"from": "123123", "to": "123123", "text": "Test, text"}`,
		Result: handler.Response{
			Message: "",
			Error:   "unknown failure",
		},
	}

	req, err := http.NewRequest("POST", "/outbound/sms", bytes.NewBuffer([]byte(testData.Body)))
	if err != nil {
		t.Fatal(err)
	}
	hFunc.ServeHTTP(rr, req)

	response := handler.Response{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != testData.Result.Message && response.Error != testData.Result.Error {
		t.Errorf("handler returned wrong response data: got %v want %v",
			response, testData.Result)
	}

	rr.Body.Reset()
}

func TestKannelResponseIncorrect(t *testing.T) {
	rr := httptest.NewRecorder()
	hndl := handler.Handler{
		Kannel: service.NewKannelService("http://35.154.95.122:13013", "chandramouli", "", service.GSM7),
	}
	hFunc := http.HandlerFunc(hndl.OutboundSmsPost)

	testData := caseData{
		Body: `{"from": "123123", "to": "123123", "text": "Test, text"}`,
		Result: handler.Response{
			Message: "",
			Error:   "unknown failure",
		},
	}

	req, err := http.NewRequest("POST", "/outbound/sms", bytes.NewBuffer([]byte(testData.Body)))
	if err != nil {
		t.Fatal(err)
	}
	hFunc.ServeHTTP(rr, req)

	response := handler.Response{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != testData.Result.Message && response.Error != testData.Result.Error {
		t.Errorf("handler returned wrong response data: got %v want %v",
			response, testData.Result)
	}

	rr.Body.Reset()
}

func TestOutboundSmsPost(t *testing.T) {
	rr := httptest.NewRecorder()
	hndl := handler.Handler{
		Kannel: service.NewKannelService("http://35.154.95.122:13013", "chandramouli", "Test12", service.GSM7),
	}
	hFunc := http.HandlerFunc(hndl.OutboundSmsPost)

	var cases = []caseData{
		caseData{
			Body: `{"from": 123123}`,
			Result: handler.Response{
				Message: "",
				Error:   "unknown failure",
			},
		},
		caseData{
			Body: `{"from": "123123", "to": "", "text": "Test, text"}`,
			Result: handler.Response{
				Message: "",
				Error:   "to is missing",
			},
		},
		caseData{
			Body: `{"from": "123", "to": "123123", "text": "Test, text"}`,
			Result: handler.Response{
				Message: "",
				Error:   "from is invalid",
			},
		},
		caseData{
			Body: `{"from": "123123", "to": "123123", "text": "Test, text"}`,
			Result: handler.Response{
				Message: "outbound sms ok",
				Error:   "",
			},
		},
	}

	for _, caseData := range cases {
		req, err := http.NewRequest("POST", "/outbound/sms", bytes.NewBuffer([]byte(caseData.Body)))
		if err != nil {
			t.Fatal(err)
		}
		hFunc.ServeHTTP(rr, req)

		response := handler.Response{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Message != caseData.Result.Message && response.Error != caseData.Result.Error {
			t.Errorf("handler returned wrong response data: got %v want %v",
				response, caseData.Result)
		}

		rr.Body.Reset()
	}
}
