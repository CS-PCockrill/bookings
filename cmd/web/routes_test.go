package main

import (
	"fmt"
	config2 "github.com/CS-PCockrill/bookings-app/internal/config"
	"github.com/go-chi/chi"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config2.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// Do nothing test passed
		break
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux, type is %T", v))
	}

}
