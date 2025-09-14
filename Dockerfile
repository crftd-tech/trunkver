# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox:1.37.0-glibc@sha256:a2c55ed708c564a69a695e0a3bb16a4c47d2bb268d2ebd06f0d77336801b80de

ARG TARGETOS
ARG TARGETARCH
COPY dist/trunkver_${TARGETOS}_${TARGETARCH} /usr/local/bin/trunkver
COPY dist/trunkver_${TARGETOS}_${TARGETARCH}.cosign.bundle /usr/local/share/trunkver.cosign.bundle

ENTRYPOINT ["/usr/local/bin/trunkver"]
