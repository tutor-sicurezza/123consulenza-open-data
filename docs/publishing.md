# Publishing — pubblicazione del repository open-data su GitHub

Guida operativa per pubblicare il repository `123consulenza-open-data` su GitHub
sotto l'organizzazione `tutor-sicurezza`. **NON è stato eseguito automaticamente** —
i passi seguenti vanno effettuati manualmente da un manutentore.

## 1. Pre-requisiti

- `gh` CLI installata e autenticata: `gh auth login`.
- Permessi di creazione repo sull'org `tutor-sicurezza`.
- Repo locale già allineato con i contenuti della cartella `open-data/`.

## 2. Estrazione della cartella `open-data/` in repo standalone

Dalla root del monorepo `123consulenza`:

```bash
# Copia in directory fuori dal monorepo
cp -r open-data /tmp/123consulenza-open-data
cd /tmp/123consulenza-open-data

git init
git add .
git commit -m "chore: initial public release"
```

## 3. Creazione del repo remoto

```bash
gh repo create tutor-sicurezza/123consulenza-open-data \
  --public \
  --description "Open data e template open source per sicurezza sul lavoro e HACCP in Italia." \
  --homepage "https://www.123consulenza.com" \
  --license "" # licenze custom: CC0 per dati, MIT per codice
```

Quindi aggiungere il remote e fare push:

```bash
git remote add origin https://github.com/tutor-sicurezza/123consulenza-open-data.git
git branch -M main
git push -u origin main
```

## 4. Topic e About

```bash
gh repo edit tutor-sicurezza/123consulenza-open-data \
  --add-topic "open-data" \
  --add-topic "sicurezza-sul-lavoro" \
  --add-topic "haccp" \
  --add-topic "italia" \
  --add-topic "dlgs-81-08" \
  --add-topic "asr-17-04-2025"
```

## 5. Release iniziale

```bash
gh release create v0.1.0 \
  --title "v0.1.0 — initial public release" \
  --notes "Dataset enti-italia, normativa-italia-ssl, sanzioni-ssl, scadenze-formazione, allergeni-1169-2011. Template fac-simile e esempi API."
```

## 6. Mirror raw URL nel sito principale

Il sito `www.123consulenza.com` referenzia direttamente:

```
https://raw.githubusercontent.com/tutor-sicurezza/123consulenza-open-data/main/data/<file>.json
```

Verificare che i link siano attivi dopo il push.

## 7. CI/CD opzionale

Workflow consigliato in `.github/workflows/validate.yml`:

- validazione JSON con `jq -e .`
- lint markdown su `templates/` e `docs/`
- check link rotti nei docs

## 8. Versioning

Vedere [`dataset-schema.md`](./dataset-schema.md) per la policy semver dei dataset.

## 9. Annuncio

- Aggiungere link al `README.md` del sito principale.
- Pubblicare un post sul blog `/news` di 123 Consulenza che annunci la release.
- Condividere su LinkedIn della società.
