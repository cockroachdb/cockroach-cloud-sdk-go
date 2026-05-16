# CLAUDE.md

## GitHub Actions

- When writing or modifying GitHub Actions workflows, always look up the latest major version of each action before using it. Do not assume you know the current version.

## Changelog

Follow Keep a Changelog conventions when updating CHANGELOG.md.

**Unreleased Section Structure:**
- The `## [Unreleased]` section starts with the `## [Unreleased]` header and ends when the next `##` level header appears (typically a version like `## [x.y.z]`)
- It contains subsections: `### Added`, `### Changed`, `### Deprecated`, `### Removed`, `### Fixed`, `### Security`
- Subsections must appear in this exact order (top to bottom): Added, Changed, Deprecated, Removed, Fixed, Security
- There should be only ONE of each subsection type within `## [Unreleased]`

**Adding Entries:**
- Add all new entries under the `## [Unreleased]` section
- Choose the appropriate subsection based on the change type:
  - Added: new features
  - Changed: changes in existing functionality
  - Deprecated: soon-to-be removed features
  - Removed: now removed features
  - Fixed: bug fixes
  - Security: vulnerabilities
- If the subsection already exists: add your entry at the top of that subsection's list - DO NOT create a duplicate subsection header
- If the subsection doesn't exist: create it within the `## [Unreleased]` section, positioned according to the order above (Added, Changed, Deprecated, Removed, Fixed, Security)

**How to Write Changelog Entries:**
- Be brief. One sentence per entry.
- Write from the user's perspective. Describe what they can now do, not what code changed.
- Start each entry with a verb (e.g. Add, Update, Fix, Remove).
- Prefix breaking changes with "Breaking Change: ".
- Group related changes into one entry when they form a single feature (e.g. "Add support for folder management operations").
- Don't list individual fields, models, or implementation details unless they're the primary change.
- For OpenAPI spec syncs: mention the new operations and capabilities, not the sync process itself.

**Examples:**
- Good: `Add support for JWT issuer management operations (CreateJwtIssuer, GetJwtIssuer, ListJwtIssuers, UpdateJwtIssuer, DeleteJwtIssuer).`
- Bad: `Sync OpenAPI spec from managed-service PR #1234 and add JwtIssuer CRUD operations with Audience, IssuerUrl, JwksUri, and ClaimMap fields.`
