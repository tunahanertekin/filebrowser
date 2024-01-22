//go:build !dev
// +build !dev

package http

// global headers to append to every response
var globalHeaders = map[string]string{
	"Cache-Control":                    "no-cache, no-store, must-revalidate",
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Headers":     "*",
	"Access-Control-Allow-Methods":     "*",
	"Access-Control-Allow-Credentials": "true",
}
