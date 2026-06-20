# API Reference — 123 Consulenza Open Data

API pubblica REST/JSON. **CORS aperto a tutti gli origin**, nessuna autenticazione richiesta per gli endpoint di sola lettura.

**Base URL:** `https://www.123consulenza.com/api/v1`

Tutte le risposte hanno la forma:

```json
{
  "data": ...,
  "meta": { "version": "1.0", "total": 12, "license": "CC BY 4.0", "source": "www.123consulenza.com" }
}
```

In caso di errore:

```json
{ "error": { "code": "...", "message": "..." } }
```

---

## GET `/enti`

Elenco enti di vigilanza per provincia (ASL/ATS, INL, VVF).

**Query parameters**

| Parametro | Tipo | Default | Descrizione |
|---|---|---|---|
| `regione` | string | — | Filtra per regione (es. `Lazio`, `Lombardia`). |
| `limit` | int | 100 | Max 200. |
| `offset` | int | 0 | Paginazione. |

**Esempio**

```bash
curl "https://www.123consulenza.com/api/v1/enti?regione=Lazio&limit=10"
```

---

## GET `/normativa`

Elenco articoli normativi chiave (D.Lgs. 81/08, Reg. CE 852/2004, Reg. CE 178/2002, Reg. UE 2021/382, Reg. UE 1169/2011, ASR 17/04/2025).

**Query parameters**

| Parametro | Tipo | Descrizione |
|---|---|---|
| `settore` | enum | `lavoro` \| `haccp` \| `alimentare`. |

**Esempio**

```bash
curl "https://www.123consulenza.com/api/v1/normativa?settore=haccp"
```

**Schema item**

```json
{
  "id": "dlgs81-art-28",
  "fonte": "D.Lgs. 81/08",
  "numero": "Art. 28",
  "titolo": "Oggetto della valutazione dei rischi",
  "descrizione": "...",
  "riferimento": "D.Lgs. 81/08 — art. 28",
  "settore": ["lavoro"]
}
```

---

## GET `/allergeni`

Lista dei 14 allergeni dell'Allegato II Reg. UE 1169/2011.

```bash
curl "https://www.123consulenza.com/api/v1/allergeni"
```

**Schema item**

```json
{
  "id": 1,
  "nomeItaliano": "Cereali contenenti glutine",
  "nomeTecnico": "Grano, segale, orzo, avena ...",
  "esempiCibi": ["pane", "pasta", "..."],
  "esempiNascosti": ["salse a base di soia", "..."],
  "prevalenza": "celiachia ~1%",
  "gravita": "moderata-severa"
}
```

---

## GET `/sanzioni`

Elenco violazioni e relativi range sanzionatori (sicurezza lavoro, HACCP, etichettatura).

**Query parameters**

| Parametro | Tipo | Descrizione |
|---|---|---|
| `categoria` | enum | `sicurezza-lavoro` \| `haccp` \| `etichettatura-alimenti`. |

**Esempio**

```bash
curl "https://www.123consulenza.com/api/v1/sanzioni?categoria=sicurezza-lavoro"
```

---

## GET `/scadenza-formazione`

Restituisce la data di scadenza dell'aggiornamento per un ruolo specifico. Vedere il sorgente live per i parametri.

---

## GET `/health`

Status endpoint, sempre `200`.

---

## CORS

Tutti gli endpoint pubblici espongono:

```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

## Rate limit

Nessun rate limit hard sul tier pubblico. Si raccomanda di cachare le risposte (i nostri header `Cache-Control` consentono fino a 24h di stale-while-revalidate).

## Licenza

Dati: **CC BY 4.0**. Indicare attribuzione `123 Consulenza — www.123consulenza.com`.
