prepare:
	rm -rf ./bin/downloads/*.json && \
	rm -rf ./bin/output/*.csv

dependencies:
	go mod download

build: dependencies
	rm -rf bin/file-parser && \
	go build -o bin/file-parser

run: build prepare
	./bin/file-parser

windows-build: dependencies
	rm -rf bin/dist/windows/ && \
	GOOS=windows GOARCH=amd64 go build -o bin/dist/windows/file-parser.exe

windows-run:
	./bin/dist/windows/file-parser.exe

mac-build: dependencies
	rm -rf bin/dist/mac/ && \
	GOOS=darwin GOARCH=amd64 go build -o bin/dist/mac/file-parser

mac-run: prepare
	./bin/dist/mac/file-parser

linux-build: dependencies
	rm -rf bin/dist/linux/ && \
	GOOS=linux GOARCH=amd64 go build -o bin/dist/linux/file-parser

linux-run: prepare
	./bin/dist/linux/file-parser

distribute: windows-build mac-build linux-build
