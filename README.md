# RetroPort

This package provides a simple Go interface to the [USB RetroPort] [hardware].
Only the SNES adapter is supported for now, but I expect it'd be very simple to
add support for the others.


## Documentation

The docs can be found at [godoc.org] [docs]. It's pretty simple.


## Usage

```go
package main

import (
  "fmt"
  "os"
  "time"
  "github.com/adammck/retroport"
)

func main() {

  // open the device
  f, err := os.Open("/dev/hidraw0")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // keep the state synchronized in a goroutine
  c := retroport.MakeSNES(f)
  go c.Run()

  // keep dumping the buttons until L and R are held together
  for {
    if c.Any() {
      if c.L && c.R {
        os.Exit(0)
      }

      fmt.Println(c.Buttons())
    }

    time.Sleep(50 * time.Millisecond)
  }
}

```


## License

[MIT] [license].


## Author

[Adam Mckaig] [adammck] made this.


[hardware]: http://www.retrousb.com/index.php?cPath=21
[license]:  https://github.com/adammck/retroport/blob/master/LICENSE
[docs]:     https://godoc.org/github.com/adammck/retroport
[adammck]:  http://github.com/adammck
