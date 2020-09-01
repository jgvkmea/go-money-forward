package moneyforward

import (
	"time"

	"github.com/sclevine/agouti"
)

// PortfolioPage 資産割合ページ
type PortfolioPage struct {
	*agouti.Page
}

func newPortfolioPage(page *agouti.Page) *PortfolioPage {
	portfolioPage := &PortfolioPage{
		Page: page,
	}
	return portfolioPage
}

// ScreenShot スクリーンショット
func (p *PortfolioPage) ScreenShot(path string) {
	// 画面サイズ
	p.Size(1920, 1080)
	// 画面の描画待ち
	time.Sleep(3 * time.Second)
	p.Screenshot(path)
}

// GoToHistoryPage 資産推移へ遷移
func (p *PortfolioPage) GoToHistoryPage() (*HistoryPage, error) {
	err := p.Navigate(historyURL)
	if err != nil {
		return nil, err
	}
	return newHistoryPage(p.Page), nil
}
