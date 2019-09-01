> 从校外网络访问USTB校内网络

![wssocks-ustb v0.4.0](https://img.shields.io/badge/wssocks--ustb-v0.4.0-orange.svg)

## 使用示例:服务端
 在内网的服务端的主机(如地址为`proxy.gensh.me`的主机)上执行:
 ```bash
 # For instance, listener on port 80.
 # And make sure your port has been added to the white list of your firewall.
 wssocks server --addr :80
 ```
 例如上面的命令，在服务端监听80端口(或者是别的端口,但client端`remote`选项需要和其相一致)。另外，需要确保你监听的端口已经被添加到主机防火墙的白名单列表中。

 see more: https://github.com/genshen/wssocks#server-side.

## 使用示例:客户端
1. 打开命令行,执行如下命令,运行客户端:
   > 下面命令中, `wssocks-ustb`可执行程序均指代包含`wssocks-plugin-ustb`插件功能的`wssocks`程序.
   ```bash
   wssocks-ustb client --remote=ws://proxy.gensh.me --http -http-addr=:1086 --vpn-enable   --vpn-host=vpn4.ustb.edu.cn --vpn-force-logout --vpn-host-encrypt
   ```
   以上命令通过启用`--vpn-enable`选项启用通过vpn连接校内到网络。
   随后, 要求输入vpn的用户名和密码登录`n.ustb.edu.cn`以获取其cookie (用户名和密码也可以在命令中通过`--vpn-usernam`和`--vpn-password`选项指定)。  
   此外，客户端默认本地监听地址为`:1080`(即0.0.0.0:1080), 服务器端地址为`ws://proxy.gensh.me`。

   也可以通过`wssocks-ustb client --help`查看更多参数的使用。 其中, 几个主要命令参数如下:
   - `--addr` 指定客户端默认本地监听地址,默认为`:1080`(即0.0.0.0:1080);
   - `--remote` 指定服务器端地址;
   - `--http` 启用 http 和 https 代理 
   - `--http-addr` 若启用 http  和 https 代理，该选项指定http代理的本地监听地址及端口(仅http代理地址，https代理的地址和socks5一致)，默认 `:1086`
   - `--vpn-enable` 是否开启vpn模式;如不开启vpn模式, 将跳过所有以vpn开头的参数;
   - `--vpn-host` vpn服务器主机地址;
   - `--vpn-username` 登录vpn的用户名;如不在命令参数中指定,将会以交互的方式获取;
   - `--vpn-password` 登录vpn的密码; 如不在命令参数中指定,将会以交互的方式获取(为安全起见,不推荐在命令参数中指定);
   - `--vpn-force-logout` 如果账号已经在其他设备上登录,强制退出其他设备上的账号;
   - `--vpn-host-encrypt` 使用 aes 算法加密代理服务器主机名,默认启用;

2. socks5 代理客户端设置  
  使用socks代理客户端软件(如mac系统的全局代理功能), 设置代理地址。
  ![](asserts/mac-proxy.png)
  在mac中，勾选**socks代理**选项框, 并填入代理服务器的地址及端口(即wssocks客户端本地监听地址及端口)，保存生效。  
  如果你使用的是windows, 可以使用[Proxifier](https://www.proxifier.com/)软件来设置全局代理。  
  (下图proxifier界面来自于  www.proxifier.com , 以展示添加代理方法。使用wssocks时，图中各个字段(如地址和端口)和选项会有区别。)
  ![proxifier](https://www.proxifier.com/screenshots/proxy.png)   
  另外，一些软件的设置中也会有socks5代理选项，也可以针对特定软件进行设置，当然这种设置只针对该软件有效。  

4. http 与 https 代理客户端设置  
  在 mac 的系统偏好设置或 windows 的网络设置中，可设置全局的 http 和 https 代理; 部分软件可设置针对该软件有效的http和https代理(如firefox浏览器)。  
  其中，http代理服务器地址由命令行中的`--http-addr`指定，默认为`:1086`(或`127.0.1.0:1086`)，https 代理服务器地址和socks5代理服务器地址一致。

5. 访问网页  
  如果你已经设置好了 http、 https 代理，或者设置了 socks5 代理，可直接在浏览器地址栏输入对应的地址即可访问，不用任何特殊设置。

6. ssh连接(仅macOS)
   ```bash
   ssh -o ProxyCommand='nc -x 127.0.0.1:1080 %h %p' ssh.hpc.gensh.me
   ```
   ![](asserts/ssh-example.png)  
   windows和linux中可直接使用类似的`ssh ssh.hpcer.dev`命令。

7. git 命令代理
   - macOS
      ```bash
      GIT_SSH_COMMAND="ssh -o ProxyCommand='nc -x 127.0.0.1:1080 %h %p' " git clone repo_address
      ```
   - 通用设置(macOS、windows、linux)  
      设置 http(s) 代理:
      ```bash
      git config --global http.proxy http://127.0.0.1:1086
      git config --global https.proxy http://127.0.0.1:1080
      ```
      设置 socks5 代理:
      ```bash
      git config --global http.proxy socks5://127.0.0.1:1080
      git config --global https.proxy socks5://127.0.0.1:1080
      ```
      Git 取消代理设置:
      ```bash
      git config --global --unset http.proxy
      git config --global --unset https.proxy
      ```
      具体可参考 https://git-scm.com/docs/git-config.

8. 其他终端操作(macOS、linux)  
   终端设置 http(s) 代理(只对当前终端有效):  
   ```bash
   export http_proxy=http://127.0.0.1:1086
   export https_proxy=http://127.0.0.1:1080
   ```
   终端设置 socks5 代理(只对当前终端有效):  
   ```bash
   export http_proxy=socks5://127.0.0.1:1080
   export https_proxy=socks5://127.0.0.1:1080
   ```
   设置终端中的 wget、curl 等都走 socks5 代理(只对当前终端有效):
   ```bash
   export ALL_PROXY=socks5://127.0.0.1:1080
   ```
