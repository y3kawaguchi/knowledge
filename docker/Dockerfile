FROM golang:1.14-stretch

# コンテナ内に作業ディレクトリを作成
RUN mkdir /knowledge
# コンテナログイン時のディレクトリ指定
WORKDIR /knowledge
# コンテナにソースをコピー
COPY . /knowledge
# ポートを公開
EXPOSE 8080
# コンテナにコピーしたソースをビルド
RUN make build_linux
# Serverを起動
CMD ./build/knowledge