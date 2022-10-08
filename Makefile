# golangci.ymlを指定してlintを実行する
.PHONY: lint
lint:
	golangci-lint run