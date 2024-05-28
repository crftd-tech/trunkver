---
layout: layout
---

# <span class="logo">trunkVer</span>

so we can stop talking about versions and start shipping.

<div class="spacer"></div>
{% include 'hero.html' %}

## TL;DR

`trunkVer` is a semVer-compatible versioning scheme for
trunk-based applications and systems that don't follow a release scheme.

It removes the chore of manually bumping version numbers and instead
enriches the version number with three important data points: The
**when**, **what** and **how**.

## Usage

### GitHub Actions

```yaml
- name: Generate trunkver
  id: trunkver
  uses: crftd-tech/trunkver@main

- name: Print trunkver
  env:
  TRUNKVER: ${{ steps.trunkver.outputs.trunkver }}
  run: |
    echo "$TRUNKVER"
```

### Docker

```sh
docker run ghcr.io/crftd-tech/trunkver:latest --build-ref "$CI_JOB_ID" --source-ref "g$(git rev-parse --short HEAD)"
```

### Other CIs

```sh
curl -sSL https://github.com/crftd-tech/trunkver/releases/latest/download/trunkver_linux_amd64 -o trunkver
chmod +x trunkver
./trunkver
```

## FAQ

- **Why only use the MAJOR semVer part as a timestamp?**  
  Tools (e.g. compliance, audit trails) might still classify a version based on semVer semantics, hence trunkVer always defensively implies breaking changes between versions.

- **Why do you use the PRERELEASE semVer part for source information and not build?**  
  Because OCI tags don't support `+` (see [distribution/distribution#1201](https://github.com/distribution/distribution/issues/1201) and [opencontainers/distribution-spec#154](https://github.com/opencontainers/distribution-spec/issues/154). We'd rather have one consistent version across artifacts. Semantically, the only relevant portion for sorting of a trunkVer is
  the MAJOR version, and a conflict (as in creating two artifacts in the
  very same second) should be avoided.

## Links

- [https://github.com/crftd-tech/trunkver](https://github.com/crftd-tech/trunkver)
- [https://crftd.tech/](https://crftd.tech/)
