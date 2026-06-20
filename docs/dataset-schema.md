# Dataset Schema — 123 Consulenza Open Data

Documentazione degli schema JSON dei dataset pubblicati in `data/`.

Convenzioni comuni a tutti i dataset:

- Codifica: **UTF-8**.
- Tutte le date sono in formato **ISO 8601** (`YYYY-MM-DD`).
- Ogni file ha un wrapper con metadati:

```json
{
  "license": "CC0-1.0",
  "source": "https://www.123consulenza.com",
  "updated": "2026-06-20",
  "disclaimer": "...",
  "<chiave>": [ ... ]
}
```

---

## `enti-italia.json`

Elenco enti di vigilanza (ASL/ATS/AUSL, INL, VVF) per provincia.

```ts
type Ente = {
  codice: string;
  nome: string;
  tipo: 'ASL' | 'ATS' | 'AUSL' | 'INL' | 'VVF' | string;
  regione: string;
  provincia: string;
  comune?: string;
  indirizzo?: string;
  telefono?: string;
  pec?: string;
  email?: string;
  sito?: string;
};
```

---

## `normativa-italia-ssl.json`

Articoli normativi chiave.

```ts
type Articolo = {
  id: string;            // slug univoco
  fonte: string;         // es. "D.Lgs. 81/08"
  numero: string;        // es. "Art. 28"
  titolo: string;
  descrizione: string;   // 50-150 parole
  riferimento: string;   // fonte ufficiale
  settore: Array<'lavoro' | 'haccp' | 'alimentare'>;
};
```

---

## `sanzioni-ssl.json`

Database violazioni e sanzioni.

```ts
type Violazione = {
  id: string;
  categoria: 'sicurezza-lavoro' | 'haccp' | 'etichettatura-alimenti';
  titolo: string;
  descrizione: string;
  articoloViolato: string;
  fonteSanzionatoria: string;
  tipo: 'penale' | 'amministrativa' | 'penale-amministrativa';
  arrestoMinMesi?: number;
  arrestoMaxMesi?: number;
  importoMin: number;   // EUR
  importoMax: number;   // EUR
};
```

---

## `scadenze-formazione.json`

Durate e periodicità della formazione per ruolo (ASR 17/04/2025).

```ts
type Ruolo = {
  value: string;
  label: string;
  aggiornamentoAnni: number;
  formazioneInizialeOre: number | null;
  aggiornamentoOre?: number;
  nota?: string;
};
```

---

## `allergeni-1169-2011.json`

I 14 allergeni dell'Allegato II Reg. UE 1169/2011.

```ts
type Allergene = {
  id: number;            // 1..14
  nomeItaliano: string;
  nomeTecnico: string;
  esempiCibi: string[];
  esempiNascosti: string[];
  prevalenza: string;
  gravita: 'lieve-moderata' | 'moderata' | 'moderata-severa' | 'severa';
};
```

---

## Validazione

I file JSON sono validi sintatticamente. Per validazione strutturale, è possibile generare schemi JSON Schema a partire da queste definizioni TypeScript (es. con `ts-json-schema-generator`).

## Versioning

I dataset seguono **semver**:

- **patch** (`x.y.Z`) → correzione dati, nuove righe coerenti con lo schema.
- **minor** (`x.Y.0`) → nuovi campi opzionali.
- **major** (`X.0.0`) → modifiche breaking. Annunciate con almeno 30 giorni di preavviso.
