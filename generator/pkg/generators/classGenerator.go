package generators

import (
	"bytes"
	"github.com/emicklei/proto"
	"log"
	"protots/pkg/googlehttpapi"
	"text/template"
)

const classTemplateString = `
{{range .Comments}}
//{{.}}{{end}}
export class {{.ClassName}} extends ProtoAPIService implements {{.ImplementedService}}{
{{range .Functions}}	{{.Name}}(arg: {{.InputArgumentName}}): Promise<{{.OutputArgumentName}}>{
		const u = {{.Source}}
		return this.{{.HTTPMethod}}(u, arg)
}
{{end}}}
`

type classTemplateType struct {
	ClassName          string
	Comments           []string
	ImplementedService string
	Functions          []classTemplateFunctionType
}

type classTemplateFunctionType struct {
	Comments           []string
	Name               string
	InputArgumentName  string
	OutputArgumentName string
	HTTPMethod         string
	Source             string
}

// ClassGenerator generates the class for the service within the proto. It only does so if there is an RPC endpoint with
// an google http api option. For services with a combination of both RPCs with and without a google http api option,
// it will output the class but only with the RPCs with a google http option.
func ClassGenerator(p *proto.Proto) {
	for _, e := range p.Elements {
		switch e.(type) {
		case *proto.Service:
			classFunc(e.(*proto.Service))
		}
	}
}

// classFunc is a slight helper function that outputs the class
func classFunc(s *proto.Service) {
	if googlehttpapi.DoesServiceContainGoogleHTTPAPIRPCs(s) {
		var classTemplate classTemplateType

		// getting class details
		classTemplate.ClassName = s.Name
		classTemplate.ImplementedService = s.Name
		if s.Comment != nil {
			classTemplate.Comments = s.Comment.Lines
		}

		// getting method details
		for _, e := range s.Elements {
			switch e.(type) {
			case *proto.RPC:
				r := e.(*proto.RPC)
				if doesContain, _ := googlehttpapi.DoesRPCContainGoogleHTTPAPI(r); doesContain {
					classTemplate.Functions = append(classTemplate.Functions, vistHTTPAPIRPC(r))
				}
			default:
				log.Fatalf("could not map to rpc element of service %s", s.Name)
			}
		}

		// printing
		t := template.Must(template.New("").Parse(classTemplateString))
		buf := new(bytes.Buffer)
		err := t.Execute(buf, classTemplate)
		if err != nil {
			log.Fatal(err)
		}
		writerString(buf.String())
	}
}

func vistHTTPAPIRPC(r *proto.RPC) (out classTemplateFunctionType) {
	if r.Comment != nil {
		out.Comments = r.Comment.Lines
	}
	out.Name = r.Name
	out.InputArgumentName = r.RequestType
	out.OutputArgumentName = r.ReturnsType

	_, position := googlehttpapi.DoesRPCContainGoogleHTTPAPI(r)
	option := r.Options[position].AggregatedConstants[0]
	// Building url
	u := option.Literal.Source
	u, err := googlehttpapi.Parsing(u)
	if err != nil {
		log.Fatal(err)
	}

	out.Source = u
	out.HTTPMethod = option.Name
	return out
}
