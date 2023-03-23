package ext_logger

// このケースでは、設定したログレベルのprefixがログに出力されることを確認する
// func TestGetPrefix(t *testing.T) {
// 	// テスト対象の関数を実行する
// 	l, _ := NewLogger()

// 	// ログ出力を行う部分をヘルパーに流して、標準出力したログを抽出する
// 	got := testHelper(t, l.SetWriter, l.Log, LogLevelInfo, "ログです")

// 	//　設定したprefixがログ出力した文字に含まれていることを確認する
// 	if !strings.Contains(got, LogLevelInfo.GetPrefix()) {
// 		t.Errorf("failed to test. mes logLevel : %s, contains want: %s", got, LogLevelInfo.GetPrefix())
// 	}
// }

// // このケースでは、設定したログレベルによってログが出力制御されないことを確認する
// func TestSetLogLevelOutputIsNotControlled(t *testing.T) {
// 	// テスト対象の関数を実行する
// 	l, _ := NewLogger()

// 	// ログレベル最低を設定する
// 	l.SetLogLevel(LogLevelNone)

// 	// ログ出力を行う部分をヘルパーに流して、標準出力したログを抽出する
// 	got := testHelper(t, l.SetWriter, l.Log, LogLevelInfo, "ログです")

// 	//　設定したprefixがログ出力した文字に含まれていることを確認する
// 	if len(strings.TrimSpace(got)) == 0 {
// 		t.Errorf("failed to test. mes logLevel : %s, contains want: %s", got, LogLevelInfo.GetPrefix())
// 	}
// }

// // このケースでは、設定したログレベルによってログが出力制御されることを確認する
// func TestSetLogLevelOutputControlled(t *testing.T) {
// 	// テスト対象の関数を実行する
// 	l, _ := NewLogger()

// 	// ログレベル最高を設定する
// 	l.SetLogLevel(LogLevelCritical)

// 	// ログ出力を行う部分をヘルパーに流して、標準出力したログを抽出する
// 	got := testHelper(t, l.SetWriter, l.Log, LogLevelInfo, "ログです")

// 	//　設定したprefixがログ出力した文字に含まれていることを確認する
// 	if len(strings.TrimSpace(got)) != 0 {
// 		t.Errorf("failed to test. mes logLevel : %s, contains want: %s", got, LogLevelInfo.GetPrefix())
// 	}
// }
