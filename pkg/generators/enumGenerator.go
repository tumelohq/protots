package generators

import (
	"bytes"
	"github.com/emicklei/proto"
	"log"
	"text/template"
)

// EnumGenerator generates the messages
func EnumGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithEnum(func(e *proto.Enum) {
		var templateType enumTemplateType
		templateType.Name = e.Name
		templateType.Comments = e.Comment.Lines
		for _, e := range e.Elements {
			switch e.(type) {
			case *proto.EnumField:
				field := e.(*proto.EnumField).Name
				templateType.Fields = append(templateType.Fields, field)
			default:
				log.Fatal("could not ")
			}
		}
		t := template.Must(template.New("").Parse(enumTemplateString))
		buf := new(bytes.Buffer)
		err := t.Execute(buf, templateType)
		if err != nil {
			log.Fatal(err)
		}
		writerString(buf.String())
	}))
}

type enumTemplateType struct {
	Name     string
	Comments []string
	Fields   []string
}

const enumTemplateString = `
{{range .Comments}}
//{{.}}{{end}}
export enum {{.Name}} {
{{range .Fields}}	{{.}} = "{{.}}",
{{end}}}
`
