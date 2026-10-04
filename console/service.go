// https://github.com/Senzdetta/TwrapKit

package console

import (
    "os"
    "github.com/Senzdetta/TwrapKit/module/service"
    "github.com/Senzdetta/TwrapKit/utils/invinput"
)

type Service struct{}
func (c Service) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    service.Runner(args[2])
}

// Copyright (c) 2026 Senzdetta