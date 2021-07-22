package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHandler myHandler

	h := NoSurf(&myHandler)
	switch v := h.(type) {
	case http.Handler:
		// do nothing
		break
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myHandler myHandler

	h := SessionLoad(&myHandler)
	switch v := h.(type) {
	case http.Handler:
		// do nothing
		break
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler %T", v))
	}
}

