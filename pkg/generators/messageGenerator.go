package generators

import (
	"fmt"
	"github.com/emicklei/proto"
	"io"
	"protots/pkg/mappingTypes"
)

// MessageGenerator generates the messages
func MessageGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithMessage(message))
}

func message(m *proto.Message) {
	printCommentLines(m.Comment.Lines, 0)
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
	_, err := io.WriteString(writer, fmt.Sprintf("\t%s: %s\n", f.Name, field))
	if err != nil {
		panic(err)
	}
}
