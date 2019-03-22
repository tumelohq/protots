package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"protots/pkg/googlehttpapi"
	"log"
)

// MessageGenerator generates the messages
func ClassGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithService(classFunc))
}

func classFunc(r *proto.Service) {
	writerString(fmt.Sprintf("export class %s extends ProtoAPIService implements %s{\n", r.Name, r.Name))
	for _, element := range r.Elements {
		element.Accept(classVisitor{})
	}
	writerString(fmt.Sprintf("}\n"))
}

type classVisitor struct {
	BaseVisitor
}

// TODO Check if there are any google.http.api functions

func (classVisitor) VisitRPC(r *proto.RPC) {
	httpAPI, position := googlehttpapi.DoesRPCContainGoogleHTTPAPI(r)
	if httpAPI {
		if r.Comment != nil {
			if r.Comment.Lines != nil {
				printCommentLines(r.Comment.Lines, 1)
			}
		}
		writerString(fmt.Sprintf("\t%s(arg: %s): Promise<%s>{ \n", r.Name, r.RequestType, r.ReturnsType))
		option := r.Options[position].AggregatedConstants[0]
		// Building url
		u := option.Literal.Source
		u, err := googlehttpapi.Parsing(u)
		if err != nil {
			log.Fatal(err)
		}
		writerString(fmt.Sprintf("\t\tconst u = %s;\n", u))
		writerString(fmt.Sprintf("\t\treturn this.%s(u, arg)\n", option.Name))
		writerString(fmt.Sprintf("\t}\n\n"))
	}
}

