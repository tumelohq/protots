package generators

import (
	"github.com/emicklei/proto"
)

type BaseVisitor struct{}

func (BaseVisitor) VisitMessage(m *proto.Message)         {}
func (BaseVisitor) VisitService(s *proto.Service)         {}
func (BaseVisitor) VisitSyntax(s *proto.Syntax)           {}
func (BaseVisitor) VisitPackage(p *proto.Package)         {}
func (BaseVisitor) VisitOption(o *proto.Option)           {}
func (BaseVisitor) VisitImport(i *proto.Import)           {}
func (BaseVisitor) VisitNormalField(i *proto.NormalField) {}
func (BaseVisitor) VisitEnumField(i *proto.EnumField)     {}
func (BaseVisitor) VisitEnum(e *proto.Enum)               {}
func (BaseVisitor) VisitComment(e *proto.Comment)         {}
func (BaseVisitor) VisitOneof(o *proto.Oneof)             {}
func (BaseVisitor) VisitOneofField(o *proto.OneOfField)   {}
func (BaseVisitor) VisitReserved(r *proto.Reserved)       {}
func (BaseVisitor) VisitRPC(r *proto.RPC)                 {}
func (BaseVisitor) VisitMapField(f *proto.MapField)       {}
func (BaseVisitor) VisitGroup(g *proto.Group)             {}
func (BaseVisitor) VisitExtensions(e *proto.Extensions)   {}
