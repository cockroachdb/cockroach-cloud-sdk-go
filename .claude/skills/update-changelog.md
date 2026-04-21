I need you to analyze the code changes in this SDK repository and update the CHANGELOG.md file.

Changes were made by syncing the OpenAPI spec from managed-service.

## Environment Variables

The following environment variables are set. Read each one individually
using `printenv <VAR>` (one variable per command):

- MANAGED_SERVICE_PR_URL — The managed-service PR URL that triggered this sync
- MANAGED_SERVICE_SHA — (optional) The managed-service commit SHA if the PR was merged

## Instructions

Please follow these steps:

1. Use git diff to see all changes in the working tree
2. Read the current CHANGELOG.md file to understand its existing structure
3. Update CHANGELOG.md following Keep a Changelog conventions (https://keepachangelog.com/):
   - Add all new entries under the [Unreleased] section
   - For each change, choose the appropriate subsection:
     - Added: for new features
     - Changed: for changes in existing functionality
     - Deprecated: for soon-to-be removed features
     - Removed: for now removed features
     - Fixed: for any bug fixes
     - Security: in case of vulnerabilities
   - If a subsection header (e.g., ### Added) does not exist under [Unreleased], create it
   - If the subsection header already exists, add new entries at the top of that subsection
   - Maintain the subsection order: Added, Changed, Deprecated, Removed, Fixed, Security
4. For each change, determine if it's a breaking change for SDK consumers
5. If a change is breaking, prefix that specific entry with "Breaking Change: "
6. Write entries as bullet points using past tense (e.g., "Added X field", not "Add X field")

IMPORTANT: Only modify CHANGELOG.md. Do not modify any other files including generated code, spec files, or configuration files.

After updating CHANGELOG.md, you MUST output structured metadata in the following format.
This is critical - end your response with these exact markers:

---COMMIT_TITLE---
<one line commit title - keep it under 72 characters>
---COMMIT_BODY---
<multi-line commit body explaining the changes - wrap lines at 72 characters>

Managed-service-pr-url: $MANAGED_SERVICE_PR_URL
---PR_TITLE---
<one line PR title - keep it under 72 characters>
---PR_DESCRIPTION---
<multi-line PR description with summary and details>

IMPORTANT GUIDANCE FOR COMMIT AND PR TITLES:
- Describe the actual SDK/API changes from a user perspective
- NEVER write titles like "Update CHANGELOG.md" - the changelog update is just a side effect, not the main change
- If changes are focused on one area, describe that (e.g., "Add MFA audit log actions", "Update invoice field descriptions")
- If there are multiple unrelated changes that can't fit in 72 characters, use "Sync OpenAPI spec from managed-service"
- Focus on WHAT changed in the API/SDK, not the process of updating files

The commit body MUST include the trailer "Managed-service-pr-url: $MANAGED_SERVICE_PR_URL" exactly as shown.
Do not add any text after the PR description section.
