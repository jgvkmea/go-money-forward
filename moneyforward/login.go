package moneyforward

import "github.com/sclevine/agouti"

// LoginPage ログインページ
type LoginPage struct {
	*agouti.Page
	email       *agouti.Selection
	password    *agouti.Selection
	loginButton *agouti.Selection
}

// GoToLoginPage ログインページ取得
func GoToLoginPage(page *agouti.Page) (*LoginPage, error) {
	err := page.Navigate(loginURL)
	if err != nil {
		return nil, err
	}

	return newLoginPage(page), nil
}

func newLoginPage(page *agouti.Page) *LoginPage {
	loginPage := &LoginPage{
		Page:        page,
		email:       page.FindByID("sign_in_session_service_email"),
		password:    page.FindByID("sign_in_session_service_password"),
		loginButton: page.FindByID("login-btn-sumit"),
	}
	return loginPage
}

// Login ログイン処理
func (p *LoginPage) Login(email, password string) (*TopPage, error) {
	// email, password 入力
	err := p.email.Fill(email)
	if err != nil {
		return nil, err
	}
	err = p.password.Fill(password)
	if err != nil {
		return nil, err
	}

	// ログインボタン押下
	err = p.loginButton.Submit()
	if err != nil {
		return nil, err
	}
	return newTopPage(p.Page), nil
}
