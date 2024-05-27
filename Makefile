VERSION ?= 0.0.0-HEAD-local

PLATFORMS := darwin_arm64 darwin_amd64 linux_amd64 linux_arm64 windows_amd64

ifeq ($(OS),Windows_NT)
  LOCAL_PLATFORM := windows_amd64
else
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Darwin)
	LOCAL_PLATFORM := darwin_amd64
  else
	LOCAL_PLATFORM := linux_amd64
  endif
endif

ifeq ($(LOCAL_PLATFORM),darwin_amd64)
  SMOKE_PLATFORM := Darwin-arm64
else ifeq ($(LOCAL_PLATFORM),linux_amd64)
  SMOKE_PLATFORM := Linux-x86_64
endif


.PHONY: all
all: test spec $(addprefix out/trunkver_, $(PLATFORMS))

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf out

out:
	@mkdir -p out || true

.PHONY: $(addprefix out/trunkver_, $(PLATFORMS))
$(addprefix out/trunkver_, $(PLATFORMS)): out test
	GOOS=$(word 2,$(subst _, ,$@)) \
	  GOARCH=$(word 3,$(subst _, ,$@)) \
	  go build -ldflags "-X github.com/crftd-tech/trunkver/cmd.Version=$(VERSION)" -o $@ ./main.go

out/smoke: out
	curl -sL https://github.com/SamirTalwar/smoke/releases/download/v2.4.0/smoke-v2.4.0-$(SMOKE_PLATFORM) -o $@
	chmod a+x $@

.PHONY: spec
spec: out/trunkver_$(LOCAL_PLATFORM) out/smoke
	./out/smoke --command="./out/trunkver_$(LOCAL_PLATFORM) --timestamp 2024-05-22T16:25:48+02:00" .
