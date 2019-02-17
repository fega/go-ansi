# ansi

Ansi package, get bash colorized strings

## Usage

```go
package main

import (
	"log"

	ansi "github.com/fega/go-ansi"
)

func main() {
	log.Print(ansi.Cyan("Cyan string"))
	log.Print(ansi.Yellow("Yellow"))
	log.Print(ansi.Black("Black"))
	log.Print(ansi.Blue("Blue"))
	log.Print(ansi.Purple("Purple"))
	log.Print(ansi.Red("Red"))
	log.Print(ansi.Green("Green"))
	log.Print(ansi.White("White"))
}
```