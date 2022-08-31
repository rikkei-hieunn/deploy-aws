/*
Package model declare all models use in application.
*/
package model

// APIResponse contains info response api
type APIResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}