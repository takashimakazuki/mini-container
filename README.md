# mini-container

## これは何？
簡易コンテナを作成するコマンドプログラムです。

プログラム内で行っていることは以下です
- Linux namespace機能を用いた実行環境の隔離。つまりコンテナ化を行う。
- chrootによってファイルシステムを分離
- cgroupを使ってコンテナが使える計算リソースを制限する
- コンテナプロセス内でのコマンド実行

```bash
$ sudo go run cmd/main.go run ps -ef
Running
Running Child
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 17:14 pts/1    00:00:00 /proc/self/exe child ps -ef
root           5       1  0 17:14 pts/1    00:00:00 ps -ef
```

## namespace, resourceの分離

|namespace|o/x|
|---|---|
|Process Tree|o|
|Mount Point|o|
|User/Group|x|
|Host/Domain Name|o|
|Network Stack|x|
|IPC|x|

## メモ


- forkの実行にスーパユーザ権限が必要になる。(以下のようなエラーが出る)どうにかして回避する方法はないか？dockerの実行にsudo権限が必要なのはこのあたりに起因するものなのかな？
```bash
$ go run cmd/main.go run echo hello
Running
panic: fork/exec /usr/bin/echo: operation not permitted
```

- dockerデーモンの場合、Unix socketにバインドするためにsudo権限が必要らしい。(デフォルトではUnix socketのownerがroot)
> The Docker daemon binds to a Unix socket instead of a TCP port. By default that Unix socket is owned by the user root and other users can only access it using sudo. The Docker daemon always runs as the root user.


- procファイルシステムとは
  - カーネルの現在の状態を示す特別なファイルが置いてある階層
  - アプリケーションやユーザがカーネルのシステム情報を見ることがDきる
  - カーネルの状態とは
    - 疑似端末に関する情報 /proc/tty
    - カーネルのバージョン情報 /proc/version
    - マウントの一覧 /proc/mount
    - システム上のRAM使用率など /proc/meminfo
  - /proc/selfとは？？
    - 現在実行中のプロセスへのリンク（自身のプロセスIDを知らなくても状況を把握できる）
  - /proc/self/exeとは？？
    - 現在実行中のファイルへのシンボリックリンク
    - 以下のコードでは，現在実行中のファイル(cmd/main.go)を，実行中のプロセスから起動している．（引数にchildを渡しているためChild()関数が実行される）
    ```go
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    ```


## 参考資料
- https://kaminashi-developer.hatenablog.jp/entry/dive-into-swamp-container-scratch
- https://qiita.com/woshahua/items/35c4ee5e90fd2c782eca
- https://docs.docker.com/engine/install/linux-postinstall/