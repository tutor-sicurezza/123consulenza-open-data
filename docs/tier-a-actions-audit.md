# GitHub Actions cost-optimization — Tier A / B / C playbook

Date: 2026-07-12 (updated 2026-07-12)

Goal: eliminate unnecessary GitHub Actions spend across the `tutor-sicurezza` account by:

- **Tier A** — static/data/docs repos: disable Actions entirely (repo settings) or neutralize file triggers.
- **Tier B** — active library repos: keep only essential CI, add `concurrency`/`cancel-in-progress`, restrict to PR + default branch only.
- **Tier C** — no cron jobs were found in any repo during this audit.

Audit method:

1. `list_workflows` via GitHub API for every candidate repository.
2. Sampled `README.md` and workflow YAML for repos with active workflows.
3. Cross-checked against the "skip" list (active apps/tools).

---

## Tier A — Static repos: already clean (no workflows)

These 17 repositories returned `total_count: 0` from `list_workflows`. Nothing to do.

`awesome-civic-tech`, `awesome-italia-opensource`, `awesome-italian-public-datasets`,
`awesome-legal-data`, `awesome-legal-nlp`, `awesome-legaltech`, `awesome-mcp-servers`,
`awesome-mcp-servers-1`, `awesome-nextjs`, `awesome-open-data`, `awesome-water-wastewater`,
`comuni-italiani-istat`, `dl-19-2024-patente-crediti-cantieri`, `dlgs-81-08-testo-unico`,
`haccp-italia-normativa-regionale`, `italian-sanzioni-dlgs-81-08`,
`verifiche-periodiche-inail-attrezzature`

---

## Tier A — Static repos: need workflow neutralization

Four static dataset/docs repos still have an active `ci.yml` that runs on every `push` and `pull_request`. The fix is the same for all four: **replace automatic triggers with `workflow_dispatch` only**.

> **Quickest path**: go to each repo → Settings → Actions → General → set "Disable actions" and save.
> No PR needed. The YAML patches below are the fallback for when you prefer to keep the file visible but inactive.

### `accordi-stato-regioni-sicurezza-lavoro`

Current triggers: `push` (main/master) + `pull_request`. Also has a `pages-build-deployment` dynamic workflow (cannot be removed by file, tied to Pages setting).

**Action**: disable in Settings → Actions → General (preferred), or apply the patch below.
If you want to keep Pages working, only disable the `CI` workflow — the `pages-build-deployment` dynamic workflow will also stop running if Actions are fully disabled.

```yaml
# .github/workflows/ci.yml — NEUTRALIZED (workflow_dispatch only)
name: CI

on:
  workflow_dispatch:  # manual only — automatic push/PR triggers removed (Tier A cost-saving)

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  build-validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - name: Install
        run: |
          if [ -f package-lock.json ]; then npm ci --no-audit --no-fund; else npm install --no-audit --no-fund; fi
      - name: Build (if script exists)
        run: |
          if jq -e '.scripts.build' package.json > /dev/null; then npm run build; fi
      - name: Test (if script exists)
        run: |
          if jq -e '.scripts.test' package.json > /dev/null; then npm test; fi
      - name: Validate JSON metadata
        run: |
          find data -name "*.json" -print0 | xargs -0 -I {} sh -c 'jq empty "{}" && echo "ok: {}"'
```

### `dlgs-81-08-glossario`

Current triggers: `push` (main/master) + `pull_request`.

```yaml
# .github/workflows/ci.yml — NEUTRALIZED
name: CI

on:
  workflow_dispatch:  # manual only — Tier A cost-saving

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - name: Install (if package-lock exists)
        run: |
          if [ -f package-lock.json ]; then npm ci --no-audit --no-fund; fi
      - name: Validate JSON datasets
        run: |
          find data -name "*.json" -print0 | xargs -0 -I {} sh -c 'jq empty "{}" && echo "ok: {}"'
      - name: Smoke import
        run: node -e "const d = require('./index.js'); if (!d) process.exit(1); console.log('module loads ok');"
```

### `italian-ateco-database`

Current triggers: `push` (main/master) + `pull_request`.

```yaml
# .github/workflows/ci.yml — NEUTRALIZED
name: CI

on:
  workflow_dispatch:  # manual only — Tier A cost-saving

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - name: Install (if package-lock exists)
        run: |
          if [ -f package-lock.json ]; then npm ci --no-audit --no-fund; fi
      - name: Validate JSON datasets
        run: |
          find data -name "*.json" -print0 | xargs -0 -I {} sh -c 'jq empty "{}" && echo "ok: {}"'
      - name: Smoke import
        run: node -e "const d = require('./index.js'); if (!d) process.exit(1); console.log('module loads ok');"
```

### `italian-province-regioni-dataset`

Current triggers: `push` (main/master) + `pull_request`.

```yaml
# .github/workflows/ci.yml — NEUTRALIZED
name: CI

on:
  workflow_dispatch:  # manual only — Tier A cost-saving

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  build-validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - name: Install
        run: |
          if [ -f package-lock.json ]; then npm ci --no-audit --no-fund; else npm install --no-audit --no-fund; fi
      - name: Build (if script exists)
        run: |
          if jq -e '.scripts.build' package.json > /dev/null; then npm run build; fi
      - name: Test (if script exists)
        run: |
          if jq -e '.scripts.test' package.json > /dev/null; then npm test; fi
      - name: Validate JSON datasets
        run: |
          find data -name "*.json" -print0 | xargs -0 -I {} sh -c 'jq empty "{}" && echo "ok: {}"'
```

### `123consulenza-open-data` (this repository)

No checked-in workflow files. The only active workflow is a dynamic `Copilot cloud agent` workflow — this is not a file-based workflow and cannot be neutralized by a PR. To fully disable it go to: **Settings → Actions → General → Disable actions**.

---

## Tier B — Active library repos: hardened CI

These repos have real build/test workflows that should keep running, but only on PRs and pushes to the default branch — with `concurrency` added to cancel stale runs.

No other changes are needed (no matrix, no cron, no expensive extra jobs were found).

### `next-seo-italian-helpers`

Current: `push` (main/master) + `pull_request`. **Apply**: add `concurrency` + `cancel-in-progress`.

```yaml
# .github/workflows/ci.yml — TIER B (concurrency added)
name: CI

on:
  push:
    branches: [main, master]
    paths-ignore:
      - '**.md'
      - 'docs/**'
  pull_request:
    paths-ignore:
      - '**.md'
      - 'docs/**'

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  build-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - name: Install
        run: |
          if [ -f package-lock.json ]; then npm ci --no-audit --no-fund; else npm install --no-audit --no-fund; fi
      - name: Build (if script exists)
        run: |
          if jq -e '.scripts.build' package.json > /dev/null; then npm run build; fi
      - name: Test (if script exists)
        run: |
          if jq -e '.scripts.test' package.json > /dev/null; then npm test; fi
```

### `scadenze-formazione-calculator`

Current: `push` (main/master) + `pull_request`. **Apply**: add `concurrency` + `cancel-in-progress`.

```yaml
# .github/workflows/ci.yml — TIER B (concurrency added)
name: CI

on:
  push:
    branches: [main, master]
    paths-ignore:
      - '**.md'
      - 'docs/**'
  pull_request:
    paths-ignore:
      - '**.md'
      - 'docs/**'

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  build-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - name: Install
        run: |
          if [ -f package-lock.json ]; then npm ci --no-audit --no-fund; else npm install --no-audit --no-fund; fi
      - name: Build (if script exists)
        run: |
          if jq -e '.scripts.build' package.json > /dev/null; then npm run build; fi
      - name: Test (if script exists)
        run: |
          if jq -e '.scripts.test' package.json > /dev/null; then npm test; fi
```

---

## Tier C — Cron jobs

No scheduled (`schedule:`) triggers were found in any repository during this audit. No Tier C action required.

---

## Skipped repos (active apps/tools — no Tier A/B change)

| Repository | Workflows | Reason |
| --- | --- | --- |
| `chrome-extension-italian-workplace-safety` | 0 | Active browser extension |
| `goccia` | 0 | Potentially active, name unclear |
| `lims` | 0 | Potentially active, name unclear |
| `mcp-italian-workplace-safety` | 0 | Active MCP server |
| `italian-inail-rapporto-annuale-parser` | 0 | Parser/tool, not a static dataset |
| `apd-core` | 0 | Core library/app code |
| `apis-guru-openapi-directory` | 0 | Sync/tooling repo |
| `schemastore` | 0 | Active schema maintenance |

---

## Implementation checklist

Cross-repo changes cannot be applied from this repository session. Each item below is a separate agent session or a manual admin action in GitHub Settings.

**Tier A — Disable in Settings (fastest, no PR needed):**

- [ ] `accordi-stato-regioni-sicurezza-lavoro` → Settings → Actions → General → Disable actions
- [ ] `dlgs-81-08-glossario` → Settings → Actions → General → Disable actions
- [ ] `italian-ateco-database` → Settings → Actions → General → Disable actions
- [ ] `italian-province-regioni-dataset` → Settings → Actions → General → Disable actions
- [ ] `123consulenza-open-data` (this repo) → Settings → Actions → General → Disable actions

**Tier A — Fallback: neutralize file trigger (open a PR in each repo):**

Apply the `workflow_dispatch`-only YAML from the patches above. This keeps the file visible but stops automatic runs.

**Tier B — Harden active lib CI (open a PR in each repo):**

- [ ] `next-seo-italian-helpers` — apply the Tier B patch above
- [ ] `scadenze-formazione-calculator` — apply the Tier B patch above

**Tier C:** No action needed (zero cron jobs found).
