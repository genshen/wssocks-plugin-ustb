
<a name="v0.5.1"></a>
## [v0.5.1](https://github.com/genshen/wssocks-plugin-ustb/compare/v0.5.0...v0.5.1)

> 2020-05-01

### Docs

* change default vpn host to n.ustb.edu.cn, rather than vpn4.ustb.edu.cn
* add document of using SwitchyOmega extension as proxy client in Chrome or new Edge
* update document to use socks5 proxy on chromium based browser
* update usage document, add gui document and correct document errors
* **changelog:** add CHANGELOG.md file and git-chglog config
* **readme:** add document of available clients to README.md

### Feat

* **go-module:** update websocket to v1.4.2 and fyne to 1.2.4

### Fix

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

