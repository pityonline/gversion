# gversion

Go project Version from Git tag

## How to use

```Go
import (
	"fmt"
    "github.com/pityonline/gversion"
)

func main() {
	version := Version("./")
	fmt.Printf("Current tag is %q.\n", version)
}
```

### When on a tag

Your app/lib version will be `git describe --always --tags`

### When no tags

Your app/lib version will be `git describe --long --all`
