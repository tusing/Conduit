package main

import (
  "fmt"
  "github.com/foo/conduit/provider"
)

func main() {
  fmt.Println(provider.New().Run())
}