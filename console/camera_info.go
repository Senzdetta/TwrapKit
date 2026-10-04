// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/camerainfo"
)

type CameraInfo struct{}
func (c CameraInfo) Execute(args []string) {
    camerainfo.CamInfo()
}

// Copyright (c) 2026 Senzdetta