# Randomizer

Randomizer generates random stuff. To use it, install the package:

```
$ go get github.com/microwaves/randomizer
```

Then, import on your code:

```
package main

import (
    "fmt"
    
    "github.com/microwaves/randomizer"
)

func main() {
    // the amount of chars for the string output.
    n := 10
    fmt.Println(randomizer.GenerateRandomString(n))

    // random and cute UUIDs based on /dev/urandom.
    fmt.Println(randomizer.GenerateUUID())
}
```

## Author

Stephano Zanzin <sz@shitty.pizza>

## License

Please, refer to the [LICENSE](LICENSE) file.
