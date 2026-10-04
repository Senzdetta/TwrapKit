// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/wificonninfo"
)

type WifiConnInfo struct{}
func (c WifiConnInfo) Execute(args []string) {
    wificonninfo.WFConn()
}

// Copyright (c) 2026 Senzdetta