package test

import (
    "strings"
    "testing"
)

func FuzzLogEntryFuzz(f *testing.F) {
    f.Add("cmd=SET:key:value")
    f.Fuzz(func(t *testing.T, input string) {
        if strings.Contains(input, "\n") {
            t.Errorf("Invalid log entry: newline detected in '%s'", input)
        }
    })
}
