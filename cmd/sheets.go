package cmd

import (
	"fmt"

	"github.com/fomiller/scribe/sheets"
	"github.com/spf13/cobra"
)

// sheetsCmd represents the sheets command
var sheetsCmd = &cobra.Command{
	Use:   "sheets",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fieldNames := sheets.GetSpreadsheetColumnNames()
		rowData := sheets.GetRowData()
		spreadsheetData := sheets.FmtSpreadsheetData(fieldNames, rowData)
		for _, v := range spreadsheetData {
			for _, vv := range v {
				fmt.Printf("%v: %v\n", vv.FieldName, vv.FieldValue)
			}
			fmt.Println("-----------------------------")
		}

	},
}

func init() {
	rootCmd.AddCommand(sheetsCmd)
}