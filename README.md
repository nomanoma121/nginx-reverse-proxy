# nginx-reverse-proxy

リバースプロキシの勉強に使ったリポジトリです


```mermaid
graph TD
  client[Client<br>192.168.0.2]

  subgraph backend
    proxy[Reverse Proxy<br>Nginx<br>192.168.0.10, 172.28.0.3]
    server1[server-1<br>172.28.0.5]
    server2[server-2<br>172.28.0.20]
  end

  subgraph frontend
    client
  end

  client --> proxy
  proxy --> server1
  proxy --> server2

```

## 使い方

以下のコマンドでコンテナをビルド、起動します

`docker compose up --build -d`

次に`client`コンテナに入ります

`docker exec -it <client-container-id> sh`

※ client コンテナをわざわざ使っている理由は、server-1.local などのホスト名を使用するためにホスト側で設定をする必要があり、ローカル環境を汚したくないからです。

コンテナ内で以下のコマンドを実行して、プロキシが正しく動作していることを確認します
```bash
curl http://server-1.local
curl http://server-2.local
```
それぞれ異なるレスポンス(文字列)が返ってくれば、リバースプロキシとして正常に振り分けできていることが確認できます。

`/etc/hosts`を見ると、ホスト`server-1.local`、`server-2.local`に同じipが割り当てられていることがわかります。
そのため、実際には同じ Nginx リバースプロキシにアクセスしており、プロキシ側でホスト名をみて`server-1`と`server-2`へルーティングされていることが確認できます。
