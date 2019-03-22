package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"io"
	"protots/pkg/mappingTypes"
	"strings"
)

// MessageGenerator generates the messages
func MessageGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithMessage(message))
}

func message(m *proto.Message) {
	if m.Comment != nil {
		if m.Comment.Lines != nil {
			printCommentLines(m.Comment.Lines, 1)
		}
	}
	_, err := io.WriteString(writer, fmt.Sprintf("export interface %s {\n", m.Name))
	if err != nil {
		panic(err)
	}
	visitor := messageVisitor{}
	for _, e := range m.Elements {
		e.Accept(visitor)
	}
	_, err = io.WriteString(writer, fmt.Sprintf("}\n\n"))
	if err != nil {
		panic(err)
	}
}

type messageVisitor struct {
	BaseVisitor
}

func (messageVisitor) VisitNormalField(f *proto.NormalField) {
	field := mappingTypes.MapType(f.Type, f.Repeated)
	s := strings.Split(field, ".")
	_, err := io.WriteString(writer, fmt.Sprintf("\t%s: %s\n", f.Name, s[len(s)-1]))
	if err != nil {
		panic(err)
	}
}
