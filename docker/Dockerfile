FROM golang:1.14

# コンテナ内に作業ディレクトリを作成
RUN mkdir /knowledge
# コンテナログイン時のディレクトリ指定
WORKDIR /knowledge
# ソースをコピー
COPY . /knowledge
# ポートを公開
EXPOSE 8080
# ビルド
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o build/knowledge -v ./cmd/knowledge
# Serverを起動
CMD ./build/knowledge