package test

import (
	"blog-server/utils"
	"testing"
)

func TestLogger(t *testing.T) {
	utils.InitLog("log")
	//utils.LogRus.Debug("this is debug log")
	utils.LogRus.Info("this is info log")
	utils.LogRus.Warn("this is warn log")
	//utils.LogRus.Error("this is error log")
	//utils.LogRus.Panic("this is panic log")
}
