package service

import (
	"fmt"
	"os"

	"github.com/jgvkmea/go-money-forward/moneyforward"
	"github.com/sclevine/agouti"
	"github.com/sirupsen/logrus"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Chrome/35.0.1916.114 Safari/537.36"
)

// GetAssetGraphImage 資産推移と割合のグラフをスクショして返す
func GetAssetGraphImage() {
	logger := logrus.New()
	logger.Infoln("Start GetAssetGraphImage()")

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
		return
	}
	portfolioPage.ScreenShot("img/portfolio.jpg")

	// 推移スクショ
	historyPage, err := portfolioPage.GoToHistoryPage()
	if err != nil {
		logger.Errorf("failed to go history page: %v", err)
		return
	}
	historyPage.ScreenShot("img/history.jpg")
}

// UpdateBankData 口座情報更新
func UpdateBankData() error {
	logger := logrus.New()
	logger.Infoln("Start UpdateBankData()")

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			fmt.Sprintf("--user-agent=%s", userAgent),
		}),
	)
	err := driver.Start()
	if err != nil {
		return fmt.Errorf("failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		return fmt.Errorf("failed to create page: %v", err)
	}

	// ログインページに遷移する
	loginPage, err := moneyforward.GoToLoginPage(page)
	if err != nil {
		return fmt.Errorf("failed to go login page: %v", err)
	}
	title, _ := loginPage.Find(".p-user-sign-in > .for-x").Text()
	fmt.Fprintf(os.Stdout, "login: %v\r\n", title)

	// ログイン
	topPage, err := loginPage.Login(email, password)
	if err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}
	title, _ = topPage.Find(".global-menu > ul > li > .active").Text()
	fmt.Fprintf(os.Stdout, "top: %v\r\n", title)

	// 口座ページへ移動
	bankAccountPage, err := topPage.GoToBankAccountPage()
	if err != nil {
		return fmt.Errorf("failed to go to bank account page: %v", err)
	}

	// 更新
	err = bankAccountPage.UpdateBankAccounts()
	if err != nil {
		return fmt.Errorf("failed to update bank accounts: %v", err)
	}
	return nil
}
