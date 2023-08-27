build:
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap

zip:
	zip -j bootstrap.zip bootstrap
