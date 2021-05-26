package service

import (
	"fmt"

	"github.com/jgvkmea/go-money-forward/moneyforward"
	"github.com/sclevine/agouti"
	"github.com/sirupsen/logrus"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Chrome/35.0.1916.114 Safari/537.36"
)

// UpdateBankData 口座情報更新
func UpdateBankData(email string, password string) error {
	logger := logrus.New()
	logger.Infoln("start to update account data")

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"no sandbox",
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

	// ログイン
	topPage, err := loginPage.Login(email, password)
	if err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}

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

	logger.Infoln("finish updating account data")
	return nil
}
