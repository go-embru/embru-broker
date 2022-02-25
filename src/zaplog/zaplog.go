package zaplog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var Logger *zap.SugaredLogger


func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())


	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder


	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getLogWriter() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./embru.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	//return zapcore.AddSync(lumberJackLogger)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))


	//file, _ := os.Create("./test.log")
	//return zapcore.AddSync(file)
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	//_logger := zap.New(core)
	_logger := zap.New(core, zap.AddCaller())

	Logger = _logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		Logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		Logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func Test1() {
	defer Logger.Sync()

	for i := 0; i < 100; i++ {
		Logger.Debugf("Trying to hit GET request for %s", "http://www.baidu.com")
		Logger.Infof("Success! statusCode = %s for URL %s", "http://www.baidu.com",200)
		Logger.Errorf("Error fetching URL %s : Error = %s", "http://www.baidu.com",400)
	}


	//simpleHttpGet("www.google.com")
	//simpleHttpGet("http://www.google.com")
}
