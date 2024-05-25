
VERSION ?= 0.0.0-HEAD-local

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf out

out:
	@mkdir -p out || true

PLATFORMS := darwin_arm64 darwin_amd64 linux_amd64 linux_arm64 windows_amd64

.PHONY: all
all: $(addprefix out/trunkver_, $(PLATFORMS))

.PHONY: $(addprefix out/trunkver_, $(PLATFORMS))
$(addprefix out/trunkver_, $(PLATFORMS)): out
	GOOS=$(word 2,$(subst _, ,$@)) \
	  GOARCH=$(word 3,$(subst _, ,$@)) \
	  go build -ldflags "-X github.com/crftd-tech/trunkver/cmd.Version=$(VERSION)" -o $@ ./main.go

out/smoke: out
	curl -sL https://github.com/SamirTalwar/smoke/releases/download/v2.4.0/smoke-v2.4.0-Linux-x86_64 -o $@
	chmod a+x $@

.PHONY: spec
spec: out/trunkver_linux_amd64 out/smoke
	./out/smoke .
