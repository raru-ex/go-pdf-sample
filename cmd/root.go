package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var templateFilePath string
var defaultTemplateFilePath = "./"

var rootCmd = &cobra.Command{
	Use:   "pdf-diary",
	Short: "output pdf diary from template",
	Long:  `output pdf diary from template`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&templateFilePath, "template", "t", defaultTemplateFilePath, "tamplate pdf file path")
}
