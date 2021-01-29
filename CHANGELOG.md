
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
* **swiftui:** add ui for macOS client which is built by swiftui
* **swiftui:** add a button to open network proxies preference on mac
* **swiftui:** ability to set `SkipTSLVerify` option in swiftui
* **swiftui:** migrate go api used by swiftui to wssocks v0.5
* **swiftui:** show primary style "Start" button on OSX 11.0
* **swiftui:** now we can pass client handles pointer between swift side and Go side
* **swiftui:** add submit button action to start or stop client task
* **swiftui:** add app icon, version. And disable window resizing
* **swiftui:** add swift-go binding to enable to estabilish connections with server
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

* update usage document, add gui document and correct document errors
* change default vpn host to n.ustb.edu.cn, rather than vpn4.ustb.edu.cn
* add document of using SwitchyOmega extension as proxy client in Chrome or new Edge
* update document to use socks5 proxy on chromium based browser
* **changelog:** add CHANGELOG.md file and git-chglog config
* **changelog:** update changelog for release v0.5.1
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

