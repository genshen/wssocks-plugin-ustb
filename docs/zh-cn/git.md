# Git 访问校内 Git 服务器

## SSH 协议访问
如果你使用 SSH 协议访问校内的 git 服务器，那么命令会很简单。
以 git clone 为例（仓库地址为`ssh://git@git.hpcer.dev:2222/cli/cli.git`）：
- macOS (ssh 协议)
```bash
GIT_SSH_COMMAND="ssh -o ProxyCommand='nc -x 127.0.0.1:1080 %h %p' " git clone ssh://git@git.hpcer.dev:2222/cli/cli.git
```
- Windows Git Bash (ssh 协议)
```bash
GIT_SSH_COMMAND="ssh -o ProxyCommand='connect -x 127.0.0.1:1080 %h %p' " git clone ssh://git@git.hpcer.dev:2222/cli/cli.git
```

## HTTP/HTTPS 协议访问
wssocks-ustb 客户端需要开启 http/https 代理。
这里全局配置 https://git.hpcer.dev 域名下的所有仓库都通过代理访问，https 代理服务器的地址为 http://127.0.0.1:1080 (与 wssocks-ustb的 socsk5 address 一致)。
```
git config --global http.https://git.hpcer.dev.proxy http://127.0.0.1:1080
git clone https://git.hpcer.dev/genshen/my-project.git
```
注意，如果你的服务器只支持 http 协议，git 的配置的 http 代理服务器地址改为 http://127.0.0.1:1086 (与wssocks-ustb的 htt(s) address 一致):
```
git config --global http.http://git.hpcer.dev.proxy http://127.0.0.1:1086
```

如果需要取消以上的全局配置，可用 `git config --global --unset  http.https://git.hpcer.dev.proxy` 命令(https 协议)
或者 `git config --global --unset  http.http://git.hpcer.dev.proxy` (http 协议)，或者直接修改 `~/.gitconfig` 文件。
