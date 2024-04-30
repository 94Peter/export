package excel

import (
	"fmt"
	"io"

	"github.com/tealeg/xlsx"
)

type ExcelReport interface {
	Write(w io.Writer) error
}

func New(ds DS) ExcelReport {
	return &basic{
		DS: ds,
	}
}

type basic struct {
	DS
}

func (c *basic) Write(w io.Writer) error {
	wb := xlsx.NewFile()
	for p, ok := c.DS.NextPage(); ok; p, ok = c.DS.NextPage() {
		sh, err := wb.AddSheet(p.GetName())
		if err != nil {
			return err
		}
		rowNum := 0
		for num, rowVal := p.Next(); num != -1; num, rowVal = p.Next() {
			for j, v := range rowVal {
				sh.Cell(rowNum, j).SetString(fmt.Sprintf("%v", v))
			}
		}
	}
	return wb.Write(w)
}
