package generators

import (
	"bytes"
	"github.com/emicklei/proto"
	"log"
	"protots/pkg/mappingTypes"
	"strings"
	"text/template"
)

// MessageGenerator generates the messages types, it walks through each to make sure that "sub-messages" are also
// generated
func MessageGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithMessage(func(m *proto.Message) {
		var templateMessage messageTemplateType
		templateMessage.Name = m.Name
		if m.Comment != nil {
			templateMessage.Comments = m.Comment.Lines
		}
		for _, e := range m.Elements {
			switch e.(type) {
			case *proto.NormalField:
				f, ok := e.(*proto.NormalField)
				if !ok {
					log.Fatal("not no normal field")
				}
				field := mappingTypes.MapType(f.Type, f.Repeated)
				s := strings.Split(field, ".")
				templateMessage.Fields = append(templateMessage.Fields, messageTemplateFieldName{Name: f.Name, Type: s[len(s)-1]})
			default:
				log.Fatalf("%+v could not be mapped to normal field", e)
			}
		}
		t := template.Must(template.New("").Parse(messageTemplateString))
		buf := new(bytes.Buffer)
		err := t.Execute(buf, templateMessage)
		if err != nil {
			log.Fatal(err)
		}
		writerString(buf.String())
	}))
}

const messageTemplateString = `
{{range .Comments}}
//{{.}}{{end}}
export interface {{.Name}} {
{{range .Fields}}	{{.Name}}: {{.Type}},
{{end}}}
`

type messageTemplateType struct {
	Comments []string
	Name     string
	Fields   []messageTemplateFieldName
}

type messageTemplateFieldName struct {
	Name string
	Type string
}
