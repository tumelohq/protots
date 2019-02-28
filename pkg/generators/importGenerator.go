package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"path/filepath"
	"protots/pkg/googlehttpapi"
	"strings"
)

// MessageGenerator generates the messages
func ImportGenerator(p *proto.Proto) {
	for _, element := range p.Elements {
		if im, ok := element.(*proto.Import); ok {
			if !(im.Filename == googlehttpapi.FilePath) {
				importFunc(im)
			}
		}
	}
	writerString(fmt.Sprintf("\n"))
}

func importFunc(i *proto.Import) {
	path := i.Filename
	ext := filepath.Ext(path)
	importPath := strings.TrimSuffix(path, ext)
	writerString(fmt.Sprintf("import * from \"src/%s\"\n", importPath))
}
