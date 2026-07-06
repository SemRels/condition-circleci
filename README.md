# condition-circleci

[![Latest Release](https://img.shields.io/github/v/release/SemRels/condition-circleci?label=version&color=blue)](https://github.com/SemRels/condition-circleci/releases/latest)

Allows releases only when semrel is running inside CircleCI.

This plugin is distributed as the standalone Go binary `semrel-plugin-condition-circleci`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/condition-circleci/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/condition-circleci:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/condition-circleci:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/condition-circleci/.github/workflows/release.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```

## Configuration

```yaml
plugins:
  - name: condition-circleci
    path: ~/.semrel/plugins/semrel-plugin-condition-circleci
    env:
      {}
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_BRANCH` | no | Optional branch name override. When set, require `CIRCLE_BRANCH` to match this value. | unset |

## `SEMREL_*` release context used

This plugin does not consume any `SEMREL_*` release context variables directly.

## Example behavior

The plugin checks the CI environment and succeeds when `CIRCLECI=true`. If `SEMREL_PLUGIN_BRANCH` is set, it also requires `CIRCLE_BRANCH` to match. Outside CircleCI it exits non-zero to stop the release.

## License

Apache-2.0
