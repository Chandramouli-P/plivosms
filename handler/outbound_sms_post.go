package handler

import (
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/pressly/chi/render"
)

// Response struct with data for answer
type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// SMSData struct for binding data from request
type SMSData struct {
	From string `json:"from" validate:"required,gte=6,lte=16"`
	To   string `json:"to" validate:"required,gte=6,lte=16"`
	Text string `json:"text" validate:"required,gte=1,lte=120"`
}

// OutboundSmsPost accept SMS data and handle it
func (h *Handler) OutboundSmsPost(w http.ResponseWriter, r *http.Request) {
	smsData := SMSData{}

	// Bind data to struct
	err := render.Bind(r.Body, &smsData)
	if err != nil {
		render.JSON(w, r, Response{
			Message: "",
			Error:   "unknown failure",
		})
		return
	}

	// Validate data
	err = validator.New().Struct(&smsData)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			switch err.Tag() {
			case "required":
				render.JSON(w, r, Response{
					Message: "",
					Error:   field + " is missing",
				})
			default:
				render.JSON(w, r, Response{
					Message: "",
					Error:   field + " is invalid",
				})
			}
			return
		}
	}

	resp, err := h.Kannel.SendSMS(smsData.From, smsData.To, smsData.Text)
	if err != nil {
		render.JSON(w, r, Response{
			Message: "",
			Error:   "unknown failure",
		})
		return
	}

	if resp != "0: Accepted for delivery" {
		render.JSON(w, r, Response{
			Message: "",
			Error:   "unknown failure",
		})
		return
	}

	render.JSON(w, r, Response{
		Message: "outbound sms ok",
		Error:   "",
	})
}
