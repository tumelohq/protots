package generators

import (
	"github.com/emicklei/proto"
	"io"
	"log"
)

var writer io.Writer

func Generate(p *proto.Proto, w io.Writer, generators []Generator) {
	writer = w
	for _, g := range generators {
		g(p)
	}
}

type Generator func(p *proto.Proto)

func writerString(s string) {
	_, err := io.WriteString(writer, s)
	if err != nil {
		log.Fatal(err)
	}
}
