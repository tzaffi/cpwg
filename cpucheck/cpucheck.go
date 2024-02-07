package main

import (
       "fmt"
       "runtime"
)

func main() {
     fmt.Printf("Numer of CPU cores: %d\n", runtime.NumCPU())
     fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
       