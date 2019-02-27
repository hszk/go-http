package main

import (
    "bytes"
    // フォーマットI/O
    "fmt"
    // http処理関連
    "net/http"

    // カスタム
    "go-http/cmd/http-echo-server/pkg/logger"
)

func main() {
    // httpハンドラを作成
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, r)
    }
    http.HandleFunc("/", handler)
    // httpServer構造体にListenするポートを指定
    err := http.ListenAndServe(":8888", nil)

    var buf bytes.Buffer
    // logger.go:Newを呼び出し
    lgr := logger.New(&buf)
    lgr.Log(err)
    fmt.Println(buf.String())
}
