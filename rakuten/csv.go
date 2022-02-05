package rakuten

import (
	"encoding/csv"
	"net/http"
)

const statementPageUrl = "https://www.rakuten-card.co.jp/e-navi/members/statement/index.xhtml?downloadAsCsv=1"

func ReadStatementCsv(cookies []*http.Cookie) ([][]string, error) {
	records, err := readCsvFromUrl(statementPageUrl, cookies)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func readCsvFromUrl(url string, cookies []*http.Cookie) ([][]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	c := new(http.Client)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)
	r.LazyQuotes = true
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
