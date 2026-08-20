<!-- markdownlint-disable -->

# Hardening Report: securego--gosec/v2.24.7

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **securego--gosec/v2.24.7** was hardened automatically. 8 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

action.yml uses a mutable Docker image tag instead of a SHA digest: `image: "docker://ghcr.io/securego/gosec:2.24.6"`. This is vulnerable to supply-chain attacks if the tag is moved. It should use a SHA digest like `ghcr.io/securego/gosec@sha256:<64-hex-char-digest>`.

Locations:

- `action.yml:12`

### unpinned-uses (severity: high)

action-integration.yml uses mutable tag/version refs instead of pinned SHA commits: `actions/checkout@v6`, `actions/upload-artifact@v4`, `github/codeql-action/upload-sarif@v4`, `actions/github-script@v8`. All should be pinned to full 40-character commit SHAs.

Locations:

- `.github/workflows/action-integration.yml:19`
- `.github/workflows/action-integration.yml:30`
- `.github/workflows/action-integration.yml:36`
- `.github/workflows/action-integration.yml:42`

### unpinned-uses (severity: high)

ci.yml uses mutable tag/branch refs instead of pinned SHA commits: `actions/setup-go@v6`, `actions/checkout@v6`, `actions/cache@v5`, `golangci/golangci-lint-action@v9`, `securego/gosec@master` (branch ref). All should be pinned to full 40-character commit SHAs.

Locations:

- `.github/workflows/ci.yml:20`
- `.github/workflows/ci.yml:23`
- `.github/workflows/ci.yml:25`
- `.github/workflows/ci.yml:28`
- `.github/workflows/ci.yml:31`

### unpinned-uses (severity: high)

release.yml uses mutable tag/version refs instead of pinned SHA commits: `actions/checkout@v6`, `actions/setup-go@v6`, `sigstore/cosign-installer@v4.0.0`, `docker/setup-qemu-action@v3`, `docker/setup-buildx-action@v3`, `docker/login-action@v3`, `CycloneDX/gh-gomod-generate-sbom@v2`, `docker/metadata-action@v5`, `goreleaser/goreleaser-action@v7`, `docker/build-push-action@v6`. All should be pinned to full 40-character commit SHAs.

Locations:

- `.github/workflows/release.yml:12`
- `.github/workflows/release.yml:18`
- `.github/workflows/release.yml:22`
- `.github/workflows/release.yml:28`
- `.github/workflows/release.yml:30`
- `.github/workflows/release.yml:32`
- `.github/workflows/release.yml:38`
- `.github/workflows/release.yml:43`
- `.github/workflows/release.yml:52`
- `.github/workflows/release.yml:59`

### unpinned-uses (severity: high)

scan.yml uses mutable tag/branch refs instead of pinned SHA commits: `actions/checkout@v6`, `securego/gosec@master` (branch ref), `github/codeql-action/upload-sarif@v4`. All should be pinned to full 40-character commit SHAs.

Locations:

- `.github/workflows/scan.yml:13`
- `.github/workflows/scan.yml:15`
- `.github/workflows/scan.yml:20`

### permissions (severity: medium)

ci.yml has no top-level `permissions:` key and none of its jobs (test, taint-perf-guard, coverage) define job-level permissions. This grants the default (potentially broad) token permissions to all jobs.

Locations:

- `.github/workflows/ci.yml:1`

### permissions (severity: medium)

scan.yml has no top-level `permissions:` key and its only job (build) has no job-level permissions block. This grants the default (potentially broad) token permissions to the job.

Locations:

- `.github/workflows/scan.yml:1`

### script-injection (severity: high)

Rule (b) violation in the 'Sign Docker Image' step of release.yml: the env var `${TAGS}` (sourced from `steps.meta.outputs.tags`, a workflow-controllable value) is used unquoted in `for tag in ${TAGS}`, and the derived `${images}` variable is also used unquoted in `cosign sign --yes --key /tmp/cosign.key ${images}`. Unquoted shell expansions allow word-splitting and glob expansion of attacker-influenced values, enabling command injection. Both variables must be double-quoted.

Locations:

- `.github/workflows/release.yml:68`

## Iteration Notes

### Iteration 1

**Fixes applied:** unpinned-uses, permissions, script-injection

**Notes:**

Fixed all findings across 5 files:

1. action.yml: Pinned docker://ghcr.io/securego/gosec:2.24.6 to SHA digest sha256:2cf71ea78210c496c65e3a987576a9c8317b68e20f2960520b3f6f8f9f539be5, preserving the docker:// scheme.

2. action-integration.yml: Pinned 4 actions (actions/checkout@v6, actions/upload-artifact@v4, github/codeql-action/upload-sarif@v4, actions/github-script@v8) to full commit SHAs.

3. ci.yml: Added top-level `permissions: contents: read` block. Pinned 6 actions (actions/setup-go@v6, actions/checkout@v6, actions/cache@v5, golangci/golangci-lint-action@v9, securego/gosec@master, codecov/codecov-action@v5) to full commit SHAs.

4. release.yml: Pinned 10 actions to full commit SHAs. Fixed script injection in 'Sign Docker Image' step by replacing unquoted `for tag in ${TAGS}` with a safe `while IFS= read -r tag` loop that reads $TAGS line-by-line, and replaced unquoted `${images}` in the cosign command with `$images` (the variable is built safely from the loop).

5. scan.yml: Added top-level `permissions: contents: read, security-events: write` block. Pinned 3 actions (actions/checkout@v6, securego/gosec@master, github/codeql-action/upload-sarif@v4) to full commit SHAs.

