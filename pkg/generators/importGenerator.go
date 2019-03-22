package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"path/filepath"
	"protots/pkg/googlehttpapi"
)

// MessageGenerator generates the messages
func ImportGenerator(rootPath string) func(p *proto.Proto) {
	return func(p *proto.Proto) {
		for _, element := range p.Elements {
			if im, ok := element.(*proto.Import); ok {
				if !(im.Filename == googlehttpapi.FilePath) {
					importFunc(im)
				}
			}
		}
		writerString(fmt.Sprintf("\n"))
	}
}

func importFunc(roootPath string, i *proto.Import) {
	path := filepath.Join(roootPath, i.Filename)



	proto.NewParser()


}
