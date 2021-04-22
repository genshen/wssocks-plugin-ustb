# wssocks-ustb 简介
为了访问校内网络，可通过 [https://github.com/genshen/wssocks-plugin-ustb](https://github.com/genshen/wssocks-plugin-ustb) 工具进行连接。
由于本文档主要是面向普通用户，这里仅对其原理做一个简单的介绍，了解这些原理对于使用该工具会有所帮助。

## 技术实现原理
由于 NAT 的限制，在外网（校外）是无法直接连接到校内网络，而学校提供的 [vpn](https://n.ustb.edu.cn) 仅能访问网页 (基于http 协议的请求)，默认还只是限定为几个特定的网站。

[https://github.com/genshen/wssocks](https://github.com/genshen/wssocks) 实现了 SOCKS5 代理协议，并将代理的数据包用 websocket 协议进行包装。 
其包含一个客户端和服务端，客户端和服务端建立一个 websocket 连接，并通过该 websocket 连接传递 SOCKS5 代理协议的数据，提供代理访问功能。  
数据的流动看起来是这样的（一般的，"Your laptop" 和 wssocks client 位于同一台机器上，即在你自己电脑上运行 wssocks-client）:
```
Your laptop ------->wssocks-client ----(gateway)---> wssocks server ----> Target service
                   (Proxy client)                    (Proxy server)       (other network)
```

这样，如果你需要访问的某服务（如ssh 连接到某服务器）在防火墙内部（即不能直接访问该服务）。
但网络管理员允许通过 http/websocket 协议访问到防火墙内部，那么我们就可实现通过 wooscks 代理工具访问防火墙内部的所有基于 TCP 协议的网络(如 ssh 连接)。

例如，北科大校内网络就在一个防火墙内部，我们在校外没法通过 ssh 直接连接到实验室的机器。
但好在我们可以通过学校的 vpn 访问校内的网站(仅仅是基于 http 协议的网站)，同时 vpn 也允许 websocket 数据进入到校内网络。  
基于此，我们就可以在 wssocks client 一侧，将 ssh 连接(是一种TCP 协议)的数据包，利用 SOCCKS 代理协议，并将其塞入到 websocket 数据包中。
然后随着 vpn 进入到校内网络，到达 wssocks server， wssocks server 从 websocket 数据包中解包，并按照 SOCCKS 代理协议将数据的发送到目标服务器。
从目标服务器返回的数据，先是到达 wssocks server，wssocks server 将其按照 SOCCKS 代理协议规范放入 websocket 数据包中，
通过学校 vpn 发送给 wssocks client，随后 wssocks client 解包，将数据通过 SOCCKS 客户端发送给应用程序(如 ssh 客户端)。

总之，数据就是在 "你的本地电脑" <-> wssocks client <-> 学校 VPN <-> wssocks server <-> target service 之间来回传送，
但是，逻辑上，就好像你的本地电脑可以直接访问 target service 一样。
例如你在校外，需要访问校内的服务器，通过 wssocks，就能实现校服务器的访问，就好像你在校内直接访问服务器一样。

当然，学校的 vpn 还有一些特殊的地方，例如需要登录才能使用，不然就不允许访问。
因此，还开发了一个 wssocks 插件，[https://github.com/genshen/wssocks-plugin-ustb/](https://github.com/genshen/wssocks-plugin-ustb/) ，用于登录学校 vpn 服务。
