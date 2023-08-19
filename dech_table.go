package dechtable

import (
	"log"
	"os"

	"github.com/aquasecurity/table"
)

type tableService struct {
	outfile        *os.File
	header         []string
	columnAligment []table.Alignment
	data           [][]string
}

func New(outfile *os.File, header []string, columnAlignment []table.Alignment, data [][]string) *tableService { // If return as pointer then use nil
	return &tableService{
		outfile:        outfile,
		header:         header,
		columnAligment: columnAlignment,
		data:           data,
	}
}

func (rcv *tableService) ExportTableToFile() {
	if rcv.outfile == nil {
		log.Println("Output file is not null")
		return
	}

	//* Console
	// t := table.New(os.Stdout)

	//* File
	tb := table.New(rcv.outfile)

	//? Header Table
	tb.SetDividers(table.UnicodeRoundedDividers)
	tb.SetHeaders(rcv.header...)
	tb.SetAlignment(rcv.columnAligment...)

	//? Data Table
	for _, v := range rcv.data {
		tb.AddRow(v...)
	}

	tb.Render()
}

func (rcv *tableService) DisplayTableToConsole() {
	//* Console
	tb := table.New(os.Stdout)

	//* File
	// tb := table.New(outfile)

	//? Header Table
	tb.SetDividers(table.UnicodeRoundedDividers)
	tb.SetHeaders(rcv.header...)
	tb.SetAlignment(rcv.columnAligment...)

	//? Data Table
	for _, v := range rcv.data {
		tb.AddRow(v...)
	}

	tb.Render()
}
