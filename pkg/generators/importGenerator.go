package generators

import (
	"fmt"
	"github.com/emicklei/proto"
)

// MessageGenerator generates the messages
func ImportGenerator(p *proto.Proto) {
	for _, element := range p.Elements {
		if im, ok := element.(*proto.Import); ok {
			importFunc(im)
		}
	}
	writerString(fmt.Sprintf("\n"))
}

func importFunc(i *proto.Import) {
	writerString(fmt.Sprintf("import \"%s\"\n", i.Filename))
}
