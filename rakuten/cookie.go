package rakuten

import (
	"fmt"
	"net/http"

	"github.com/sclevine/agouti"
)

const targetPageUrl = "https://www.rakuten-card.co.jp/e-navi/members/statement/index.xhtml"

func GetLoggedInCookies(id string, password string, selectCardNo string) ([]*http.Cookie, error) {
	driver := buildHeadlessDriver()
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		return nil, err
	}

	page, err := driver.NewPage()
	if err != nil {
		return nil, err
	}

	// ご利用明細ページへ遷移
	if err := page.Navigate(targetPageUrl); err != nil {
		return nil, err
	}

	if err := login(*page, id, password); err != nil {
		return nil, err
	}

	if err := selectCard(*page, selectCardNo); err != nil {
		return nil, err
	}

	cookies, err := page.GetCookies()
	if err != nil {
		return nil, err
	}

	return cookies, nil
}

func buildHeadlessDriver() *agouti.WebDriver {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
			"start-maximized",
			"--disable-dev-shm-usage",
			"--window-size=1920,1080",
		})
	driver := agouti.ChromeDriver(options)

	return driver
}

func login(page agouti.Page, id string, password string) error {
	idEle := page.FindByID("u")
	passwordEle := page.FindByID("p")
	idEle.Fill(id)
	passwordEle.Fill(password)

	if err := page.FindByID("indexForm").Submit(); err != nil {
		return err
	}

	return nil
}

func selectCard(page agouti.Page, selectCardNo string) error {
	selector := fmt.Sprintf(`#j_idt609\:card > option:nth-child(%s)`, selectCardNo)
	if err := page.Find(selector).Click(); err != nil {
		return err
	}

	return nil
}
