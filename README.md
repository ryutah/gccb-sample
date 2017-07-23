# GCCB自動デプロイサンプル
[Google Cloud Container Builder(GCCB)](https://cloud.google.com/container-builder/docs/?hl=ja)を使用し、AppEngine(Standard Environment)へアプリケーションをデプロイするサンプル

# フロー
1. [Cloud Source Repositries](https://cloud.google.com/source-repositories/docs/?hl=ja)へコードをPushする
2. GCCBのトリガーが実行される
3. `cloudbuild.yaml`に従い、ビルドが実行される

# Dockerビルドコマンド
## ビルド用コンテナ
```console
$ docker build -t gcr.io/{PROJECT_NAME}/glide -f Dockerfile.deploy .
```

## デプロイ用コンテナ
```console
$ docker build -t gcr.io/project-ryutah/deploy \
  -f Dockerfile.deploy \
  --build-arg PROJECT_NAME={YOUR_PROJECT_NAME} \
  --build-arg SERVICE_ACCOUNT_ID={SERVICE_ACCOUNT_ID} \
  --build-arg KEY_FILE_NAME={SERVICE_ACCOUNT_KEY_FILE_NAME} \
  . 
```

# 注意点
* プロジェクトルートに、サービスアカウントのキーファイルを配置してからコンテナのビルドをする
