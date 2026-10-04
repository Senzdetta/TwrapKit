// https://github.com/Senzdetta/TwrapKit

package console

type Command interface {
    Execute(args []string)
}

// Copyright (c) 2026 Senzdetta