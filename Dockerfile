# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox:1.38.0-glibc@sha256:3ba030337caebbfc2232b22b1e435eb213b28e5844a34942c74555bf904a265a

ARG TARGETOS
ARG TARGETARCH
COPY dist/trunkver_${TARGETOS}_${TARGETARCH} /usr/local/bin/trunkver
COPY dist/trunkver_${TARGETOS}_${TARGETARCH}.cosign.bundle /usr/local/share/trunkver.cosign.bundle

ENTRYPOINT ["/usr/local/bin/trunkver"]
