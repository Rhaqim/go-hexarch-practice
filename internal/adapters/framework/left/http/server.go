package http

import (
	"github.com/Rhaqim/hexarch/internal/ports"
	// "github.com/gin-gonic/gin"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) Run() {
	// gin gonic server
}
