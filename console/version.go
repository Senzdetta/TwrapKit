// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.ShowVersion()
}

// Copyright (c) 2026 Senzdetta