package ext_logger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// ログテスト用のヘルパー関数（※要名称変更）
// 参考
// https://zenn.dev/glassonion1/articles/8ac939208bd455
func testHelper(t *testing.T, fn func(io.Writer), fnc func(logLevel, string) error, l logLevel, mes string) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w

	// writerの更新
	fn(w)

	// テスト対象の関数を実行する
	fnc(l, mes)

	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()

	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}

	println(buf.String())
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}
