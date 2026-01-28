
<a name="v0.8.0-beta"></a>
## [v0.8.0-beta](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.7.0...v0.8.0-beta)

> 2026-01-28

### Docs

* **browser:** use plugin `ZeroOmega` to replace no longer maintained plugin `SwitchyOmega`
* **changelog:** add changelog for release v0.8.0-beta

### Feat

* **client-ui:** add ui to customize chrome installation and pass the selected path to vpn plugin
* **client-ui:** add chromedp-based webview implementation for client-ui
* **version:** bump version to v0.8.0-beta
* **vpn-qr-code:** add support of qr-code vpn login and qr-code login based wssocks connection
* **vpn-qr-code:** feature of saving and storing "auth method" setting value
* **vpn-qr-code:** design ui for both password and qr-code authentication
* **vpn-qr-code:** generate qr code via its content directly, instead of requsting qr image
* **vpn-qr-code:** feature of loading QR code image for QR code login
* **vpn-webview:** add new vpn login method: chromedp-based webview

### Fix

* **compile:** fix building error of swiftui: cannot assign type GoUintptr to type UintPtr
* **swiftui:** fix building error

### Merge

* Merge pull request [#39](https://github.com/genshen/wssocks-plugin-ustb/issues/39) from genshen/style-code-format-and-typo-fixes
* Merge pull request [#32](https://github.com/genshen/wssocks-plugin-ustb/issues/32) from genshen/docs-update-browser-plugin
* Merge pull request [#29](https://github.com/genshen/wssocks-plugin-ustb/issues/29) from genshen/fix-gh-deprecated-artifact-actions
* **gh-action:** Merge pull request [#37](https://github.com/genshen/wssocks-plugin-ustb/issues/37) from genshen/bump-gh-action-os-and-go-version
* **gomodule:** Merge pull request [#40](https://github.com/genshen/wssocks-plugin-ustb/issues/40) from genshen/feature-bump-fyne-2.7.0
* **swiftui:** Merge pull request [#36](https://github.com/genshen/wssocks-plugin-ustb/issues/36) from genshen/fix-swiftui-building
* **vpn-plugin:** Merge pull request [#38](https://github.com/genshen/wssocks-plugin-ustb/issues/38) from genshen/refactor-vpn-login-and-vpn-ui
* **vpn-qr-code:** Merge pull request [#27](https://github.com/genshen/wssocks-plugin-ustb/issues/27) from genshen/feature-vpn-qrcode-login
* **vpn-webview:** Merge pull request [#41](https://github.com/genshen/wssocks-plugin-ustb/issues/41) from genshen/feature-chromedp-based-webview-vpn-login

### Refactor

* **client-ui:** refactor ui: move vpn ui and its preferences to another window
* **vpn-plugin:** move different vpn login methods to different dirs


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.6.0...v0.7.0)

> 2024-02-07

### Docs

* add new screenshot for swiftui client and add screenshots to porject README
* update docs for git lfs and correct/update git and ssh document
* add tips of "connect.exe" path on windows
* add document for macOS client
* update document about versions for downloading
* add downloads number badge and action status badge
* add new version, complete document for wssocks-ustb usage
* **changelog:** update changelog before v0.7.0 release
* **changelog:** update changelog for release v0.7.0
* **client:** correct client filename in document due to filename changing in v0.7.0 release
* **git:** correct ProxyCommand of git under ssh protocol on windows Git Bash

### Feat

* **docs:** change docs theme to `theme-simple` and add edit-on-github button
* **gui:** update document url of fyne-based client to github page
* **gui:** add a tab for separating basic and vpn settings
* **gui:** wait and show fyne-based client error message after the client is started
* **gui:** add input for accept the proxy auth token
* **gui:** feature of copying proxy command of git/ssh/http/https for fyne-based client
* **swiftui:** adjust "quit app" button to right of "Start/Stop" button
* **swiftui:** add feature of copying proxy command for git/ssh/http/https
* **swiftui:** wait and show client error message after the client is started
* **swiftui:** add an easter egg
* **swiftui:** big code refactor and UI redesign(use "popover menu" now)
* **swiftui:** change the swiftui based wssocks client app to menu bar style
* **version:** bump version to v0.7.0

### Fix

* **swiftui:** fix the front color of the status images: only use multicolor for easter egg
* **swiftui:** fix compatibility issue of swiftui client on mac ventura
* **swiftui:** fix bug of "can not open network proxy preference"

### Merge

* Merge pull request [#19](https://github.com/genshen/wssocks-plugin-ustb/issues/19) from genshen/docs-update
* Merge pull request [#7](https://github.com/genshen/wssocks-plugin-ustb/issues/7) from genshen/new-version-document
* **gh-action:** Merge pull request [#9](https://github.com/genshen/wssocks-plugin-ustb/issues/9) from genshen/fix-ci-building-error
* **gomodule:** Merge pull request [#11](https://github.com/genshen/wssocks-plugin-ustb/issues/11) from genshen/dependabot/go_modules/golang.org/x/crypto-0.1.0
* **gui:** Merge pull request [#16](https://github.com/genshen/wssocks-plugin-ustb/issues/16) from genshen/feature-fyne-based-client-mac-arm64-support
* **gui:** Merge pull request [#23](https://github.com/genshen/wssocks-plugin-ustb/issues/23) from genshen/feature-fyne-based-client-copy-proxy-command
* **gui:** Merge pull request [#21](https://github.com/genshen/wssocks-plugin-ustb/issues/21) from genshen/feature-client-ui-wait-error
* **gui:** Merge pull request [#22](https://github.com/genshen/wssocks-plugin-ustb/issues/22) from genshen/feature-new-fyne-based-client-ui-redesign
* **swiftui:** Merge pull request [#17](https://github.com/genshen/wssocks-plugin-ustb/issues/17) from genshen/feature-macos-copy-proxy-command
* **swiftui:** Merge pull request [#14](https://github.com/genshen/wssocks-plugin-ustb/issues/14) from genshen/feature-client-macos-arm-support
* **swiftui:** Merge pull request [#12](https://github.com/genshen/wssocks-plugin-ustb/issues/12) from genshen/feature-swiftui-client-wait
* **swiftui:** Merge pull request [#20](https://github.com/genshen/wssocks-plugin-ustb/issues/20) from genshen/fix-swiftui-status-image-color
* **swiftui:** Merge pull request [#10](https://github.com/genshen/wssocks-plugin-ustb/issues/10) from genshen/fix-swiftui-menu-bar-ventura
* **swiftui:** Merge branch 'fix-swiftui-client-building-error' into 'master'
* **swiftui:** Merge pull request [#8](https://github.com/genshen/wssocks-plugin-ustb/issues/8) from genshen/feature-swiftui-menu-bar


<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.5.1...v0.6.0)

> 2021-01-30

### Docs

* correct ssh ProxyCommand when using wssocks on windows Git Bash
* **changelog:** add changelog for version 0.6.0
* **git-proxy:** correct git http proxy document
* **readme:** update OSDN night release url
* **readme:** update cli/gui installation document and night release installation document
* **swiftui:** add document for building swiftui client

### Feat

* update client and vpn-plugin code to make it compatible with wssocks v0.5.x
* **cli:** better exit code for cli client: Help will not exit with code 1
* **gui:** apply customized fyne theme to client app
* **gui:** bump fyne version to v2.0.0
* **gui:** redesign version display ui and also show plugin version in client ui
* **gui:** add ability of reconfiguring vpn plugin (e.g vpn host) with new input fields under fyne-ui and swiftui
* **gui:** ability to set `SkipTSLVerify` option in client-ui
* **plugin-option:** add option plugin to gui/swiftui and cli client
* **plugin-vpn:** use `SkipTLSVerify` core option to control vpn connection with/without tsl verify
* **swiftui:** ability to set `SkipTSLVerify` option in swiftui
* **swiftui:** add a button to open network proxies preference on mac
* **swiftui:** migrate go api used by swiftui to wssocks v0.5
* **swiftui:** show primary style "Start" button on OSX 11.0
* **swiftui:** now we can pass client handles pointer between swift side and Go side
* **swiftui:** add submit button action to start or stop client task
* **swiftui:** add app icon, version. And disable window resizing
* **swiftui:** add swift-go binding to enable to estabilish connections with server
* **swiftui:** add ui for macOS client which is built by swiftui
* **version:** bump version to v0.6.0

### Fix

* **plugin-vpn:** fix "x509: certificate signed by unknown authority" when connecting
* **swiftui:** set/load preference for "skip tsl verify"
* **swiftui:** also disable username/password/vpn-host input box if vpn is disabled
* **swiftui:** go to status of 'Start' (not status of 'Stop') if error occurs when starting client

### Merge

* Merge pull request [#6](https://github.com/genshen/wssocks-plugin-ustb/issues/6) from genshen/gh-action-build-release
* Merge pull request [#5](https://github.com/genshen/wssocks-plugin-ustb/issues/5) from genshen/migrate-wssocks-v0.5
* Merge pull request [#3](https://github.com/genshen/wssocks-plugin-ustb/issues/3) from genshen/feature-swift-ui-client

### Refactor

* **gui:** move client-ui/background.go to extra/ to provide api for client-ui and swiftui-client


<a name="v0.5.1"></a>
## [v0.5.1](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.5.0...v0.5.1)

> 2020-05-01

### Docs

* change default vpn host to n.ustb.edu.cn, rather than vpn4.ustb.edu.cn
* add document of using SwitchyOmega extension as proxy client in Chrome or new Edge
* update document to use socks5 proxy on chromium based browser
* update usage document, add gui document and correct document errors
* **changelog:** update changelog for release v0.5.1
* **changelog:** add CHANGELOG.md file and git-chglog config
* **readme:** add document of available clients to README.md

### Feat

* **go-module:** update websocket to v1.4.2 and fyne to 1.2.4
* **version:** update version to 0.5.1

### Fix

* **gui:** create Entry and password Entry from struct, thus there is option to show password as plain text
* **gui:** create some of widgets by function (not struct) to fix bug of "initial ui is not freshed"
* **gui:** change default value of vpn host to n.ustb.edu.cn
* **logs:** fix incorrect log format of printing `https(ssl) enabled` information


<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.4.0...v0.5.0)

> 2020-02-24

### Feat

* **go-module:** update wssocks package version to 0.4.1
* **gui:** close tasks when closing gui window
* **gui:** add user data preferences loading and saving
* **gui:** let service(e.g. websocket, tcp connection) of gui client closable and restartable
* **gui:** basic implementation of gui client powered by fyne
* **plugin:** add support for ssl protocol enabled vpn server
* **version:** update version to 0.5.0

### Fix

* **plugin:** fix no cookie found error if http -> https redirection is enabled on vpn site

### Refactor

* **plugin:** make members of struct UstbVpn (for storing vpn config) exported

### Test

* **plugin:** correct parameters in vpn url testing when calling func vpn.vpnUrl()


<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.3.0...v0.4.0)

> 2019-09-01

### Docs

* add document for wssocks-ustb v0.4.0.

### Feat

* **version:** update version to 0.4.0.
* **version-plugin:** add version plugin to handle different versions of client and server.


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.3.0-alpha...v0.3.0)

> 2019-06-16

### Feat

* **logs:** update logs: use logrus log package.
* **version:** update version to 0.3.0.

### Refactor

* **plugin:** use golang.org/x/crypto/ssh/terminal package to read user password, instead of github.com/howeyc/gopass.


<a name="v0.3.0-alpha"></a>
## [v0.3.0-alpha](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.2.0...v0.3.0-alpha)

> 2019-05-13

### Docs

* update document to version 0.3.0.
* add figure for proxifier on windows.

### Feat

* **plugin:** add vpn force logout feature due to the changes of vpn4.ustb.edu.cn.
* **plugin:** add feature of encrypt proxy host due to the changes of vpn4.ustb.edu.cn.

### Fix

* add wssocks-ustb version badge in document.
* **auto-login:** fix logout address due to mistake function name spelling.
* **version:** update version (version sub-command) to v0.3.0-alpha

### Refactor

* **plugin:** change vpn-login-url option to vpn-host.

### Pull Requests

* Merge pull request [#2](https://github.com/genshen/wssocks-plugin-ustb/issues/2) from genshen/develop


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.1.0...v0.2.0)

> 2019-04-11

### Docs

* update document (client side) for this plugin.

### Feat

* **plugin:** add username and password reading for this plugin.
* **plugin:** add auto login feature for this plugin in command line.
* **version:** update version to v0.2.0

### Merge

* **plugin:** Merge pull request [#1](https://github.com/genshen/wssocks-plugin-ustb/issues/1) from genshen/develop


<a name="v0.1.0"></a>
## v0.1.0

> 2019-03-10

