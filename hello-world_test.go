package main

import (
    "bytes"
    "os"
    "testing"
)

func TestMainOutput(t *testing.T) {
    // Save the original stdout
    old := os.Stdout
    defer func() { os.Stdout = old }()

    // Capture the stdout
    var buf bytes.Buffer
    os.Stdout = &buf

    // Call the main function
    main()

    // Verify the output
    expected := "hello world\n"
    if buf.String() != expected {
        t.Errorf("Expected %q but got %q", expected, buf.String())
    }
}
