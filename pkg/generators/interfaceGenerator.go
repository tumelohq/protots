package generators

import (
	"fmt"
	"github.com/emicklei/proto"
)

// MessageGenerator generates the messages
func InterfaceGenerator(p *proto.Proto) {
	proto.Walk(p, proto.WithService(interfaceFunc))
}

func interfaceFunc(r *proto.Service) {
	writerString(fmt.Sprintf("export interface %s {\n", r.Name))
	for _, element := range r.Elements {
		element.Accept(InterfaceVisitor{})
	}
	writerString(fmt.Sprintf("}\n"))
}

// TODO Need to check if it http api google exists

type InterfaceVisitor struct {
	BaseVisitor
}

func (InterfaceVisitor) VisitRPC(r *proto.RPC) {
	printCommentLines(r.Comment.Lines, 1)
	writerString(fmt.Sprintf("\t%s(arg: %s): Promise<%s> \n", r.Name, r.RequestType, r.ReturnsType))
}
