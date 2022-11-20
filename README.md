# sample-webhook

CloudNative Days Tokyo 2022のセッション「Kubernetes Admission Webhook Deep Dive」のサンプルプログラムです。

- セッション情報
    - https://event.cloudnativedays.jp/cndt2022/talks/1579
- 補足記事
    - https://zenn.dev/zoetro/articles/admission-webhook-deep-dive

## サンプルプログラムの実行方法

- Docker をインストール
    -  https://docs.docker.com/get-docker/
- aqua をインストール
    - https://aquaproj.github.io/docs/tutorial-basics/quick-start
- Kubernetes クラスターを起動
    - `make start`
- tilt を起動します
    - `tilt up`
- http://localhost:10350/ にアクセス

## 実装している Webhook

- [api/v1/sampleresource_webhook.go](./api/v1/sampleresource_webhook.go)
    - Defaulter/Validator を利用
    - カスタムリソースの Mutating と Validating を実装
- [hooks/pod_webhook.go](./hooks/pod_webhook.go)
    - CustomDefaulter を利用
    - Pod の imagePullPolicy を強制的に上書き
- [hooks/namespace_webhook.go](./hooks/namespace_webhook.go)
    - CustomValidator を利用
    - アノテーションの付与されていない Namespace の削除を禁止
- [hooks/deployment_webhook.go](./hooks/deployment_webhook.go)
    - Handler を利用
    - deployments リソースと deployments/scale サブリソースに対応
    - replicas の範囲を制限
