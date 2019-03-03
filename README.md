## wssocks-plugin-ustb
wssocks-plugin-ustb is a **wssocks** plugin, 
used for accessing internal network of [USTB](http://www.ustb.edu.cn) 
when the internal network is not available directly (such as at home).  

This plugin is based on [wssocks](https://github.com/genshen/wssocks), 
which is a socks5 proxy application over websocket protocol.  
See more about wssocks: https://github.com/genshen/wssocks.

## install
```bash
go get -u github.com/genshen/wssocks-plugin-ustb/wssocks-ustb
wssocks-ustb --help
```

Or download from [release](https://github.com/genshen/wssocks-plugin-ustb/release) page.

Note: this binary contains **wssocks** and wssocks-plugin-ustb plugin.

## Document
- [zh-cn](docs/zh-cn/README.md)
