# AtCoder コンテスト提出結果取得
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/woinary/AtCoderABCReport?filename=go.mod) ![GitHub](https://img.shields.io/github/license/woinary/AtCoderABCReport)
## 概要

AtCoder Beginner Contestの提出結果を取得してCSV出力します。
以下の条件で絞り込みができます。

* 問題
* 言語
* 結果
* ユーザ

検索条件や何回目のABCの情報を取得するかは設定ファイル(YAML)に記載します。
記載したファイルを引数に指定して実行すると、設定ファイル名の拡張子を.csvに変えたファイル名で出力します。
なお、負荷軽減のため、2ページ目以降の取得時には1分のウェイトを入れています。

## 備考

* コードはGo 1.18.1で確認しています。

以上
