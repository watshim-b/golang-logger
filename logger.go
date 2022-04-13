package ext_logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/uuid"
)

type extLogger struct {
	uniqueKey    string
	baseLogLevel logLevel
	writer       io.Writer
}

// すべてのログレベルを出力する場合にこちらの関数で初期化する
func NewLogger() (*extLogger, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &extLogger{uniqueKey: u.String(), baseLogLevel: LogLevelNone, writer: os.Stdout}, nil
}

//　出力すべきログレベルを指定する
func (l *extLogger) SetLogLevel(ll logLevel) {
	l.baseLogLevel = ll
}

//　出力すべきログレベルを指定する
func (l *extLogger) SetWriter(w io.Writer) {
	l.writer = w
}

// ログレベル次第で出力をするかどうかを判断して、出力する必要がある場合は、ログを出力する
func (l *extLogger) Log(ll logLevel, message string) error {
	// ログレベルの設定上出力する必要があるかどうかを判断し、出力する必要がない場合は、nilだけ返す
	if !l.baseLogLevel.shouldOutLog(ll) {
		return nil
	}

	// ログを出力する
	log := fmt.Sprintf(strings.Join([]string{"%s", message, " : unique_key【%s】\n"}, ""), ll.GetPrefix(), l.uniqueKey)
	_, err := l.writer.Write([]byte(log))
	return err
}

// ログを出力する（interfaceを柔軟に置換してくれるログ出力機能を作成したい）
// func (l *extLogger) Log(template string, i interface{}...) {
// 	fmt.Sprintf(strings.Join("%s", template), l.logLevel.GetPrefix(), )
// }
