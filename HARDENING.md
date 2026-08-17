<!-- markdownlint-disable -->

# Hardening Report: securego--gosec/v2.28.0

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **securego--gosec/v2.28.0** was hardened automatically. 2 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

Two `uses:` references in ci.yml use mutable branch refs instead of pinned 40-character commit SHAs, making the workflow vulnerable to supply-chain attacks if those branches are compromised:
- Line 43: `uses: securego/gosec@master` — branch ref
- Line 68: `uses: ccojocar/barry@main` — branch ref
Both should be replaced with full SHA digests (e.g. `uses: securego/gosec@<40-hex-sha> # vX.Y.Z`).

Locations:

- `.github/workflows/ci.yml:43`
- `.github/workflows/ci.yml:68`

### script-injection (severity: high)

Sub-rule (b) violation in the 'Sign Docker Image' step of release.yml: the env var `TAGS` is sourced from `steps.meta.outputs.tags` (a workflow-controllable context value) and `DIGEST` from `steps.relimage.outputs.digest`, but both are expanded unquoted inside the `run:` shell script. Specifically:
- `for tag in ${TAGS}; do` — unquoted expansion of workflow-controlled value
- `cosign sign --yes --key /tmp/cosign.key ${images}` — `${images}` is built from unquoted `${TAGS}` and `${DIGEST}`
An attacker who can influence the tag metadata output could inject shell metacharacters. All expansions should be double-quoted: `for tag in "${TAGS}"; do` and `cosign sign --yes --key /tmp/cosign.key "${images}"`.

Locations:

- `.github/workflows/release.yml:71`
- `.github/workflows/release.yml:74`

## Iteration Notes

### Iteration 1

**Fixes applied:** unpinned-uses, script-injection

**Notes:**

Fixed three security issues: (1) Pinned securego/gosec@master to full SHA 2eb4727ebc7504944ddbcdd2f7157db492026c3f in ci.yml; (2) Pinned ccojocar/barry@main to full SHA 66654985f3a1b5fe417656467fa0d0b1e9f6e1db in ci.yml; (3) Fixed script injection in release.yml's 'Sign Docker Image' step by replacing the unquoted `for tag in ${TAGS}` loop and unquoted `${images}` expansion with a safe `while IFS= read -r tag` loop that builds a bash array and passes it as `"${images[@]}"` to cosign sign.

