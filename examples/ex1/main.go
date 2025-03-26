package main

import (
	"fmt"
	"strings"

	link "github.com/Beadko/gophercizes_link"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	links, err := link.Parse(strings.NewReader(exampleHTML))
	if err != nil {
		fmt.Printf("Could not parse the links from HTML %s:%v", exampleHTML, err)
		return
	}
	fmt.Printf("%+v\n", links)
}
