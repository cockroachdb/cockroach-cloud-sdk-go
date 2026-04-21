# CLAUDE.md

## GitHub Actions

- When writing or modifying GitHub Actions workflows, always look up the latest major version of each action before using it. Do not assume you know the current version.

## Versioning

CHANGELOG.md is the single source of truth for the version. When committing changes, add a user-facing entry under the `## [Unreleased]` section describing what changed. Follow the [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) format with these subsections: `### Added`, `### Changed`, `### Deprecated`, `### Removed`, `### Fixed`, `### Security`. Prefix breaking changes with "Breaking Change: ".

When releasing a new version, move `[Unreleased]` entries in CHANGELOG.md to a new `[x.y.z]` section.

### Writing good changelog entries

- Be brief. One sentence per entry.
- Write from the user's perspective. Describe what they can now do, not what code changed.
- Start each entry with a verb (e.g. Add, Update, Fix, Remove).
- Group related changes into one entry when they form a single feature (e.g. "Add support for folder management operations").
- Don't list individual fields, models, or implementation details unless they're the primary change.
- For OpenAPI spec syncs, mention the new operations and capabilities, not the sync process itself. The managed-service PR reference belongs in the commit message,
not the changelog.

Good: `Add support for JWT issuer management operations (CreateJwtIssuer, GetJwtIssuer, ListJwtIssuers, UpdateJwtIssuer, DeleteJwtIssuer).`
Bad: `Sync OpenAPI spec from managed-service PR #1234 and add JwtIssuer CRUD operations with Audience, IssuerUrl, JwksUri, and ClaimMap fields.`
