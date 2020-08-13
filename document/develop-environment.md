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

## 動作確認方法
1. リポジトリをCloneし、対象のディレクトリでVSCodeを開く
    ```
    git clone git@github.com:qiitacopy/article.git
    cd article
    code .
    ```
2. 画面左下の[><]を押下後に、`Remote-Containers: Reopen in Container`を選択する
    初回はコンテナのビルドを行うため時間がかかる
3. `src/github.com/qiitacopy/article/server.go`を開いた後に、F5キーを押下してVSCodeのデバックモードを開始する(Goのサーバが起動する)
4. localhost:9000 にアクセスし、`Hello,`と返ってくることを確認する

## 開発環境更新に関するルール
### VSCodeで用いる開発用の外部ライブラリのインストール
GOPATHを`/workspace`で設定しているため、コンテナ起動後にGUI等でインストールしようとすると、`/workspace/bin`配下にダウンロードされる  
開発用パッケージはリポジトリ管理として共有したいため、追加やアップデートしたい場合は`.devcontainer/Dockerfile`を更新すること
