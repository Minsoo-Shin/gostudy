package logger

import "go.uber.org/zap"

type fieldType string

const (
	ID        fieldType = "id"
	UserID    fieldType = "userID"
	RequestID fieldType = "request_id"

	URI     fieldType = "URI"
	URIPath fieldType = "URI_path"
	Status  fieldType = "status"
	Method  fieldType = "method"
	Latency fieldType = "latency"
)

type Field = zap.Field

func String(key fieldType, value string) Field {
	return zap.String(string(key), value)
}

func StringP(key fieldType, value *string) Field {
	return zap.String(string(key), *value)
}

func Int(key fieldType, value int) Field {
	return zap.Int(string(key), value)
}

func IntP(key fieldType, value *int) Field {
	return zap.Intp(string(key), value)
}

func Error(err error) Field {
	return zap.Error(err)
}
