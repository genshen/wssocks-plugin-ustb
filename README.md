# wssocks-plugin-ustb
![github release](https://badgen.net/github/release/genshen/wssocks-plugin-ustb/stable?label=github-release)
![OSDN](https://img.shields.io/badge/OSDN-night%20release-red?link=https://osdn.net/projects/wssocks-ustb/releases/)
![license](https://badgen.net/github/license/genshen/wssocks-plugin-ustb)

wssocks-plugin-ustb is a **wssocks** plugin, 
used for accessing internal network of [USTB](http://www.ustb.edu.cn) 
when the internal network is not available directly (such as at home).  

This plugin is based on [wssocks](https://github.com/genshen/wssocks), 
which is a socks5 proxy application over websocket protocol.  
See more about wssocks: https://github.com/genshen/wssocks.

## Clients
The available clients for different platforms are list as follows:
- cli: command line client.
- client-ui: From v0.5.0, we also provide GUI [client](client-ui).

Note: **wssocks** and wssocks-plugin-ustb plugin are all included in both cli and client-ui clients.

### install cli client
```bash
go get -u github.com/genshen/wssocks-plugin-ustb/wssocks-ustb
wssocks-ustb --help
```

Or download from github [releases](https://github.com/genshen/wssocks-plugin-ustb/releases) page,
with file name `wssocks-ustb-$OS-$ARCH`.

### install client-ui
You can obtain GUI client from [github releases](https://github.com/genshen/wssocks-plugin-ustb/releases), with file name `client-ui-$OS-$ARCH`.

## Night release clients
If you would like to try new features, you can download wssocks-ustb night release from
[OSDN](https://osdn.net/rel/wssocks-ustb/Night%20release).

## Document
- [zh-cn](docs/zh-cn/README.md)
