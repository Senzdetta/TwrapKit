// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/audioinfo"
)

type AudioInfo struct{}
func (c AudioInfo) Execute(args []string) {
    audioinfo.Audio()
}

// Copyright (c) 2026 Senzdetta