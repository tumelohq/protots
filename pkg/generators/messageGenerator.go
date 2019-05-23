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
		if m.Comment != nil {
			if m.Comment.Lines != nil {
				printCommentLines(m.Comment.Lines, 1)
			}
		}
		var templateMessage messageTemplateType
		templateMessage.Name = m.Name
		for _, e := range m.Elements {
			switch e.(type) {
			case *proto.NormalField:
				f := e.(*proto.NormalField)
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

const messageTemplateString = `export interface {{.Name}} {
{{range .Fields}}	{{.Name}}: {{.Type}},
{{end}}}
`

type messageTemplateType struct {
	Name   string
	Fields []messageTemplateFieldName
}

type messageTemplateFieldName struct {
	Name string
	Type string
}

