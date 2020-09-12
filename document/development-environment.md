# GoLang
アプリの開発環境に関するドキュメント  
開発環境構築は基本的に [VSCode Remote Container](https://code.visualstudio.com/docs/remote/remote-overview) を用いて行う  

## 事前インストール
VSCode Remote Containers で開発環境を作る場合は以下の3つのインストールが必要  
(それ以外のGoやVSCodeの拡張機能のインストールは不要)
* Visual Studia Code
* Remote - Containers（VSCode拡張機能）
* Docker Desktop for Windows or Mac

## 環境構築設定の解説
TODO(後で頑張るかも)

## 初期構築作業

### 動作確認
1. リポジトリをCloneし、対象のディレクトリでVSCodeを開く
    ```
    git clone git@github.com:qiitacopy/article.git
    cd article
    code .
    ```
2. 画面左下の[><]を押下後に、`Remote-Containers: Reopen in Container`を選択する
    初回はコンテナのビルドを行うため時間がかかる
3. `src/github.com/qiitacopy/article/server.go`を開いた後に、F5キーを押下してVSCodeのデバックモードを開始する(Goのサーバが起動する)
4. `grpc_cli ls localhost:9000 grpc.ArticleService -l` コマンドで登録されたgRPCのサービスが確認できる
5. `grpc_cli call localhost:9000 GetByID "id: 1"` コマンドでgRPCのサービスをコマンドラインで呼び出せる
6. `grpc_cli call localhost:9000 CreateArticle "username: 'テストユーザ', title: 'タイトル', text: 'テキスト'"` コマンドでgRPCのサービスをコマンドラインで呼び出せる

## パイプラインの起動
本リポジトリではコンテナをビルドし、DockerhubへPUSHを行うパイプラインが構築されている  
このパイプラインは以下のコマンドでバージョンのタグをつけた際に実行される  
バージョンの付け方のネーミングルールは先頭にvを付け、その後にMAJOR.MINOR.PATCHを付けるセマンティックバージョニングを採用する  
1. `git tag -a vX.X.X -m "comment" commitID`
2. `git push origin vX.X.X`

## 開発環境更新に関するルール
### VSCodeで用いる開発用の外部ライブラリのインストール
GOPATHを`/workspace`で設定しているため、コンテナ起動後にGUI等でインストールしようとすると、`/workspace/bin`配下にダウンロードされる  
開発用パッケージはリポジトリ管理として共有したいため、追加やアップデートしたい場合は`.devcontainer/Dockerfile`を更新すること  

## FAQ
### VSCodeから `git push`できない
参考:https://qiita.com/y-tsutsu/items/ec984831e6c8262d3ff7