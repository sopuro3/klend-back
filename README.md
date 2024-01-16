# klend-back

golangci-lintで使うlinterの絞り込みをしていないので、厳しいと感じたら都度いじる

.envか環境変数に

```
POSTGRES_USER=ユーザー名
POSTGRES_PASSWORD=パスワード
POSTGRES_DB=DBの名前
DEVELOP=1 #定義したらSQLのデータを流しこむ
DOCKER_COMPOSE=0 #0の場合ローカルでklend-backを動かす
```

を追加してから起動してくれ
