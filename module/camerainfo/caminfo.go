// https://github.com/Senzdetta/TwrapKit

package camerainfo

import (
    "fmt"
    "encoding/json"
    "github.com/Senzdetta/TwrapKit/utils/shell"
    "github.com/Senzdetta/TwrapKit/utils/color"
)

func CamInfo() {
    result, err := shell.ExecShell("termux-camera-info")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get camera info!\n",
            color.R, color.N,
        )
        return
    }

    var cams []Info

    err = json.Unmarshal([]byte(result.Stdout), &cams)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(cams)
}

// Copyright (c) 2026 Senzdetta