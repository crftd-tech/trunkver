# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox:1.37.0-glibc@sha256:3f9777e7e82e8591542f72b965ec7db7e8b3bdb59692976af1bb9b2850b05a4e

ARG TARGETOS
ARG TARGETARCH
COPY dist/trunkver_${TARGETOS}_${TARGETARCH} /usr/local/bin/trunkver
COPY dist/trunkver_${TARGETOS}_${TARGETARCH}.cosign.bundle /usr/local/share/trunkver.cosign.bundle

ENTRYPOINT ["/usr/local/bin/trunkver"]
