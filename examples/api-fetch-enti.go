// Fetch enti di vigilanza (ASL/ATS, INL, VVF) from the 123 Consulenza public API.
// Run: go run examples/api-fetch-enti.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const base = "https://www.123consulenza.com/api/v1"

type Ente struct {
	Nome         string `json:"nome"`
	Denominazione string `json:"denominazione"`
	Provincia    string `json:"provincia"`
	Regione      string `json:"regione"`
}

type Response struct {
	Data []Ente `json:"data"`
}

func fetchEnti(regione string, limit int) ([]Ente, error) {
	u, err := url.Parse(base + "/enti")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("limit", fmt.Sprintf("%d", limit))
	if regione != "" {
		q.Set("regione", regione)
	}
	u.RawQuery = q.Encode()

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http %d", resp.StatusCode)
	}
	var out Response
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

func main() {
	enti, err := fetchEnti("Lazio", 50)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Errore:", err)
		os.Exit(1)
	}
	fmt.Printf("Trovati %d enti nel Lazio:\n", len(enti))
	max := 5
	if len(enti) < max {
		max = len(enti)
	}
	for _, e := range enti[:max] {
		nome := e.Nome
		if nome == "" {
			nome = e.Denominazione
		}
		fmt.Printf("- %s (%s)\n", nome, e.Provincia)
	}
}
