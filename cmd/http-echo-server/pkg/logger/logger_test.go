package logger

import (
	// バイトスライスを操作するための関数
	"bytes"
	// Goパッケージの自動テストをサポート
	"testing"
)

// ファイル名：{テスト対象}_test.go
// 関数名：Test{テスト対象}(t *testing.T)
func TestLog(t *testing.T) {
	// bytes.Bufferを使用（標準入出力の代わり）
	var buf bytes.Buffer
	// logger.goのNewを呼び出す
	logger := New(&buf)
	if logger == nil {
		// loggerに何も入ってなければエラー
		t.Error("New() returned nil")
	} else {
		// 
		logger.Log("Hello, I am a logger!")
		if buf.String() != "Hello, I am a logger!" {
			t.Errorf("\nhave: '%s'\nwant: '%s'", buf.String(), "Hello, I am a logger!")
		}
	}
}
