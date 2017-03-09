package service_test

import (
	"testing"

	"testproject/service"
)

func TestKannelServiceLookupUnknown(t *testing.T) {
	kannel := service.NewKannelService("http://unknown", "chandramouli", "Test12", service.GSM7)

	_, err := kannel.SendSMS("79991002010", "79992001020", "Test")
	if err == nil {
		t.Fatal("error: got %v want %v", err, "...dial tcp: lookup unknown...")
	}
}

func TestKannelServiceGSM7(t *testing.T) {
	kannel := service.NewKannelService("http://35.154.95.122:13013", "chandramouli", "Test12", service.GSM7)

	resp, err := kannel.SendSMS("79991002010", "79992001020", "Test")
	if err != nil {
		t.Fatal(err)
	}

	expectedData := "0: Accepted for delivery"
	if resp != expectedData {
		t.Fatal("error: got %v want %v", resp, expectedData)
	}
}

func TestKannelServiceUCS2(t *testing.T) {
	kannel := service.NewKannelService("http://35.154.95.122:13013", "chandramouli", "Test12", service.UCS2)

	resp, err := kannel.SendSMS("79991002010", "79992001020", "Test")
	if err != nil {
		t.Fatal(err)
	}

	expectedData := "0: Accepted for delivery"
	if resp != expectedData {
		t.Fatal("error: got %v want %v", resp, expectedData)
	}
}
