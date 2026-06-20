// Fetch enti di vigilanza (ASL/ATS, INL, VVF) from the 123 Consulenza public API.
// Node.js >= 18 (built-in fetch).
// Run: node examples/api-fetch-enti.js

const BASE = 'https://www.123consulenza.com/api/v1';

async function fetchEnti({ regione, limit = 50 } = {}) {
  const url = new URL(`${BASE}/enti`);
  if (regione) url.searchParams.set('regione', regione);
  url.searchParams.set('limit', String(limit));

  const res = await fetch(url, { headers: { Accept: 'application/json' } });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  const json = await res.json();
  return json.data;
}

(async () => {
  try {
    const lazio = await fetchEnti({ regione: 'Lazio' });
    console.log(`Trovati ${lazio.length} enti nel Lazio:`);
    for (const ente of lazio.slice(0, 5)) {
      console.log(`- ${ente.nome ?? ente.denominazione} (${ente.provincia ?? ''})`);
    }
  } catch (err) {
    console.error('Errore:', err.message);
    process.exit(1);
  }
})();
