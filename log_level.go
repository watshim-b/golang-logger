package ext_logger

type logLevel int

const (
	// 何もログレベルの情報を出さない場合に利用する
	LogLevelNone logLevel = iota

	// デバッグレベルのログ出力
	LogLevelDebug

	//  情報レベルのログ出力
	LogLevelInfo

	// 警告レベルのログ出力
	LogLevelWarning

	// エラー発生時のログ出力
	LogLevelError

	// エラーの中で、アラート発砲が必要なエラーが発生した場合に利用するログレベルy
	LogLevelAlert

	// 致命的な問題が発生した際のログ出力
	LogLevelCritical
)

// ログ出力の際にどのレベルのログなのかがわかる情報を出力する
func (l logLevel) GetPrefix() string {
	switch l {
	case LogLevelDebug:
		return "【DEBUG】"

	case LogLevelInfo:
		return "【INFO】"

	case LogLevelWarning:
		return "【WARN】"

	case LogLevelError:
		return "【ERROR】"

	case LogLevelAlert:
		return "【ERR-A】"

	case LogLevelCritical:
		return "【ERR-CR】"

	case LogLevelNone:
		return ""

	default:
		return ""

	}
}

/**
* ログ出力の際にどのレベルのログなのかがわかる情報を出力する
* 引数で渡されたログレベルが自身よりも小さい場合は、ログに出力するという制御をかける
 */
//
func (l logLevel) shouldOutLog(ll logLevel) bool {
	if l <= ll {
		return true
	}
	return false
}
