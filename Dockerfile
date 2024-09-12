# ベースイメージとしてGoの公式イメージを使用
FROM golang:1.23.0

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルをコピー
COPY go.mod go.sum ./

# Goモジュールの依存関係をダウンロード
RUN go mod tidy

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# コンテナが起動したときに実行するコマンド
CMD ["./main"]

# ポート8080を公開
EXPOSE 8080
