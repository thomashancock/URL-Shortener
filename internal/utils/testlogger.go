package utils

import "fmt"

// TestLogger implements the Logger interface but records the last log
type TestLogger struct {
	PrevLog string
}

// Infoln stores the log rather than printing it
func (l *TestLogger) Infoln(args ...interface{}) {
	l.PrevLog = fmt.Sprintln(args...)
}

// Infof stores the log rather than printing it
func (l *TestLogger) Infof(format string, args ...interface{}) {
	l.PrevLog = fmt.Sprintf(format, args...)
}

// Errorln stores the log rather than printing it
func (l *TestLogger) Errorln(args ...interface{}) {
	l.PrevLog = fmt.Sprintln(args...)
}

// Errorf stores the log rather than printing it
func (l *TestLogger) Errorf(format string, args ...interface{}) {
	l.PrevLog = fmt.Sprintf(format, args...)
}

// Debugln stores the log rather than printing it
func (l *TestLogger) Debugln(args ...interface{}) {
	l.PrevLog = fmt.Sprintln(args...)
}

// Debugf stores the log rather than printing it
func (l *TestLogger) Debugf(format string, args ...interface{}) {
	l.PrevLog = fmt.Sprintf(format, args...)
}

// Fatalln stores the log rather than printing it
func (l *TestLogger) Fatalln(args ...interface{}) {
	l.PrevLog = fmt.Sprintln(args...)
}

// Fatalf stores the log rather than printing it
func (l *TestLogger) Fatalf(format string, args ...interface{}) {
	l.PrevLog = fmt.Sprintf(format, args...)
}
