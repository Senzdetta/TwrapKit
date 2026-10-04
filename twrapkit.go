// https://github.com/Senzdetta/TwrapKit

package main

import (
    "os"
    "strings"
    "github.com/Senzdetta/TwrapKit/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.TwrapKitConsole(input)
}

// Copyright (c) 2026 Senzdetta