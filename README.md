# mini-container

## これは何？
簡易コンテナを作成するコマンドプログラムです。

プログラム内で行っていることは以下です
- Linux namespace機能を用いた実行環境の隔離。つまりコンテナ化
- 
- cgroupの設定 ()
- コンテナプロセス内でのコマンド実行

```
$ sudo go run cmd/main.go child echo hello
Running Child
hello
```
## メモ


- forkの実行にスーパユーザ権限が必要になる。(以下のようなエラーが出る)どうにかして回避する方法はないか？dockerの実行にsudo権限が必要なのはこのあたりに起因するものなのかな？
```bash
$ go run cmd/main.go run echo hello
Running
panic: fork/exec /usr/bin/echo: operation not permitted
```

## 参考資料
- https://kaminashi-developer.hatenablog.jp/entry/dive-into-swamp-container-scratch