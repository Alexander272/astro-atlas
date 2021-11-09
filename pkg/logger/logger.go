package logger

type ILogger interface {
	Debug(msg string, params map[string]interface{})
	Info(msg string, params map[string]interface{})
	Error(msg string, params map[string]interface{})
	Fatal(msg string, params map[string]interface{})
}
