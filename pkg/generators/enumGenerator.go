package generators

import (
	"fmt"
	"github.com/emicklei/proto"
)

// EnumGenerator generates the messages
func EnumGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithEnum(enum))
}

func enum(e *proto.Enum) {
	printCommentLines(e.Comment.Lines, 0)
	writerString(fmt.Sprintf("export enum %s {\n", e.Name))
	visitor := enumVisitor{}
	for _, e := range e.Elements {
		e.Accept(visitor)
	}
	writerString(fmt.Sprintf("}\n\n"))
}

type enumVisitor struct {
	BaseVisitor
}

func (enumVisitor) VisitEnumField(i *proto.EnumField) {
	writerString(fmt.Sprintf("\t%s = \"%s\",\n", i.Name, i.Name))
}
