package core

// Logger defines a core interface which a logger must implement to be used
type Logger interface {
	Infoln(...interface{})
	Infof(string, ...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})
}
