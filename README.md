<p align="left">
    <a href="https://github.com/canonical/hydra/actions/workflows/ci.yaml"><img src="https://github.com/canonical/hydra/actions/workflows/ci.yaml/badge.svg?branch=canonical&event=push" alt="CI Tasks for Ory Hydra"></a>
    <a href="https://codecov.io/gh/canonical/hydra"><img src="https://codecov.io/gh/canonical/hydra/branch/canonical/graph/badge.svg?token=y4fVk2Of8a"/></a>
</p>

This is a fork of [Ory Hydra](https://github.com/ory/hydra). Ory Hydra is used
as the OAuth2/OIDC Server on the
[Canonical Identity Platform](https://charmhub.io/topics/canonical-identity-platform).
The reason for forking Hydra is that we needed to support the
[OAuth 2.0 Device Authorization Grant](https://datatracker.ietf.org/doc/html/rfc8628).
Some work was already done on upstream hydra, but the PRs were never merged. Our
implementation is heavily influnced by the work done on
https://github.com/ory/hydra/pull/3252 from:

- [supercairos](https://github.com/supercairos)
- [BuzzBumbleBee](https://github.com/BuzzBumbleBee)

We plan to keep this fork up to date with upstream Hydra and release oci-images
on https://github.com/canonical/hydra-rock/pkgs/container/hydra, until this work
is merged upstream. See the [wiki](https://github.com/canonical/hydra/wiki) for
more details on the implementation and how to try out the device flow.
