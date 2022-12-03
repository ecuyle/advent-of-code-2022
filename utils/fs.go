package utils

import (
    "os"
)

func Readfile(path string) *os.File {
    file, err := os.Open(path)

    if err != nil {
        panic(err)
    }

    return file
}

