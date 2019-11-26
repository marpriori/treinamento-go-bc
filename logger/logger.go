package logger

// Logger é para logar
type Logger interface {
	// Info representa log de info
	Info(string, map[string]interface{})
	// Error representa log de erro mesmo
	Error(string, map[string]interface{})
}
