package main

import (
    "fmt"
    "image"
    _ "image/jpeg"
    "log"
    "os"
)

func main() {
    file, _ := os.Open("hamu.png")
    defer file.Close()

    config, format, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }

    rct := config.Bounds()
    fmt.Println("画像フォーマット：" + format)
    fmt.Println("縦幅=", rct.Dy())
    fmt.Println("横幅=", rct.Dx())
}
