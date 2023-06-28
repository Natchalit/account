package interfaces

type AccountingStruct struct {
	Money     float64
	AccountNo string
}
type AccountInterface interface {
	Save(account AccountingStruct) error
}

type GetAccounting struct {
	Money     float64
	AccountNo string
}

func (getAcc *GetAccounting) Save(account AccountingStruct) error {
	return nil
}
