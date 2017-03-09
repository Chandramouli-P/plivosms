package handler

import "testproject/service"

// Handler main struct for all handler functions of API
type Handler struct {
	Kannel service.IKannel
}
