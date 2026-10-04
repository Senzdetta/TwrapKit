// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/contactlist"
)

type ContactList struct{}
func (c ContactList) Execute(args []string) {
    contactlist.DumpList()
}

// Copyright (c) 2026 Senzdetta