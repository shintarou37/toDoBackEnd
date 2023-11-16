package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Logger型を定義します。
type Logger struct{}

// ログを記録するメソッド
func (l *Logger) Log(message string) {
	log.Println(message)
}

// ハンドラーのファクトリ関数。Loggerを受け取り、http.HandlerFuncを返します。
func NewSampleHandler(logger *Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log("Sample handler called")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, DI!"))
	}
}

func T(logger *Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log("test")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, DI!"))
	}
}
func main() {
	r := chi.NewRouter()

	// Loggerのインスタンスを作成します。
	logger := &Logger{}

	// ルートの設定
	// ミドルウェアの設定
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	// r.Get("/", NewSampleHandler(logger))
	r.Group(func(r chi.Router) {
		r.Get("/", NewSampleHandler(logger))

		r.Get("/test", T(logger))
	})

	// ポートの設定とサーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // デフォルトのポート
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
