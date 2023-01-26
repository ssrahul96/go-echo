package models

type LogModel struct {
	Timestamp          int64             `json:"timestamp,omitempty"`
	Method             string            `json:"method,omitempty"`
	Url                string            `json:"url,omitempty"`
	Body               string            `json:"body,omitempty"`
	AdditionalContents string            `json:"additionalContents,omitempty"`
	Headers            map[string]string `json:"headers,omitempty"`
}
