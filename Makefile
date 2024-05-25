
VERSION ?= 0.0.0-HEAD-local

.PHONY: test
test:
	go test -v 


trunkver_linux_amd64: trunkver.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o trunkver_linux_amd64

smoke:
	curl -sL https://github.com/SamirTalwar/smoke/releases/download/v2.4.0/smoke-v2.4.0-Linux-x86_64 -o smoke
	chmod a+x smoke

.PHONY: spec
spec: trunkver_linux_amd64 smoke
	./smoke .
