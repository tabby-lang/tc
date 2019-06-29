// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package message

import (
	"os"

	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func init() {
	timestampFormat := "Mon Jan 2 15:04:05 -0700 MST 2006"
	var disableTimestamp bool

	if len(os.Getenv("TCTimestamp")) == 1 {
		disableTimestamp = false
	} else {
		disableTimestamp = true
	}
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, DisableLevelTruncation: true, DisableTimestamp: disableTimestamp, FullTimestamp: true, TimestampFormat: timestampFormat})
	logrus.SetOutput(colorable.NewColorableStdout())
}

// Error throws an error and aborts the running process.
func Error(v interface{}) {
	logrus.Error(v)
	os.Exit(1)
}

// Warning displays a warning log message.
func Warning(v interface{}) {
	logrus.Warning(v)
}

// Info displays an informational log message.
func Info(v interface{}) {
	logrus.Info(v)
}
