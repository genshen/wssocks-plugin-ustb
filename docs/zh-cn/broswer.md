## 浏览器访问校内网站

## 系统代理：Firefox
如果你的系统上安装有 firefox 浏览器，可以通过以下设置来启用 socks5 代理，当然这仅限在该浏览器中访问网页时会启用代理。 
![](asserts/socks5-firefox.png)

## [推荐] 基于 ZeroOmega 插件：Firefox、 Chrome、新版Edge

Tips: 对于基于 chromium 的浏览器，除了命令行启动外，还可以通过 [ZeroOmega 插件](https://github.com/suziwen/ZeroOmega) 来实现代理的访问。

### 安装 ZeroOmega 插件
用户可以从 [github Release](https://github.com/suziwen/ZeroOmega/release) 下载安装，或者从插件的应用商店安装：  
- [**Microsoft Store**](https://microsoftedge.microsoft.com/addons/detail/proxy-switchyomega-3-zer/dmaldhchmoafliphkijbfhaomcgglmgd) (Edge，这个是非官方版本，建议从 Chrome Web Store 安装)
- [**Chrome Web Store**](https://chromewebstore.google.com/detail/proxy-switchyomega-3-zero/pfnededegaaopdmhkdmcofjmoldfiped) (Chrome)。
- **Firefox浏览器**：可从 Addons 安装插件：https://addons.mozilla.org/en-US/firefox/addon/zeroomega/。  

?> ZeroOmega 源自于 SwitchyOmega，当[2018年以后再也没有更新了](https://github.com/FelisCatus/SwitchyOmega/issues/2513#issuecomment-2218665232)，ZeroOmega是继任者。

### ZeroOmega 配置
安装完成后，设置一条规则，其中代理协议为 SOCKS5，代理服务器和端口和 wssocks-ustb 中配置的 "socks5 address" 一致。
![](./resource/SwitchyOmega.webp)

## 命令行启动：Chrome、新版Edge

对于基于 chromium 的浏览器(如 Chrome, 新版Edge), 可以在命令行中启动浏览器以使用 socks5 代理:
 ```bash
 # chrome on windows
 # tips: windows 用户可以将这一启动命令设置为快捷方式.
 "C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" --show-app-list --proxy-server="socks5://127.0.0.1:1080" --host-resolver-rules="MAP * 0.0.0.0 , EXCLUDE localhost"
 ```
 ```bash
 # chrome on macOS
 /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --show-app-list --proxy-server="socks5://127.0.0.1:1080" --host-resolver-rules="MAP * 0.0.0.0 , EXCLUDE localhost"
 ```
 ```bash
 # new Edge on macOS
 /Applications/Microsoft\ Edge.app/Contents/MacOS/Microsoft\ Edge --show-app-list --proxy-server="socks5://127.0.0.1:1080" --host-resolver-rules="MAP * 0.0.0.0 , EXCLUDE localhost"
 ```
 参考: https://github.com/shadowsocks/shadowsocks/wiki/Forcing-Chrome-to-Use-Socks5-Proxy
