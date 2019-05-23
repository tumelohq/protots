package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
	"path/filepath"
	"protots/pkg/googlehttpapi"
	"strings"
)

// MessageGenerator generates the messages
func ImportGenerator(rootPath string) func(p *proto.Proto) {

	log.Println(rootPath)

	return func(p *proto.Proto) {
		for _, element := range p.Elements {
			if im, ok := element.(*proto.Import); ok {
				if !(im.Filename == googlehttpapi.FilePath) {
					importFunc(rootPath, im)
				}
			}
		}
		writerString(fmt.Sprintf("\n"))
	}
}

func importFunc(roootPath string, i *proto.Import) {
	path := filepath.Join(roootPath, i.Filename)

	if strings.Contains(path, "/google/protobuf/timestamp.proto") {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := proto.NewParser(file)

	parsedProto, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	var messageNames []string

	for _, element := range parsedProto.Elements {
		message, ok := element.(*proto.Message)
		if ok {
			messageNames = append(messageNames, message.Name)
		}
	}

	messagesString := strings.Join(messageNames, ", ")
	importPath := strings.TrimSuffix(path, filepath.Ext(path))
	writerString(fmt.Sprintf(`import {%s} from "src/%s";`, messagesString, importPath))
}
