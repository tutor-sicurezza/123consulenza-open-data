"""Fetch enti di vigilanza (ASL/ATS, INL, VVF) from the 123 Consulenza public API.

Requires: pip install requests
Run: python examples/api-fetch-enti.py
"""
from __future__ import annotations

import sys
import requests

BASE = "https://www.123consulenza.com/api/v1"


def fetch_enti(regione: str | None = None, limit: int = 50) -> list[dict]:
    params: dict[str, str] = {"limit": str(limit)}
    if regione:
        params["regione"] = regione
    r = requests.get(f"{BASE}/enti", params=params, timeout=15)
    r.raise_for_status()
    return r.json().get("data", [])


def main() -> int:
    try:
        lazio = fetch_enti(regione="Lazio")
    except requests.RequestException as exc:
        print(f"Errore: {exc}", file=sys.stderr)
        return 1
    print(f"Trovati {len(lazio)} enti nel Lazio:")
    for ente in lazio[:5]:
        nome = ente.get("nome") or ente.get("denominazione") or "-"
        prov = ente.get("provincia", "")
        print(f"- {nome} ({prov})")
    return 0


if __name__ == "__main__":
    sys.exit(main())
