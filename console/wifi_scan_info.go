// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/wifiscaninfo"
)

type WifiScanInfo struct{}
func (c WifiScanInfo) Execute(args []string) {
    wifiscaninfo.WFScan()
}

// Copyright (c) 2026 Senzdetta