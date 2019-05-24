package generators

import (
	"io"
	"log"
)

func printCommentLines(ins []string, numberOfTabs int) {
	out := formatCommentLines(ins, numberOfTabs)
	_, err := io.WriteString(writer, out)
	if err != nil {
		log.Printf("error trying to print comments %s", ins)
		log.Fatal(err)
	}
}

func formatCommentLines(ins []string, numberOfTabs int) (out string) {
	for _, in := range ins {
		for i := 0; i < numberOfTabs; i++ {
			out += "\t"
		}
		out += "//" + in + "\n"
	}
	return out
}
