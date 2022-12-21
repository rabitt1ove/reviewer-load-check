# reviewer-load-check
GitHub ID毎のPRレビュー負荷を数値で表示するWebアプリケーション(SPA)です。<br>
ブラウザで開き、`team`を選択して`Search`ボタンを押下することで動作します。<br>
選択した`team`はローカルセッションに保存されます。

# DEMO
![image_6483441](https://user-images.githubusercontent.com/45308877/208905985-3c528188-5efd-453c-834c-a7d123b02a45.JPG)

## 検索条件
- state:
  - ALL: 全てのPR
  - Approved: 承認済みのPR
  - None: 未承認のPR
- team:
  - GitHubのチーム
- created from:
  - PR作成開始日
- created to:
  - PR作成終了日
## 検索結果
- GitHub ID: 指定したチームに所属しているGitHub ID
- pr count: PR数
- additions: 追加数
- deletions: 削除数
- state: 入力のstateと同じ

# Requirement
- Go v1.19

# Installation
```bash
cd reviewer-load-check直下

go mod tidy
```

# Setup
## 1.Personal access token
### GitHubでPersonal access tokenを発行する
Select scopesは以下の通り
- repo：全てオン
- admin:org：read:orgのみオン
- 他は全てオフ

### ソースにPersonal access tokenを定義する
- ソース`reviewer-load-check/constants/github_const.go`の<br>
`GITHUB_TOKEN string = "<Personal access token>"`に定義

## 2.チーム設定
### GitHubにチームを追加する
- https://github.com/orgs/[オーナー名]/teams にチームを追加
### ソースにチームを定義する
- ソース`reviewer-load-check/assets/javascripts/application.js`の<br>
`teamOptions:`に定義を追加する

# Note
- 現在、PRレビュー負荷の計測対象となるのは、検索にヒットしたPRのうち直近100件までです
- `Personal access token`は外部流出させないよう注意してください
