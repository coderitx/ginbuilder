package templates

var LoggerxTemplate = `package logx

import (
	"{{.pkgname}}/config/internal_config"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DefaultLogPath = "/var/agent/" // 默认输出日志文件路径
const DefaultLogName = "default.log"

// InitLogger 初始化 log
func InitLogger(conf internal_config.LogConfig) error {
	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}

	// init log channel
	logCh = make(chan []byte, 2048)

	writeSyncer, err := getLogWriter(conf) // 日志文件配置 文件位置和切割
	if err != nil {
		return err
	}
	encoder := getEncoder(conf)          // 获取日志输出编码
	level, ok := logLevel[conf.LogLevel] // 日志打印级别
	if !ok {
		level = logLevel["info"]
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller()) // zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	zap.ReplaceGlobals(logger)
	//zap.L().Debug("")
	//zap.S().Debugf("")
	return nil
}

// getEncoder 编码器(如何写入日志)
func getEncoder(conf internal_config.LogConfig) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeTime = customTimeEncoder            // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if conf.LogFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(conf internal_config.LogConfig) (zapcore.WriteSyncer, error) {
	conf.LogPath = fmt.Sprintf("%v/%v", conf.LogPath, time.Now().Format("2006-01-02"))
	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(conf.LogPath); !exist {
		if conf.LogPath == "" {
			conf.LogPath = DefaultLogPath
		}
		if conf.LogFileName == "" {
			conf.LogFileName = DefaultLogName
		}
		if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
			conf.LogPath = DefaultLogPath
			if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}
	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFileName), // 日志文件路径
		MaxSize:    conf.LogFileMaxSize,                           // 单个日志文件最大多少 mb
		MaxBackups: conf.LogFileMaxBackups,                        // 日志备份数量
		MaxAge:     conf.LogMaxAge,                                // 日志最长保留时间
		Compress:   conf.LogCompress,                              // 是否压缩日志
	}

	// 日志同时输出到控制台、日志文件和grpc log中
	var writeSyncers []zapcore.WriteSyncer
	writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
	writeSyncers = append(writeSyncers, zapcore.AddSync(lumberJackLogger))

	grpcWrite := zapcore.AddSync(&GrpcLogChan{})
	writeSyncers = append(writeSyncers, grpcWrite)
	return zapcore.NewMultiWriteSyncer(writeSyncers...), nil

	// if conf.LogStdout {
	// 	// 日志同时输出到控制台和日志文件中
	// 	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	// } else {
	// 	// 日志只输出到日志文件
	// 	return zapcore.AddSync(lumberJackLogger), nil
	// }
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

type GrpcLogChan struct {
}

var logCh chan []byte

func (w *GrpcLogChan) Write(p []byte) (n int, err error) {
	d := make([]byte, len(p))
	copy(d, p)
	select {
	case logCh <- d:
	default:
	}

	return
}

func (w *GrpcLogChan) Sync() error {
	return nil
}`
