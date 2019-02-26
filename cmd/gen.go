package cmd

import (
	"fmt"
	"github.com/emicklei/proto"
	"path/filepath"
	"protots/pkg/generators"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var output string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen [file path]",
	Short: "gen generates the ts file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("require filepath, not %d arguments", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			inputFilePath := arg
			reader, err := os.Open(inputFilePath)
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			parser := proto.NewParser(reader)
			definition, err := parser.Parse()
			if err != nil {
				log.Fatal(err)
			}
			f, err := os.Create(outputFilename(inputFilePath))
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			generators.Generate(definition, f, []generators.Generator{
				generators.ImportGenerator,
				generators.InterfaceGenerator,
				generators.ClassGenerator,
				generators.MessageGenerator,
				generators.EnumGenerator,
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&output, "output", "o", "", "file to write to")
}

const typescriptExtension = ".ts"

func outputFilename(input string) string {
	if output == "" {
		_, fileName := filepath.Split(input)
		ext := filepath.Ext(input)
		name := strings.TrimSuffix(fileName, ext)
		return name + typescriptExtension
	}
	return output
}
