package cmd

import (
	"fmt"

	"github.com/fomiller/scribe/docs"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// RunCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		fmt.Println("<--------------------------------->")

		templateCopyId := drive.CreateTemplateCopy()
		fmt.Println("This is the template Copy id: ", templateCopyId)
		templateId := docs.UpdateTemplateFile(templateCopyId)

		fmt.Println("Successfully updated document: ", templateId)

	},
}

func init() {
	rootCmd.AddCommand(RunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}