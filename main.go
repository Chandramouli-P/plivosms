package main

import (
	"fmt"
	"net/http"

	"testproject/handler"
	"testproject/service"
	"github.com/goji/httpauth"
	"github.com/pressly/chi"
)

func main() {
	var (
		r             = chi.NewRouter()
		kannelService = service.NewKannelService("http://35.154.95.122:13013", "chandramouli", "Test12", service.UTF8)
		hndl          = handler.Handler{
			Kannel: kannelService,
		}
	)

	r.Use(httpauth.SimpleBasicAuth("chandramouli", "6ab5db01-8bc3-4960-92d7-d83155a42b0c"))

	r.Post("/outbound/sms", hndl.OutboundSmsPost)

	fmt.Println("Stating the server...")

	http.ListenAndServe(":3000", r)
}
