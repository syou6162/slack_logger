# slack_logger
コマンドラインからslackを眺めるための小さなスクリプトです。slackのchannels.historyのAPIにユーザー名やチーム名など必要そうな情報を付与します。付与した情報をjqで加工し、pecoで必要なチャネルだけをフィルタリングしながら見ると快適です。

# Usage
まず、APIを使うためのtokenを取得します。
- [Slack Web API | Slack](https://api.slack.com/web)

次に必要となるツールをインストールします。

```
brew install jq
brew install coreutils
brew install peco
```

最後に以下のコマンドを実行します。token1などは自分で取得したtokenを入れてください。

```bash
while true
do
  gtimeout --foreground 300 bash -c "./slack token1 token2 token3 | peco"
  sleep 5.0;
done
```
