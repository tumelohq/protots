package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
)

// ServiceInterfaceGenerator generates interfaces for the proto services and their endpoints
func ServiceInterfaceGenerator(p *proto.Proto) {
	// TODO Remove the walk
	proto.Walk(p, proto.WithService(func(service *proto.Service) {
		writerString(fmt.Sprintf("export interface %s {\n", service.Name))
		for _, e := range service.Elements {
			switch e.(type) {
			case *proto.RPC:
				r := e.(*proto.RPC)
				if r.Comment != nil {
					if r.Comment.Lines != nil {
						printCommentLines(r.Comment.Lines, 1)
					}
				}
				writerString(fmt.Sprintf("\t%s(arg: %s): Promise<%s> \n", r.Name, r.RequestType, r.ReturnsType))
			default:
				log.Fatalf("%+v unexpected type", e)
			}
		}
		writerString(fmt.Sprintf("}\n"))
	}))
}
