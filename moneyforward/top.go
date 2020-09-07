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

// GoToPortfolioPage 資産割合ページへ遷移
func (p *TopPage) GoToPortfolioPage() (*PortfolioPage, error) {
	err := p.Navigate(portfolioURL)
	if err != nil {
		return nil, err
	}
	return newPortfolioPage(p.Page), nil
}

// GoToBankAccountPage 口座情報ページへ遷移
func (p *TopPage) GoToBankAccountPage() (*BankAccountPage, error) {
	err := p.Navigate(bankAccountURL)
	if err != nil {
		return nil, err
	}
	return newBankAccountPage(p.Page), nil
}
