package main

import (
	"context"
	"net/http"
)

// Payload ... starting point for payload
type Payload struct {
	Code       int
	Ctx        context.Context
	W          http.ResponseWriter
	Data       []byte   `json:"data,omitempty"`
	Error      error
}

// Sample ... sample struct with diff data types
type Sample struct {
	Name string `json:"name,omitempty"`
	Number int `json:"number,omitempty"`
	List []string `json:"list,omitempty"`
	Map map[string]string `json:"map,omitempty"`
}
