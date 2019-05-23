package generators

import (
	"github.com/emicklei/proto"
)

type Base struct{}

func (Base) VisitMessage(m *proto.Message)         {}
func (Base) VisitService(s *proto.Service)         {}
func (Base) VisitSyntax(s *proto.Syntax)           {}
func (Base) VisitPackage(p *proto.Package)         {}
func (Base) VisitOption(o *proto.Option)           {}
func (Base) VisitImport(i *proto.Import)           {}
func (Base) VisitNormalField(i *proto.NormalField) {}
func (Base) VisitEnumField(i *proto.EnumField)     {}
func (Base) VisitEnum(e *proto.Enum)               {}
func (Base) VisitComment(e *proto.Comment)         {}
func (Base) VisitOneof(o *proto.Oneof)             {}
func (Base) VisitOneofField(o *proto.OneOfField)   {}
func (Base) VisitReserved(r *proto.Reserved)       {}
func (Base) VisitRPC(r *proto.RPC)                 {}
func (Base) VisitMapField(f *proto.MapField)       {}
func (Base) VisitGroup(g *proto.Group)             {}
func (Base) VisitExtensions(e *proto.Extensions)   {}
