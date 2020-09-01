package moneyforward

import (
	"time"

	"github.com/sclevine/agouti"
)

// HistoryPage 資産推移ページ
type HistoryPage struct {
	*agouti.Page
}

func newHistoryPage(page *agouti.Page) *HistoryPage {
	historyPage := &HistoryPage{
		Page: page,
	}
	return historyPage
}

// ScreenShot スクリーンショット
func (p *HistoryPage) ScreenShot(path string) {
	// 画面サイズ
	p.Size(1920, 1080)
	// 画面の描画待ち
	time.Sleep(3 * time.Second)
	p.Screenshot(path)
}
