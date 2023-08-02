# swiftui-client

> wssocks-ustb client for macOS built by swiftui

## How to build
1. build C header file and archive file from go API.  
    ```bash
    cd extra/go-api
    make
    ```

2. Build App
    ```bash
    xcodebuild -arch=x86_64 -arch=arm64
    ```
