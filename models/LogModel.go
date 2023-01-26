package models

type LogModel struct {
	Timestamp          int64
	Method             string
	Url                string
	Body               string
	AdditionalContents string
	Headers            map[string]string
}
