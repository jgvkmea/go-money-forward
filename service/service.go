package service

import (
	"github.com/jgvkmea/go-money-forward/moneyforward"
	"github.com/sclevine/agouti"
	"github.com/sirupsen/logrus"
)

// GetAssetGraphImage 資産推移と割合のグラフをスクショして返す
func GetAssetGraphImage() {
	logger := logrus.New()

	driver := agouti.PhantomJS()
	err := driver.Start()
	if err != nil {
		logger.Errorf("failed to start driver: %v", err)
		return
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		logger.Errorf("failed to create page: %v", err)
		return
	}

	// ログインページに遷移する
	loginPage, err := moneyforward.GoToLoginPage(page)
	if err != nil {
		logger.Errorf("failed to go login page: %v", err)
		return
	}

	// ログイン
	topPage, err := loginPage.Login(email, password)
	if err != nil {
		logger.Errorf("failed to login: %v", err)
		return
	}

	// 割合スクショ
	portfolioPage, err := topPage.GoToPortfolioPage()
	if err != nil {
		logger.Errorf("failed to go portfolio page: %v", err)
	}
	portfolioPage.ScreenShot("img/portfolio.jpg")

	// 推移スクショ
	historyPage, err := portfolioPage.GoToHistoryPage()
	if err != nil {
		logger.Errorf("failed to go history page: %v", err)
	}
	historyPage.ScreenShot("img/history.jpg")
}
