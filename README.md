# gokasi - Kakasi wrapper for Go

This package wraps the [Kakasi C library](http://kakasi.namazu.org/index.html.en) for Go.

### Usage
Import the package:
```
import "github.com/naogogo/gokasi/pkg/gokasi"
```
Then, initialize Kakasi by calling `gokasi.Init()` in your program (equivalent to `kakasi_getopt_argv()`). The default options initialize Kakasi so that it converts everything to ASCII:
```
gokasi.Init()
```
Then, convert something by calling `gokasi.New()` with a string argument:
```
s, err := gokasi.New("こんにちは世界")
```
At the end of your program, free up any resources (equivalent to `kakasi_close_kanwadict()`):
```
gokasi.Destroy()
```
### Complete Example
```
package main

import (
  "fmt"
  "github.com/naogogo/gokasi/pkg/gokasi"
)

func main() {
  gokasi.Init()
  s, _ := gokasi.New("僕は耳と目を閉じ口をつぐんだ人間になろうと考えた")
  fmt.Println(s)
  gokasi.Destroy()
}
```
