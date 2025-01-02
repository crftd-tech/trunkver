---
layout: layout
---

# <span class="logo">TrunkVer</span>

so we can stop talking about versions and start shipping.

<div class="spacer"></div>
{% include 'hero.html' %}

## TL;DR

`TrunkVer` is a SemVer-compatible versioning scheme for
continuously-delivered, trunk-based applications and systems that don't follow a release scheme.

It is a **drop-in replacement** for semantic versions and replaces the version with meaningful meta data, telling you at a glance what the artifact is, when it was built and where you may find the build log.

## Usage

### GitHub Actions

```yaml
- name: Generate trunkver
  id: trunkver
  uses: crftd-tech/trunkver@main

- name: Print trunkver
  env:
    TRUNKVER: ${{'{{ steps.trunkver.outputs.trunkver }}'}}
  run: |
    echo "$TRUNKVER"
```

### GitLab

#### Using our template from https://gitlab.com/crftd-tech/trunkver-gitlab-ci

```yaml
include:
- remote: 'https://gitlab.com/crftd-tech/trunkver-gitlab-ci/-/raw/main/trunkver.gitlab-ci.yml'
```

#### Downloading the CLI directly

```yaml
build:
  script:
    - curl -sSL https://github.com/crftd-tech/trunkver/releases/latest/download/trunkver_linux_amd64 -o trunkver
    - chmod +x trunkver
    - export TRUNKVER=$(./trunkver generate)
```

### Docker

```sh
docker run --rm ghcr.io/crftd-tech/trunkver:latest generate --build-ref "$CI_JOB_ID" --source-ref "g$(git rev-parse --short HEAD)"
```

### Other CIs

```sh
curl -sSL https://github.com/crftd-tech/trunkver/releases/latest/download/trunkver_linux_amd64 -o trunkver
chmod +x trunkver
./trunkver generate
```

## Rationale

We have identified a frequent source of avoidable confusion, conflict and cost in the software delivery process caused by versioning software that should not be versioned - or rather, the versioning should be automated.

Historically, countless versioning schemes have been used to signify changes to a piece of software, using a lot of not clearly defined words such as beta, final or release candidate as well as arbitrary numbering schemes that typically involve one or more digits that are incremented according to certain rules, or worse, without clear rules.

Over time, semantic versioning has been proposed and adopted by many developer teams for good reasons. It clearly defines what each part of a version number means, such as incrementing the first digit, or major version, to signify a change that is considered to be a “breaking” one. This can be used by both humans and machines to improve their work, such as a machine refusing to automatically apply an update in this case and notifying a human to adapt to the breaking change. We are fans of semantic versioning.

However, we keep encountering teams and organizations that apply semantic versioning or a custom versioning scheme to software that does not need any of that - and through this, they create an astonishing amount of unnecessary work such as arguing whether or not a certain piece of software should be called “alpha”, “beta”, “rho”, “really final v4” etc, manually creating tickets listing the changes or even specialized gatekeeper roles such as “release engineer” - in the worst case a single person in the whole organization. Because this makes it harder, boring and costly to deploy, it systematically reduces the number of deployments, and through this the delivery performance of the organization.

We acknowledge that these efforts often stem from perfectly valid requirements of various stakeholders, such as the necessity to audit the release process, finding out what version of the software is currently running or adhering to a specific certified process. Ironically, the manual process around it makes this not only costly, but often defeats the intended purpose. We have seen audit trails missing commits, full of copy/paste errors, etc. We therefore argue the only way to get auditing right is by automating it too, including the version numbering.

In an organization that creates software in teams of trusted contributors that deploy software to controlled environments together, this kind of versioning ceremony can become a major hurdle to adopting the XP and DevOps practices of trunk-based development and continuous integration/deployment/delivery - and it can and should be replaced by a tiny amount of code.

## Principles

- The TrunkVer is automatically created during the build process, typically as the first step, and then re-used across the build, e.g. to tag created images in a registry or as part of a file name.
- There is a single source of truth for the current code - usually a git branch named `trunk`, `default` or `main`.
- From the source, we create a deployable artifact of the software - e.g. a docker image, jar file or debian package. As described in the 12 factor principles, this artifact will be deployed to an environment by combining it with configuration not part of the artifact. Versioning deployments is tool-specific and out of scope
- There are only release candidates. Every push to the trunk triggers an automated process that creates a potentially releasable artifact, and usually it is automatically released shortly after
- Version numbers are useful to developers: While technically a build jobs id would be perfectly viable to use, we prepend a timestamp to make them sortable, and the source code revision to identify it without having to go through the build server.
- All secondary artifacts of a release, such as tickets or a changelog, are automated created during the build, referencing the TrunkVer.

## Specification

The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).

1. A TrunkVer is suitable for versioning of any artifact that may be released without any regard for compatibility with third parties integrating with the artifact.

2. A TrunkVer may be used for an artifact, while SemVer or other versioning schemes may be used for specific interfaces to that artifact, e.g. REST APIs.

3. A TrunkVer is syntactically compatible with SemVer, although it does not respect its semantic interpretation of the version number.

4. A TrunkVer consists of three components: A **timestamp**, a **source reference**, and **a build reference**.

5. The **timestamp** is precise to one second and always formatted in UTC i.e. `YYYYMMDDHHMMSS`, e.g. `20241230142105`. It replaces the **Major** part of a SemVer.

6. The **source reference** identifies the exact source code used to build the artifact. Identifiers MUST comprise only ASCII alphanumerics `[0-9A-Za-z]`. Identifiers MUST NOT be empty. If the source reference is a git commit checksum, it may be truncated  (e.g. by using `git rev-parse --short HEAD`) and be prefixed with a `g` (as customary when using `git describe`).

7. The **build reference** identifies the exact build job used to build the artifact. Identifiers MUST comprise only ASCII alphanumerics and hyphens `[0-9A-Za-z-]`.

8. A TrunkVer is formatted as follows: The **timestamp**, followed by `.0.0-`, followed by the **source reference**, followed by `-`, followed by the **build reference**. It may be used to replace any SemVer version.

9. Alternatively, for the purpose of e.g. prereleases of an otherwise SemVer-versioned artifact, a TrunkVer may be assembled into the prerelease part of a SemVer version as follows: The **timestamp**, followed by `-`, followed by the **source reference**, followed by `-`, followed by the **build reference**. 

### EBNF Definition

```ebnf
TRUNKVER = ( MAJOR_TRUNKVER | PRERELEASE_TRUNKVER );
MAJOR_TRUNKVER = TIMESTAMP, '.0.0-', SOURCE_REF, '-', BUILD_REF;
PRERELEASE_TRUNKVER = TIMESTAMP, '-', SOURCE_REF, '-', BUILD_REF;

TIMESTAMP = NON_ZERO_DIGIT, 11*DIGIT;
BUILD_REF = { ALPHANUMERIC | '-' };
SOURCE_REF = { ALPHANUMERIC };

DIGIT = "0" | NON_ZERO_DIGIT;
NON_ZERO_DIGIT = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9";
HEXADECIMAL = "a" | "b" | "c" | "d" | "e" | "f" | "A" | "B" | "C" | "D" | "E" | "F" | "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9";
ALPHANUMERIC = DIGIT | "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z" | "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T" | "U" | "V" | "W" | "X" | "Y" | "Z";

```

## FAQ

- **Why only use the MAJOR SemVer part as a timestamp?**  
  Tools (e.g. compliance, audit trails) might still classify a version based on SemVer semantics, hence TrunkVer always defensively implies breaking changes between versions.

- **Why do you use the PRERELEASE SemVer part for source information and not BUILD?**  
  Because OCI tags don't support `+` (see [distribution/distribution#1201](https://github.com/distribution/distribution/issues/1201) and [opencontainers/distribution-spec#154](https://github.com/opencontainers/distribution-spec/issues/154). We'd rather have one consistent version across artifacts. Semantically, the only relevant portion for sorting of a TrunkVer is
  the MAJOR version, and a conflict (as in creating two artifacts in the
  very same second) should be avoided.

## Links

- [https://github.com/crftd-tech/trunkver](https://github.com/crftd-tech/trunkver)
- [https://crftd.tech/](https://crftd.tech/)
- [https://semver.org](https://semver.org)
