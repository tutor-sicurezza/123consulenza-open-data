# 123 Consulenza Open Data

[![GitHub stars](https://img.shields.io/github/stars/tutor-sicurezza/123consulenza-open-data?style=social)](https://github.com/tutor-sicurezza/123consulenza-open-data/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/tutor-sicurezza/123consulenza-open-data?style=social)](https://github.com/tutor-sicurezza/123consulenza-open-data/network/members)
[![GitHub issues](https://img.shields.io/github/issues/tutor-sicurezza/123consulenza-open-data)](https://github.com/tutor-sicurezza/123consulenza-open-data/issues)
[![License: CC0-1.0](https://img.shields.io/badge/data-CC0%201.0-lightgrey.svg)](./LICENSE-CC0)
[![License: MIT](https://img.shields.io/badge/code-MIT-blue.svg)](./LICENSE-MIT)

> **Mission**: Open data e template open source per **sicurezza sul lavoro e HACCP in Italia**.
> Dataset machine-readable, modelli fac-simile pronti all'uso e snippet di codice riutilizzabili
> per sviluppatori, consulenti, RSPP, OSA, software house, enti di formazione e PA.

Manutenuto da [123 Consulenza](https://www.123consulenza.com) — consulenza sicurezza sul lavoro
e HACCP in Italia. Tutto il materiale e' rilasciato in licenze libere (CC0 per i dati, MIT per il codice).

---

## Quick start

Recupera l'elenco completo dei 107 enti territoriali (ASL/ATS/AUSL, INL, VVF) per provincia:

```bash
curl -sSL https://raw.githubusercontent.com/tutor-sicurezza/123consulenza-open-data/main/data/enti-italia.json \
  | jq '.[] | select(.regione=="Lazio")'
```

Oppure usa direttamente l'URL raw in qualsiasi linguaggio:

```js
const res = await fetch(
  'https://raw.githubusercontent.com/tutor-sicurezza/123consulenza-open-data/main/data/enti-italia.json'
);
const enti = await res.json();
```

---

## Dataset disponibili

Tutti i file sono in `data/` e in formato JSON UTF-8, validati da JSON Schema.

| File | Descrizione | Record |
|------|-------------|--------|
| [`enti-italia.json`](./data/enti-italia.json) | 107 province italiane con ASL/ATS/AUSL/ULSS, Ispettorato del Lavoro, Vigili del Fuoco | ~107 |
| [`normativa-italia-ssl.json`](./data/normativa-italia-ssl.json) | Articoli chiave D.Lgs. 81/08 e regolamenti HACCP (CE 852/2004, UE 2021/382) con sintesi | ~20 |
| [`sanzioni-ssl.json`](./data/sanzioni-ssl.json) | Database sanzioni sicurezza lavoro / HACCP / etichettatura con range ammenda + arresto | ~22 |
| [`ateco-rischi.json`](./data/ateco-rischi.json) | Mapping codici ATECO comuni → rischi tipici + obblighi documentali | 20 |
| [`scadenze-formazione.json`](./data/scadenze-formazione.json) | Durate di validita' per ruolo (lavoratore, preposto, RSPP, RLS, antincendio, primo soccorso) | 13 |
| [`allergeni-1169-2011.json`](./data/allergeni-1169-2011.json) | I 14 allergeni Reg. UE 1169/2011 con esempi e contaminazioni crociate | 14 |
| [`regioni-asl.json`](./data/regioni-asl.json) | Denominazione ASL/ATS/AUSL/ULSS/ASP per regione | 20 |

---

## Template fac-simile

Modelli markdown pronti per essere compilati, in `templates/`:

- [`nomina-rspp.md`](./templates/nomina-rspp.md) — Nomina Responsabile Servizio Prevenzione e Protezione
- [`nomina-rls.md`](./templates/nomina-rls.md) — Verbale di nomina RLS
- [`verbale-riunione-periodica.md`](./templates/verbale-riunione-periodica.md) — Riunione periodica ex art. 35 D.Lgs. 81/08
- [`registro-consegna-dpi.md`](./templates/registro-consegna-dpi.md) — Registro consegna DPI
- [`scheda-formazione.md`](./templates/scheda-formazione.md) — Scheda formazione individuale
- [`registro-temperature-haccp.md`](./templates/registro-temperature-haccp.md) — Registro temperature frigoriferi/celle
- [`registro-infestanti-haccp.md`](./templates/registro-infestanti-haccp.md) — Registro monitoraggio infestanti
- [`incarico-medico-competente.md`](./templates/incarico-medico-competente.md) — Lettera di incarico medico competente
- [`lettera-incarico-cse.md`](./templates/lettera-incarico-cse.md) — Incarico Coordinatore Sicurezza Esecuzione (cantieri)
- [`verbale-formazione-art-37.md`](./templates/verbale-formazione-art-37.md) — Verbale formazione art. 37

> Disclaimer: i modelli sono base di partenza, **non sostitutivi** della valutazione tecnica del caso.

---

## Esempi di consumo

Snippet pronti all'uso in `examples/`:

- Node.js — [`api-fetch-enti.js`](./examples/api-fetch-enti.js)
- Python — [`api-fetch-enti.py`](./examples/api-fetch-enti.py)
- Go — [`api-fetch-enti.go`](./examples/api-fetch-enti.go)
- WordPress shortcode — [`wordpress-shortcode.php`](./examples/wordpress-shortcode.php)
- React hook — [`react-hook-enti.tsx`](./examples/react-hook-enti.tsx)
- Google Apps Script — [`google-sheets-integration.gs`](./examples/google-sheets-integration.gs)
- Airtable sync — [`airtable-sync.js`](./examples/airtable-sync.js)
- Power Automate — [`power-automate-flow.json`](./examples/power-automate-flow.json)

---

## Come contribuire

Le contribuzioni sono benvenute! Leggi [`CONTRIBUTING.md`](./CONTRIBUTING.md) per:

1. Aprire una issue **prima** di inviare una PR (specialmente per nuovi dataset).
2. Rispettare il formato JSON schema (`docs/dataset-schema.md`).
3. Citare la fonte ufficiale per ogni dato aggiunto/modificato.

Issue template disponibili: `bug` (dato sbagliato), `feature` (nuovo dataset), `data-update`.

---

## Licenza

- **Dati** (`data/`, `templates/`): [Creative Commons Zero v1.0 Universal (CC0)](./LICENSE-CC0) — pubblico dominio.
- **Codice** (`examples/`): [MIT](./LICENSE-MIT).

### Citazione richiesta (CC0 — non obbligatoria, ma fortemente apprezzata)

Se usi questi dati in un prodotto, ricerca, articolo o servizio, ti chiediamo cortesemente di citare:

> Fonte: [123 Consulenza Open Data](https://www.123consulenza.com) — `https://github.com/tutor-sicurezza/123consulenza-open-data`

Un link a `https://www.123consulenza.com` ci permette di continuare a investire nella manutenzione dei dataset.

---

## Roadmap

Dataset in lavorazione (PR e suggerimenti benvenuti):

- [ ] `accordi-stato-regioni.json` — repertorio Accordi Stato-Regioni formazione SSL
- [ ] `medici-competenti-elenco.json` — riferimenti elenchi regionali medici competenti
- [ ] `attestati-corsi-riconosciuti.json` — soggetti formatori e attestati con valore legale
- [ ] `near-miss-template.json` — schema per registro mancati infortuni
- [ ] `dpi-categorie.json` — categorie DPI ex Reg. UE 2016/425 con marcatura e norme tecniche
- [ ] `agenti-chimici-pericolosi.json` — sintesi CLP / SDS sezione 2
- [ ] `mmc-niosh.json` — coefficienti NIOSH per movimentazione manuale carichi
- [ ] `rischio-rumore-soglie.json` — soglie azione/limite art. 189 D.Lgs. 81/08
- [ ] `microclima-uni-en-iso-7730.json` — valori PMV/PPD di riferimento

---

## Star History

Se questo progetto ti e' utile, **lascia una star** — ci aiuta a capire l'impatto e giustificare nuovi investimenti.

[![Star History Chart](https://api.star-history.com/svg?repos=tutor-sicurezza/123consulenza-open-data&type=Date)](https://star-history.com/#tutor-sicurezza/123consulenza-open-data&Date)
