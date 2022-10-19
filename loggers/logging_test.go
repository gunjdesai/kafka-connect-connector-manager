package loggers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerInit(t *testing.T) {

	log, err := New("warn")
	assert.Nil(t, err)
	assert.NotNil(t, log.log)

	l, e := New("")
	assert.Nil(t, e)
	assert.NotNil(t, l.log)

	logFatal, errFatal := New("fatal")
	assert.Nil(t, errFatal)
	assert.NotNil(t, logFatal)

	logDebug, errDebug := New("debug")
	assert.Nil(t, errDebug)
	assert.NotNil(t, logDebug)

	logInfo, errInfo := New("info")
	assert.Nil(t, errInfo)
	assert.NotNil(t, logInfo)

	logPanic, errPanic := New("panic")
	assert.Nil(t, errPanic)
	assert.NotNil(t, logPanic)

	l.Info("info")
	l.Debug("debug")
	l.Warn("warn")
	assert.True(t, true)

}
