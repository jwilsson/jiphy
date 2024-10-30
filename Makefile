build:
	env GOOS=linux GOARCH=arm64 go build -C app -tags lambda.norpc -o build/bootstrap
