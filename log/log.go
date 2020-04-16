package log

import "github.com/sirupsen/logrus"

func Errorf(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	logrus.Error(args)
}
