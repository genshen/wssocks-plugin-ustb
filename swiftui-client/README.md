# swiftui-client

> wssocks-ustb client for macOS built by swiftui

## How to build
1. build C header file and archive file from go API.  
    ```bash
    cd extra/go-api
    ./build_archive.sh
    ```

2. Update target’s build settings  
 In Xcode, set `SWIFT_INCLUDE_PATHS` (Header Search Paths) to `$(SRCROOT)`,
 where `$(SRCROOT)` is the same directory as .xcodeproj.
 Then `LIBRARY_SEARCH_PATHS` (Library Search Paths) to `$(SRCROOT)/../extra/go-api`.

3. Build application in Xcode.  
