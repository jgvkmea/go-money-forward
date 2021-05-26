package moneyforward

import (
	"github.com/sclevine/agouti"
)

// TopPage マネーフォワードトップページ
type TopPage struct {
	*agouti.Page
}

func newTopPage(page *agouti.Page) *TopPage {
	topPage := &TopPage{
		Page: page,
	}
	return topPage
}

// GoToBankAccountPage 口座情報ページへ遷移
func (p *TopPage) GoToBankAccountPage() (*BankAccountPage, error) {
	err := p.Navigate(bankAccountURL)
	if err != nil {
		return nil, err
	}
	return newBankAccountPage(p.Page), nil
}
