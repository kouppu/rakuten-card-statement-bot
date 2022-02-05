package rakuten

import (
	"net/http"

	"github.com/sclevine/agouti"
)

const targetPageUrl = "https://www.rakuten-card.co.jp/e-navi/members/statement/index.xhtml"

func GetLoggedInCookies(id string, password string, cardName string) ([]*http.Cookie, error) {
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

	if err := selectCard(*page, cardName); err != nil {
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

func selectCard(page agouti.Page, cardName string) error {
	// TODO: 引数からカードを選択できるようにする
	if err := page.FindByID("j_idt609:card").Select(cardName); err != nil {
		return err
	}

	return nil
}
