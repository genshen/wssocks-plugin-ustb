## wssocks-plugin-ustb
wssocks-plugin-ustb is a **wssocks** plugin, 
used for accessing internal network of [USTB](http://www.ustb.edu.cn) 
when the internal network is not available directly (such as at home).  

This plugin is based on [wssocks](https://github.com/genshen/wssocks), 
which is a socks5 proxy application over websocket protocol.  
See more about wssocks: https://github.com/genshen/wssocks.

## clients
The available clients for different platforms are list as follows:
- cli: command line client.
- client-ui: From v0.5.0, we also provide a [client](client-ui) with GUI.
  You can obtain it from [releases](https://github.com/genshen/wssocks-plugin-ustb/releases) page,
  whose file name is `client-ui-$OS-$ARCH`.

## install cli client
```bash
go get -u github.com/genshen/wssocks-plugin-ustb/wssocks-ustb
wssocks-ustb --help
```

Or download from [releases](https://github.com/genshen/wssocks-plugin-ustb/releases) page,
whose file name is `wssocks-ustb-$OS-$ARCH`.

Note: this binary contains **wssocks** and wssocks-plugin-ustb plugin.

## Document
- [zh-cn](docs/zh-cn/README.md)
