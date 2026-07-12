# Tier A GitHub Actions audit

Date: 2026-07-12

Scope audited from the issue inventory:

- static/documentation/data/awesome-list candidates
- potentially active repositories that were explicitly marked to skip unless clearly static

Method:

1. Checked this repository locally: there are no `.github/workflows/*` files and no build/test manifests.
2. Queried GitHub Actions workflow inventory for each repository with GitHub MCP `list_workflows`.
3. For repositories that still had workflows, sampled `README.md` and workflow files to confirm whether they are static or active before recommending Tier A.

## Disabled successfully

- None in this session.

The available GitHub tooling in this session can read workflow state, but it cannot flip the repository-level Actions setting for other repositories. This repository also has no local workflow file that could be neutralized by PR.

## Already no Actions/workflows

These repositories returned `total_count: 0` from `list_workflows`, so no Tier A file change was required:

- `awesome-civic-tech`
- `awesome-italia-opensource`
- `awesome-italian-public-datasets`
- `awesome-legal-data`
- `awesome-legal-nlp`
- `awesome-legaltech`
- `awesome-mcp-servers`
- `awesome-mcp-servers-1`
- `awesome-nextjs`
- `awesome-open-data`
- `awesome-water-wastewater`
- `comuni-italiani-istat`
- `dl-19-2024-patente-crediti-cantieri`
- `dlgs-81-08-testo-unico`
- `haccp-italia-normativa-regionale`
- `italian-sanzioni-dlgs-81-08`
- `verifiche-periodiche-inail-attrezzature`

## Needs manual follow-up

These repositories are still good Tier A candidates, but they currently have active workflows and cannot be disabled from this repository alone:

| Repository | Current workflow state | Why it still fits Tier A | Recommended action |
| --- | --- | --- | --- |
| `123consulenza-open-data` | 1 active workflow: `Copilot cloud agent` (`dynamic/copilot-swe-agent/copilot`) | Static open-data/documentation repository; local tree contains no workflow files | Disable GitHub Actions at repository level in Settings if full Tier A is desired. There is no local workflow file here to edit. |
| `accordi-stato-regioni-sicurezza-lavoro` | 2 active workflows: `CI`, `pages-build-deployment` | Structured documentation/data repository with GitHub Pages output | Disable Actions at repository level, or remove automatic triggers from `.github/workflows/ci.yml` and review whether the Pages deployment should also be paused. |
| `dlgs-81-08-glossario` | 1 active workflow: `CI` | Dataset/package repository whose README describes a JSON glossary dataset | Disable Actions at repository level, or neutralize `.github/workflows/ci.yml` in that repository. |
| `italian-ateco-database` | 1 active workflow: `CI` | Dataset/package repository centered on static ATECO JSON data | Disable Actions at repository level, or neutralize `.github/workflows/ci.yml` in that repository. |
| `italian-province-regioni-dataset` | 1 active workflow: `CI` | Dataset/package repository centered on static province/region JSON data | Disable Actions at repository level, or neutralize `.github/workflows/ci.yml` in that repository. |

## Skipped (conservative)

These repositories were intentionally left out of Tier A automatic disable because they look like active applications, libraries, or tools. They should get a separate Tier B/Tier C review instead of a blanket shutdown.

| Repository | Workflow inventory | Reason skipped |
| --- | --- | --- |
| `chrome-extension-italian-workplace-safety` | 0 workflows | Active browser extension project per README. |
| `goccia` | 0 workflows | Listed in the issue as a potentially active repository; not clearly static from the name alone. |
| `lims` | 0 workflows | Listed in the issue as a potentially active repository; not clearly static from the name alone. |
| `mcp-italian-workplace-safety` | 0 workflows | Active MCP server per README. |
| `next-seo-italian-helpers` | 1 active workflow: `CI` | Active TypeScript/Next.js library per README. |
| `scadenze-formazione-calculator` | 1 active workflow: `CI` | Active TypeScript library with test usage documented in README. |
| `italian-inail-rapporto-annuale-parser` | 0 workflows | Parser/tool repository name indicates executable code, not a static dataset. |
| `apd-core` | 0 workflows | Core library/application repository name indicates active code. |
| `apis-guru-openapi-directory` | 0 workflows | OpenAPI directory/tooling repository name indicates active code or sync processes. |
| `schemastore` | 0 workflows | Schema/tooling repository name indicates active maintenance, not a passive static archive. |

## Notes

- This repository already had the smallest possible local state for Tier A: no checked-in workflow files, only issue templates under `.github/`.
- The remaining work is administrative or cross-repository:
  - disable Actions in repository settings for the Tier A candidates above, or
  - open per-repository PRs that remove/neutralize workflow triggers where repository settings cannot be changed.
