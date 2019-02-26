package generators

import "io"

func printCommentLines(ins []string, numberOfTabs int) {
	out := ""
	for _, in := range ins {
		for i := 0; i < numberOfTabs; i++ {
			out += "\t"
		}
		out += "//" + in + "\n"
	}
	_, err := io.WriteString(writer, out)
	if err != nil {
		panic(err)
	}
}
