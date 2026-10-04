// https://github.com/Senzdetta/TwrapKit

package wifiscaninfo

import (
    "strings"
)

func isError(data string) bool {
    return strings.Contains(data, `"error"`) || strings.Contains(data, `"API_ERROR"`)
}

// Copyright (c) 2026 Senzdetta