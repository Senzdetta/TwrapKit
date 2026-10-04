// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/calllog"
)

type CallLog struct{}
func (c CallLog) Execute(args []string) {
    calllog.Dump()
}

// Copyright (c) 2026 Senzdetta