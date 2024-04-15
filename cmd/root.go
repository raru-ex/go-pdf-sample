package cmd

import (
	"fmt"
	"os"

	"github.com/raru-ex/pdf-sample/lib"
	"github.com/spf13/cobra"
)

var templateFilePath string
var defaultTemplateFilePath = "./pdf/diary_template.pdf"

var rootCmd = &cobra.Command{
	Use:   "pdf-diary",
	Short: "output pdf diary from template",
	Long:  `output pdf diary from template`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.ExportDiary()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&templateFilePath, "template", "t", defaultTemplateFilePath, "tamplate pdf file path")
}
