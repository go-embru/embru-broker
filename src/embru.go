package main

import "embru/zaplog"

func main() {
	zaplog.InitLogger()
	defer zaplog.Logger.Sync()

	zaplog.Logger.Info("ssss")
}
