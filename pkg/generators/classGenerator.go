package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"protots/pkg/googlehttpapi"
)

// ClassGenerator generates the class for the particular service only if there is a google http api function in it.
func ClassGenerator(p *proto.Proto) {
	// TODO Remove walk
	proto.Walk(p, proto.WithService(classFunc))
}

func classFunc(s *proto.Service) {
	if googlehttpapi.DoesServiceContainGoogleHTTPAPIRPCs(s) {
		writerString(fmt.Sprintf("export class %s extends ProtoAPIService implements %s{\n", s.Name, s.Name))
		for _, element := range s.Elements {
			element.Accept(classVisitor{})
		}
		writerString(fmt.Sprintf("}\n"))
	}
}

type classVisitor struct {
	Base
}

type classTemplateType struct {
	ClassName          string
	Comments           []string
	ImplementedService string
}

const classTemplateString = `
{{range .Comments}}
//{{.}}{{end}}
export class {{.ClassName}} extends ProtoAPIService implements {{.ImplementedService}}{
{{range .Fields}}	{{.}} = "{{.}}",
{{end}}}
`

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
