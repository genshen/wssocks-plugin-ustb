//
//  ContentView.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2020/4/30.
//  Copyright © 2020 genshen. All rights reserved.
//

import SwiftUI


extension VerticalAlignment {
    struct CustomAlignment: AlignmentID {
        static func defaultValue(in context: ViewDimensions) -> CGFloat {
            return context[VerticalAlignment.center]
        }
    }
 
    static let custom = VerticalAlignment(CustomAlignment.self)
}

class Configs : ObservableObject {
    @Published var uiSocks5Addr: String = "127.0.0.1:1080"
    @Published var uiRemoteAddr: String = ""
    @Published var uiHttpAddr: String = "127.0.0.1:1086"
    @Published var uiEnableHttpProxy: Bool = false
    @Published var uiSkipTSLerify: Bool = false

    @Published var uiVPNEnable: Bool = true
    @Published var uiVPNForceLogout: Bool = true
    @Published var uiVPNHostEncrypt: Bool = true
    @Published var uiVPNHost: String = "n.ustb.edu.cn"
    @Published var uiVPNUsername: String = ""
    @Published var uiVPNPassword: String = ""
}

struct ContentView: View {
    @ObservedObject var config = Configs()

    @State private var uiEnableSubmitBtn = true
    @State private var uiSubmitBtnLabel = "Start"

    @State private var showingAlert = false
    @State private var alertMessage: String = ""

    private let defaults = UserDefaults.standard
    var client = WssocksClient()

    var body: some View {
        Form {
            Section(header: Text("Basic").bold()) {
                HStack (alignment: .center, spacing: 20) {
                    VStack (alignment: .trailing, spacing: 12) {
                        Text("socks5 address").labelStyle()
                        Text("remote address").labelStyle()
                        Text("http(s) proxy").labelStyle()
                        Text("http(s) address").labelStyle()
                        Text("skip TSL verify").labelStyle()
                    }
                    VStack (alignment: .leading, spacing: 12) {
                        TextField("local socks5 listen address", text: $config.uiSocks5Addr).frame(height: 24)
                        TextField("wssocks server address", text: $config.uiRemoteAddr).frame(height: 24)
                        Toggle(isOn: $config.uiEnableHttpProxy) {
                            Text("Enable http(s) proxy")
                          .multilineTextAlignment(.leading)
                        }.frame(height:24)
                        TextField("http proxy listen address", text: $config.uiHttpAddr)
                            .disableInput(isDisabled: !config.uiEnableHttpProxy).frame(height:24)
                        Toggle(isOn: $config.uiSkipTSLerify) {
                            Text("Skip TSL verify")
                          .multilineTextAlignment(.leading)
                        }.frame(height:24)
                    }
                }
            }
            Section(header: Text("VPN").bold()) {
                HStack (alignment: .center, spacing: 20) {
                    VStack (alignment: .trailing, spacing: 12){
                        Text("Enable VPN").labelStyle()
                        Text("Force logout").labelStyle()
                        Text("Host encrypt").labelStyle()
                        Text("VPN host").labelStyle()
                        Text("Username").labelStyle()
                        Text("Password").labelStyle()
                    }
                    VStack (alignment: .leading, spacing: 12){
                        Toggle("", isOn: $config.uiVPNEnable).frame(height:24)
                        Toggle("", isOn: $config.uiVPNForceLogout).frame(height:24)
                            .disabled(!config.uiVPNEnable)
                        Toggle("", isOn: $config.uiVPNHostEncrypt)
                            .disabled(!config.uiVPNEnable).frame(height:24)
                        TextField("vpn hostname", text: $config.uiVPNHost)
                            .disabled(!config.uiVPNEnable).frame(height:24)
                        TextField("vpn username", text: $config.uiVPNUsername)
                            .disabled(!config.uiVPNEnable).frame(height:24)
                        SecureField("vpn password", text: $config.uiVPNPassword)
                            .disabled(!config.uiVPNEnable).frame(height:24)
                    }
                }
            }
            HStack (alignment: .center, spacing: 20) {
                Button(action: { self.openNetworkProxyPreferences() }) {
                Text("Network Preferences")
                }.buttonStyle(DefaultButtonStyle())
                Spacer()
                if #available(OSX 11.0, *) {
                    Button(action: { self.onSubmit() }) {
                        Text("\(uiSubmitBtnLabel)")
                    }.buttonStyle(BorderedButtonStyle()).disabled(!uiEnableSubmitBtn)
                    .keyboardShortcut(.defaultAction) // see https://stackoverflow.com/a/62727585
                    .alert(isPresented: $showingAlert) {
                        Alert(title: Text("Error"), message: Text("\(alertMessage)"), dismissButton: .default(Text("OK")))
                    }
                } else {
                    // Fallback on earlier versions
                    Button(action: { self.onSubmit() }) {
                        Text("\(uiSubmitBtnLabel)")
                    }.buttonStyle(BorderedButtonStyle()).disabled(!uiEnableSubmitBtn)
                    .alert(isPresented: $showingAlert) {
                        Alert(title: Text("Error"), message: Text("\(alertMessage)"), dismissButton: .default(Text("OK")))
                    }
                }
            }
        }
        .padding(EdgeInsets(top: 8, leading: 8, bottom: 8, trailing: 8))
    }

    func openNetworkProxyPreferences() {
//        let url = URL(string:"x-apple.systempreferences:com.apple.preference.network?Proxies")!
//        NSWorkspace.shared.open(url)
        let script = """
tell application "System Preferences"
    reveal anchor "Proxies" of pane "com.apple.preference.network"
    activate
end tell
"""
        var err: NSDictionary?
        let scriptObject = NSAppleScript(source: script)
        if let output = scriptObject?.executeAndReturnError(&err) {
            print(output.stringValue ?? "")
        } else {
            // something's wrong
        }
    }
    func onSubmit() {
        if !uiEnableSubmitBtn {
            return
        }

        if uiSubmitBtnLabel == "Start" {
            uiEnableSubmitBtn = false
            uiSubmitBtnLabel = "Launching..."
            DispatchQueue.global().async {
                let msg = self.client.startClient(config: self.config) ?? ""
                DispatchQueue.main.sync {
                    if msg != "" {
                        self.alertMessage = msg
                        self.showingAlert = true
                        self.uiSubmitBtnLabel = "Start"
                    } else {
                        self.uiSubmitBtnLabel = "Stop"
                    }
                    self.uiEnableSubmitBtn = true
                }
            }
        } else {
            uiEnableSubmitBtn = false
            uiSubmitBtnLabel = "Stopping..."
            DispatchQueue.global().async {
                let msg = self.client.stopClient() ?? ""
                DispatchQueue.main.sync {
                    if msg != "" {
                        self.alertMessage = msg
                        self.showingAlert = true
                    }
                    self.uiEnableSubmitBtn = true
                    self.uiSubmitBtnLabel = "Start"
                }
            }
        }
    }

    func StoreUserDefaults() {
        defaults.set(true, forKey: "has_preference")

        defaults.set(config.uiSocks5Addr, forKey: "socks5_addr")
        defaults.set(config.uiRemoteAddr, forKey: "remote_addr")
        defaults.set(config.uiHttpAddr, forKey: "http_addr")
        defaults.set(config.uiEnableHttpProxy, forKey: "enable_http_proxy")
        defaults.set(config.uiSkipTSLerify, forKey: "skip_tsl_verify")

        defaults.set(config.uiVPNEnable, forKey: "vpn_enable")
        defaults.set(config.uiVPNForceLogout, forKey: "vpn_force_logout")
        defaults.set(config.uiVPNHostEncrypt, forKey: "vpn_host_encrypt")
        defaults.set(config.uiVPNHost, forKey: "vpn_host")
        defaults.set(config.uiVPNUsername, forKey: "vpn_username")
    }

    func LoadUserDefaults() {
        let _has_preference = defaults.bool(forKey: "has_preference")
        if !_has_preference {
            return
        }

        // if iit falls into default value, we does not load the defaults
        let _socks5_addr = defaults.string(forKey: "socks5_addr") ?? ""
        let _remote_addr = defaults.string(forKey: "remote_addr") ?? ""
        let _http_addr = defaults.string(forKey: "http_addr") ?? ""
        config.uiEnableHttpProxy = defaults.bool(forKey: "enable_http_proxy")
        config.uiSkipTSLerify = defaults.bool(forKey: "skip_tsl_verify")

        config.uiVPNEnable = defaults.bool(forKey: "vpn_enable")
        config.uiVPNForceLogout = defaults.bool(forKey: "vpn_force_logout")
        config.uiVPNHostEncrypt = defaults.bool(forKey: "vpn_host_encrypt")
        let _vpn_host = defaults.string(forKey: "vpn_host") ?? ""
        let _vpn_username = defaults.string(forKey: "vpn_username") ?? ""

        if _socks5_addr.trimmingCharacters(in: .whitespaces) != "" {
            config.uiSocks5Addr = _socks5_addr
        }
        if _remote_addr.trimmingCharacters(in: .whitespaces) != "" {
            config.uiRemoteAddr = _remote_addr
        }
        if _http_addr.trimmingCharacters(in: .whitespaces) != "" {
            config.uiHttpAddr = _http_addr
        }

        if _vpn_host.trimmingCharacters(in: .whitespaces) != "" {
            config.uiVPNHost = _vpn_host
        }
        if _vpn_username.trimmingCharacters(in: .whitespaces) != "" {
            config.uiVPNUsername = _vpn_username
        }
    }
}

extension Text {
    func labelStyle() -> some View {
        AnyView(self.frame(height:24))
    }
}

extension TextField {
    func disableInput(isDisabled: Bool) -> some View {
        if isDisabled {
            return AnyView(self.disabled(true).foregroundColor(Color.gray))
        }
        return AnyView(self)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
