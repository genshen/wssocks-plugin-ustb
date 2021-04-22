# SSH 连接校内服务器

### macOS
```bash
ssh -o ProxyCommand='nc -x 127.0.0.1:1080 %h %p' user@ssh.hpcer.dev
```
`user@ssh.hpcer.dev` 替换为你需要连接的目标服务器用户和地址，`127.0.0.1:1080` 为前面 wssocks-ustb 客户端配置的 SOCKS5 监听地址和端口。

### windows Git Bash
如果，你的 windows 安装了 git bash，可以通过以下命令连接到目标服务器。
```bash
ssh -o ProxyCommand='connect -S 127.0.0.1:1080 %h %p' user@ssh.hpcer.dev
```
`user@ssh.hpcer.dev` 替换为你需要连接的目标服务器用户和地址，`127.0.0.1:1080` 为前面 wssocks-ustb 客户端配置的 SOCKS5 监听地址和端口。

### 其他 shell 客户端
其他平台上的一些 shell 客户端(如xshell)，也可对ssh连接进行 socks5 代理配置。
具体参见相关配置文档，主要配置内容是设置连接的代理类型为 SOSCK5，且设置 socks5 服务器地址为 wssocks-ustb 客户端对应的 "socks5 address"（如 `127.0.0.1:1080`）。
