package core

import (
	"github.com/sirupsen/logrus"
)

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
