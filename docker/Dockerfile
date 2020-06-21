FROM golang:1.13

# コンテナ内に作業ディレクトリを作成
RUN mkdir /knowledge
# コンテナログイン時のディレクトリ指定
WORKDIR /knowledge
# ソースをコピー
COPY . /knowledge
# ポートを公開
EXPOSE 80
# ビルド
RUN go build -o build/knowledge -v ./cmd/knowledge/
# Serverを起動
CMD ./build/knowledge