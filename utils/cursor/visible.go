// https://github.com/Senzdetta/TwrapKit

package cursor

import (
    "fmt"
)

func Visible() {
    fmt.Print("\x1b[?25h")
}

// Copyright (c) 2026 Senzdetta