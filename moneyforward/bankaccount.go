package moneyforward

import (
	"fmt"

	"github.com/sclevine/agouti"
)

// BankAccountPage 口座情報ページ
type BankAccountPage struct {
	*agouti.Page
}

func newBankAccountPage(page *agouti.Page) *BankAccountPage {
	bankAccountPage := &BankAccountPage{
		Page: page,
	}
	return bankAccountPage
}

// UpdateBankAccounts 口座ページにある口座を更新する
func (p *BankAccountPage) UpdateBankAccounts() error {
	// 口座情報リスト取得
	bankList := p.All("#account-table > tbody > tr")
	bankListCount, err := bankList.Count()
	if err != nil {
		return fmt.Errorf("failed to get bankListCount: %v", err)
	}

	for i := 1; i < bankListCount; i++ {
		id, err := bankList.At(i).Attribute("id")
		if err != nil {
			return fmt.Errorf("failed to get id: %v", err)
		}

		err = p.Find(fmt.Sprintf("#js-recorrect-form-%s", id)).FindByClass("btn").Submit()
		if err != nil {
			return fmt.Errorf("failed to submit recorrect: %v", err)
		}
	}
	return nil
}
