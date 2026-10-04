// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/tphonydevinfo"
)

type TPhonyDevInfo struct{}
func (c TPhonyDevInfo) Execute(args []string) {
    tphonydevinfo.TDevInfo()
}

// Copyright (c) 2026 Senzdetta