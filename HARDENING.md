<!-- markdownlint-disable -->

# Hardening Report: securego--gosec/v2.29.0

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **securego--gosec/v2.29.0** was hardened automatically. 2 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

ci.yml contains two unpinned `uses:` references that point to mutable branch names instead of full 40-character commit SHAs. `securego/gosec@master` (line 53) and `ccojocar/barry@main` (line 102) can be silently updated by the upstream repository owner, enabling supply-chain attacks. Both are annotated with `# zizmor: ignore[unpinned-uses]` but remain genuine security risks.

Locations:

- `.github/workflows/ci.yml:53`
- `.github/workflows/ci.yml:102`

### script-injection (severity: high)

Rule (b) violation in release.yml 'Sign Docker Image' step: the shell variables `${TAGS}` (line 81) and `${images}` (line 84) are expanded without double-quoting inside the `run:` script. `TAGS` is sourced from `${{steps.meta.outputs.tags}}` (a `steps.*.outputs.*` context, which is workflow-controllable). An unquoted expansion allows the shell to parse metacharacters (`;`, `|`, `&`, whitespace, glob chars) out of the value, enabling command injection. The `for tag in ${TAGS}` loop and `cosign sign --yes --key /tmp/cosign.key ${images}` should use quoted forms: `for tag in "${TAGS}"` and `cosign sign ... "${images}"`.

Locations:

- `.github/workflows/release.yml:81`
- `.github/workflows/release.yml:84`

## Iteration Notes

### Iteration 1

**Fixes applied:** unpinned-uses, script-injection

**Notes:**

Fixed three security issues: (1) Pinned `securego/gosec@master` to commit SHA `57f161386e11e13e4c8a37255bc1a6502993d079` in ci.yml line 53. (2) Pinned `ccojocar/barry@main` to commit SHA `66654985f3a1b5fe417656467fa0d0b1e9f6e1db` in ci.yml line 102. Both `# zizmor: ignore` annotations were removed since the actions are now properly pinned. (3) Fixed script injection in release.yml 'Sign Docker Image' step: replaced the unquoted `for tag in ${TAGS}` loop and `cosign sign ... ${images}` with a safe `while IFS= read -r tag` loop that builds a bash array `image_refs=()` and passes it quoted as `"${image_refs[@]}"` to cosign, preventing shell metacharacter injection from the TAGS value.

