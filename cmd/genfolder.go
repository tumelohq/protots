package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var genfolderCmd = &cobra.Command{
	Use:   "genfolder [folder path]",
	Short: "genfiles generates the ts files for all the protos in the file with the same folder structure",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("require single argument, one folder path")
		}
		fi, err := os.Stat(args[0])
		if err != nil {
			return err
		}
		if !fi.Mode().IsDir() {
			return errors.New("require folder path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := generatePackageJSONFile()
		if err != nil {
			log.Fatal(err)
		}
		allFiles, err := listAllFiles(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(allFiles)
		tsFiles, err := filterProtoFiles(allFiles)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tsFiles)
		for _, path := range tsFiles {
			genFile(path, outputFilenameForFolderGen(path))
		}
	},
}


func init() {
	rootCmd.AddCommand(genfolderCmd)
}
