all: prepare build upload

prepare:
	go get -d
	go get github.com/aktau/github-release

build:
	env GOOS=linux GOARCH=amd64 go build -o syslog_wrapper-linux-amd64

upload:
	github-release upload \
	    --user dz0ny \
	    --repo syslog_wrapper \
	    --tag v0.1.0 \
	    --name "syslog_wrapper-linux-amd64" \
	    --file syslog_wrapper-linux-amd64
