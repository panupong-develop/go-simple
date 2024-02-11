package logging

// https://www.somkiat.cc/golang-logging-with-zap/
// https://github.com/uber-go/zap

type ILogHandler interface {
}

type ILogger interface {
	// Debug: Most detailed messages, intended for development and troubleshooting.
	// They might include function calls, variable values, and internal program flow.
	// (Ex: "Entering function CalculateAverage with values [1, 2, 3]")
	Debug(msg string, fields ...any)

	// Info: Informational messages about routine events, starting/stopping services,
	// successful operations, or user logins. These messages convey
	// the application's overall health and activity.
	// (Ex: "User JohnDoe logged in successfully.")
	Info(msg string, fields ...any)

	// Warn: Messages indicating potential problems but not necessarily causing failures.
	// They might point to resource limitations, unusual conditions,
	// or configuration issues. (Ex: "Database connection pool running low")
	Warn(msg string, fields ...any)

	// Error: Messages indicating failed operations or unexpected conditions.
	// These messages require attention and might signify issues impacting functionality.
	// (Ex: "Error writing to file: Permission denied")
	Error(msg string, fields ...any)

	// Fatal: Messages indicating critical failures that prevent the application from
	// functioning correctly. These messages require immediate action and might trigger
	// alerts or restarts.
	// (Ex: "Fatal error: Application core component crashed")
	// if dev:
	// 	 DPanic()
	// else
	//   Panic()
	Panic(msg string, fields ...any)

	// FatalLevel logs a message, then calls os.Exit(1).
	Fatal(msg string, fields ...any)
}
