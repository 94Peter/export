package csv

import (
	"encoding/csv"
	"io"
)

type CsvReport interface {
	Write(w io.Writer) error
}

func NewCsv(ds DS) CsvReport {
	return &basicCsv{
		DS: ds,
	}
}

type basicCsv struct {
	DS
}

func (c *basicCsv) Write(w io.Writer) error {
	// UTF-8 Bom 避免Excel打開來是亂碼
	w.Write([]byte("\xEF\xBB\xBF"))

	writer := csv.NewWriter(w)
	defer writer.Flush()
	if h := c.GetHeader(); h != nil {
		writer.Write(h)
	}
	for d := c.Next(); d != nil; d = c.Next() {
		writer.Write(d)
	}
	return nil
}
