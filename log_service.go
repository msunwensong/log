package logservice

import (
	"errors"

	"github.com/cihub/seelog"
)

type GoLog struct {
	enable bool
}

func (g *GoLog) Init(filename string) error {
	logger, err := seelog.LoggerFromConfigAsFile(filename)
	if err != nil {
		return errors.New("GoLog init fail ")
	}
	seelog.ReplaceLogger(logger)
	g.enable = true
	return nil
}
func (g *GoLog) Error(v ...interface{}) error {
	err := g.checkEnable()
	if err != nil {
		return err
	}
	defer seelog.Flush()
	seelog.Error(v)
	return nil
}
func (g *GoLog) Info(v ...interface{}) error {
	err := g.checkEnable()
	if err != nil {
		return err
	}
	defer seelog.Flush()
	seelog.Info(v)
	return nil
}
func (g *GoLog) Debug(v ...interface{}) error {
	err := g.checkEnable()
	if err != nil {
		return err
	}
	defer seelog.Flush()
	seelog.Debug(v)
	return nil
}

func (g *GoLog) checkEnable() error {
	if !g.enable {
		return errors.New("GoLog Not init")
	}
	return nil
}
