// https://github.com/Senzdetta/TwrapKit

package console

import (
    "time"
    "fmt"
    "github.com/Senzdetta/TwrapKit/module/uwu"
    "github.com/Senzdetta/TwrapKit/utils/cursor"
)

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Uwu(5 * time.Second)
    cursor.Visible()

    fmt.Println()
}

// Copyright (c) 2026 Senzdetta