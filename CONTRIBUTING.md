# Contribuire al progetto

Grazie per il tuo interesse! Questo repository raccoglie dataset open e template
fac-simile per la sicurezza sul lavoro e l'HACCP in Italia. Tutti i contributi
sono benvenuti: correzioni, nuovi dati, traduzioni, esempi in nuovi linguaggi.

## Licenze

- **Dati** (`data/`) → CC0-1.0 (pubblico dominio).
- **Codice** (`examples/`, script) → MIT.
- **Template** (`templates/`) → CC0-1.0.

Contribuendo, dichiari di avere il diritto di rilasciare il contenuto sotto la
licenza corrispondente.

## Come contribuire

1. Forka il repo.
2. Crea un branch: `git checkout -b feat/nome-modifica`.
3. Effettua le modifiche rispettando lo schema (`docs/dataset-schema.md`).
4. Valida i JSON: `jq -e . data/<file>.json > /dev/null`.
5. Apri una Pull Request indicando:
   - **Cosa**: dataset/template modificato.
   - **Perché**: motivo (errore, aggiornamento normativo, nuova fonte).
   - **Fonte**: link/citazione di legge o documento ufficiale.

## Convenzioni di stile

### Dataset JSON

- Encoding UTF-8.
- Indentazione 2 spazi.
- Chiavi in `camelCase`.
- Date in ISO 8601.
- Wrapper con `license`, `source`, `updated`, `disclaimer`.

### Template Markdown

- Heading di primo livello con titolo.
- Citazione iniziale (`>`) con riferimento normativo e licenza.
- Campi compilabili con `_______________________________________________________`.
- Footer con sorgente `_Fonte: 123 Consulenza — https://www.123consulenza.com_`.

### Commit message

Stile [Conventional Commits](https://www.conventionalcommits.org/):

- `feat(data): aggiunto dataset X`
- `fix(sanzioni): corretto importo violazione Y`
- `docs(api): aggiornato schema Z`

## Code of conduct

Tutti i partecipanti sono tenuti a un comportamento rispettoso. Vedere il
Contributor Covenant 2.1 (https://www.contributor-covenant.org).

## Domande

Apri una issue con il template `feature` o scrivi a `tutorsicurezza@gmail.com`.
