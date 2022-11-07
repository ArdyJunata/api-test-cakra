package logger

type LogKey string

const (
	RESPONSE_TIME LogKey = "RESPONSE_TIME"
	RESPONSE_TYPE LogKey = "RESPONSE_TYPE"
	SERVICES_NAME LogKey = "SERVICES_NAME"
	TRACER_ID     LogKey = "TRACER_ID"
	PATH          LogKey = "PATH"
	METHOD        LogKey = "METHOD"
	DATA          LogKey = "DATA"
)
