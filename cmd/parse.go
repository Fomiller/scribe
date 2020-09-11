package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

var TemplateData = map[string]interface{}{}

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Return a list of fields from a template",
	Long:  `Call Parse command to return a list of fields inside the specified template.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if template flag is not used prompt for the template name

		if TemplateName == "" {
			prompt := &survey.Input{
				Message: "What Template would you like to use",
			}
			survey.AskOne(prompt, &TemplateName)
			fmt.Println("\n")
		}
		// get the docID of the template that needs to be parsed
		templateId, err := drive.GetFileId(TemplateName)
		if err != nil {
			// log.Fatal(err)
			log.Fatalf("File could not be found, %v", err)
		}
		// insert parsedID and return []string of fields in the template
		parsedFields := drive.ParseTemplateFields(templateId)

		qs := []*survey.Question{}
		for _, v := range parsedFields {
			q := &survey.Question{
				Name: v,
				Prompt: &survey.Input{
					Message: fmt.Sprintf("%v:", v),
				},
				Validate: survey.Required,
			}
			qs = append(qs, q)
		}
		fmt.Println("Enter data for the corresponding fields in your template")
		err = survey.Ask(qs, &TemplateData)
		if err != nil {
			log.Fatal(err)
		}

		// convert TemplateData to Fieldmap type map[string]string
		mapStrStr := make(map[string]string)
		for k, v := range TemplateData {
			strKey := fmt.Sprintf("%v", k)
			strValue := fmt.Sprintf("%v", v)

			mapStrStr[strKey] = strValue
		}

		// set Global FieldMap equal to local mapStrStr for use in other functions if needed
		FieldMap = mapStrStr
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// parseCmd.Flags().StringVarP(&ParseTemplate, "parse", "p", "nil", "Enter the name of the template you would like to get the fields of.")
}
