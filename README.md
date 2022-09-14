1. 実行方法
コンパイル
```
GOOS=linux GOARCH=amd64 go build -o hanlder main.go
```
2. 圧縮
```
zip function.zip hanlder
```

3. ```function.zip```をAWSコンソール上でアップロード
