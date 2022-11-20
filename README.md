# sample-webhook

CloudNative Days Tokyo 2022のセッション「Kubernetes Admission Webhook Deep Dive」のサンプルプログラムです。

- セッション情報
    - https://event.cloudnativedays.jp/cndt2022/talks/1579
- 補足記事
    - https://zenn.dev/zoetro/articles/admission-webhook-deep-dive

## サンプルプログラムの実行方法

- Docker をインストールします。
    -  https://docs.docker.com/get-docker/
- aqua をインストールします
    - https://aquaproj.github.io/docs/tutorial-basics/quick-start
- Kubernetes クラスターを起動します
    - `make start`
- tilt を起動します
    - `tilt up`
- http://localhost:10350/ にアクセスします
