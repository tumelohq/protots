package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var output string

// genfilesCmd represents the gen command
var genfilesCmd = &cobra.Command{
	Use:   "genfiles [file path 1] [file path 2] ...",
	Short: "genfiles generates the ts file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("require filepath, not %d arguments", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			genFile(arg, outputFilenameForSingleFile(arg))
		}
	},
}

func init() {
	rootCmd.AddCommand(genfilesCmd)
	genfilesCmd.Flags().StringVarP(&output, "output", "o", "", "file to write to")
}

