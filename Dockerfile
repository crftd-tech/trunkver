# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox:1.37.0-glibc@sha256:3bf024f5b91b256d55fcecaa910a7f671bdd2b6bb5bb22ac6b774cc4678f2093

ARG TARGETOS
ARG TARGETARCH
COPY dist/trunkver_${TARGETOS}_${TARGETARCH} /usr/local/bin/trunkver
COPY dist/trunkver_${TARGETOS}_${TARGETARCH}.cosign.bundle /usr/local/share/trunkver.cosign.bundle

ENTRYPOINT ["/usr/local/bin/trunkver"]
