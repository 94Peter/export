package csv

type DS interface {
	GetHeader() []string
	// 取得下一行資料，無資料回傳nil
	Next() []string
}
