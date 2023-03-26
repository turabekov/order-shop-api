package models

type PatchRequest struct {
	ID     string `json:"id"`
	Fields map[string]interface{}
}
