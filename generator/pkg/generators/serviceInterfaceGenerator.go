package generators

import (
	"bytes"
	"github.com/emicklei/proto"
	"log"
	"protots/pkg/googlehttpapi"
	"text/template"
)

const serviceInterfaceTemplate = `
{{range .Comments}}
//{{.}}{{end}}
export interface {{.Name}} {
{{range .Functions}}
	{{range .Comments}}//{{.}}{{end}}
	{{.Name}}(arg: {{.Input}}):Promise<{{.Output}}>
{{end}}
}
`

type serviceInterfaceType struct {
	Name      string
	Comments  []string
	Functions []serviceInterfaceFunction
}

type serviceInterfaceFunction struct {
	Comments []string
	Name     string
	Input    string
	Output   string
}

// ServiceInterface generates interfaces for the proto services and their endpoints. Similarly to the
// ClassGenerator, it only does so if there is an RPC endpoint with an google http api option. For services with a
// combination of both RPCs with and without a google http api option, it will output the class but only with the RPCs
// with a google http option.
func ServiceInterface(p *proto.Proto) {
	for _, e := range p.Elements {
		switch e.(type) {
		case *proto.Service:
			s := e.(*proto.Service)
			if googlehttpapi.DoesServiceContainGoogleHTTPAPIRPCs(s) {
				serviceGeneratorHelperFunction(s)
			}
		}
	}
}

func serviceGeneratorHelperFunction(s *proto.Service) {
	// building service
	var templateOutput serviceInterfaceType
	templateOutput.Name = s.Name
	if s.Comment != nil {
		templateOutput.Comments = s.Comment.Lines
	}

	// getting functions
	for _, e := range s.Elements {
		switch e.(type) {
		case *proto.RPC:
			r := e.(*proto.RPC)
			var comments []string
			if r.Comment != nil {
				comments = r.Comment.Lines
			}
			templateOutput.Functions = append(templateOutput.Functions, serviceInterfaceFunction{
				Comments: comments,
				Name:     r.Name,
				Input:    r.RequestType,
				Output:   r.ReturnsType,
			})
		default:
			log.Fatalf("%+v unexpected type", e)
		}
	}

	// printing output
	t := template.Must(template.New("").Parse(serviceInterfaceTemplate))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, templateOutput)
	if err != nil {
		log.Fatal(err)
	}
	writerString(buf.String())
}
