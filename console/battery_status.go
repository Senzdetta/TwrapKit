// https://github.com/Senzdetta/TwrapKit

package console

import (
    "github.com/Senzdetta/TwrapKit/module/batterystat"
)

type BatteryStatus struct{}
func (c BatteryStatus) Execute(args []string) {
    batterystat.Battery()
}

// Copyright (c) 2026 Senzdetta