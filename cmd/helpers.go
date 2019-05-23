package cmd

import (
	"github.com/emicklei/proto"
	"log"
	"os"
	"path/filepath"
	"protots/pkg/generators"
	"regexp"
	"strings"
)

func genFile(rootFilePath string, inputFilePath string, outputFilePath string) {
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
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	generators.Generate(definition, f, []generators.Generator{
		generators.ImportGenerator(rootFilePath),
		generators.ServiceInterfaceGenerator,
		generators.ClassGenerator,
		generators.MessageGenerator,
		generators.EnumGenerator,
	})
}

const typescriptExtension = ".ts"
const protoExtension = ".proto"

func outputFilenameForSingleFile(input string) string {
	_, fileName := filepath.Split(input)
	ext := filepath.Ext(input)
	name := strings.TrimSuffix(fileName, ext)
	return name + typescriptExtension
}

func outputFilenameForFolderGen(input string) string {
	return filepath.Join("gen", filepath.Dir(input), outputFilenameForSingleFile(input))
}

func listAllFiles(folderPath string) (filepaths []string, err error) {
	err = filepath.Walk(folderPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			filepaths = append(filepaths, path)
			return nil
		})
	if err != nil {
		return nil, err
	}
	return filepaths, nil
}

func filterProtoFiles(paths []string) ([]string, error) {
	var outs []string
	for _, path := range paths {
		r, err := regexp.MatchString(protoExtension, path)
		if err != nil {
			return nil, err
		}
		if r {
			outs = append(outs, path)
		}
	}
	return outs, nil
}

const packageJSONFile = `{
	"name": "src"
}`

func generatePackageJSONFile() error {
	path := "gen/pkg.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create("gen/pkg.json")
	if err != nil {
		return err
	}
	_, err = f.WriteString(packageJSONFile)
	if err != nil {
		return err
	}
	return nil
}
