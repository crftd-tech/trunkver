VERSION ?= 0.0.0-HEAD-local
IMAGE := ghcr.io/rradczewski/trunkver

SPACE := $() $()
COMMA := ,

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
  SMOKE_BINARY := smoke-v2.4.0-Darwin-arm64
else ifeq ($(LOCAL_PLATFORM),linux_amd64)
  SMOKE_BINARY := smoke-v2.4.0-Linux-x86_64
endif

.PHONY: all
all: \
	validate \
	build \
	sign

.PHONY: validate
validate: \
	test \
	spec

.PHONY: build
build: \
	$(addprefix dist/trunkver_, $(PLATFORMS)) \
	$(addsuffix .sbom.json,$(addprefix dist/trunkver_, $(PLATFORMS))) \
	dist/checksums.txt

.PHONY: sign
sign: \
	dist/checksums.txt.cosign.bundle \
	$(addsuffix .sbom.json.cosign.bundle,$(addprefix dist/trunkver_, $(PLATFORMS))) \
	$(addsuffix .cosign.bundle,$(addprefix dist/trunkver_, $(PLATFORMS)))

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf dist/

README.md: website/index.md README.md.head
	cp README.md.head $@
	sed '1,/{% # CUT FOR README %}/d' website/index.md >> $@
	# Replace double escaped YAML substitution in GithubAction reference
	sed -i "s/'{{ \(.*\) }}'/ \1 /" $@

.PHONY: $(addprefix dist/trunkver_, $(PLATFORMS))
$(addprefix dist/trunkver_, $(PLATFORMS)): test
	@mkdir -p dist || true
	GOOS=$(word 2,$(subst _, ,$@)) \
	  GOARCH=$(word 3,$(subst _, ,$@)) \
	  go build -ldflags "-X github.com/crftd-tech/trunkver/internal.Version=$(VERSION)" -o $@ ./main.go

dist/%.cosign.bundle: dist/%
	cosign sign-blob --yes --bundle $@ $^

dist/%.sbom.json: dist/%
	syft scan --output spdx-json=$@ file:$<


dist/checksums.txt: $(addprefix dist/trunkver_, $(PLATFORMS))
	cd dist; sha256sum $(subst dist/,,$^) | tee checksums.txt

.PHONY: docker
docker: 
	for i in linux/arm64 linux/amd64; do \
		docker buildx build \
			--platform $$i \
			-f Dockerfile \
			-t ${IMAGE}:${VERSION} \
			.; \
	done

ext/${SMOKE_BINARY}:
	curl --location https://github.com/SamirTalwar/smoke/releases/download/v2.4.0/${SMOKE_BINARY} -o $@
	(cd ext; sha256sum --check ${SMOKE_BINARY}.sha256sum)
	chmod a+x $@

.PHONY: spec
spec: dist/trunkver_$(LOCAL_PLATFORM) ext/${SMOKE_BINARY}
	./ext/${SMOKE_BINARY} --command="./dist/trunkver_$(LOCAL_PLATFORM)" .
