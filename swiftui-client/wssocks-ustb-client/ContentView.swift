//
//  ContentView.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2020/4/30.
//  Updated to menu bar app by genshen on 2021/8/18.
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

struct ContentView: View {
    @State private var selection = 0
    
    // The only config instance is here.
    // When the window and view is appreared, the config is loaded from preference file.
    // When the app is terminated, the config will be store to preference file.
    // While running, the config is shared by this instance.
    @ObservedObject var config = Configs()

    @State private var uiEnableSubmitBtn = true
    @State private var uiSubmitBtnLabel = "Start"

    @State private var showingAlert = false
    @State private var alertMessage: String = ""

    var body: some View {
        TabView {
            Form {
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
            }.padding(EdgeInsets(top: 8, leading: 8, bottom: 8, trailing: 8))
            .tabItem {
                if #available(macOS 11.0, *) {
                    Label("General", systemImage: "list.dash")
                } else {
                    Text("General")
                }
            }
            Form {
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
            }.padding(EdgeInsets(top: 8, leading: 8, bottom: 8, trailing: 8))
            .tabItem {
                if #available(macOS 11.0, *) {
                    Label("USTB VPN", systemImage: "list.dash")
                } else {
                    Text("USTB VPN")
                }
            }
        }.padding(EdgeInsets(top: 10, leading: 0, bottom: 0, trailing: 0))
        .onAppear{
            let defaults = UserDefaults.standard
            config.LoadUserDefaults(defaults: defaults)
        } // load preference
    }

    func onSubmit() {
        if !uiEnableSubmitBtn {
            return
        }
    }
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

    func StoreUserDefaults(defaults: UserDefaults) {
        defaults.set(true, forKey: "has_preference")

        defaults.set(self.uiSocks5Addr, forKey: "socks5_addr")
        defaults.set(self.uiRemoteAddr, forKey: "remote_addr")
        defaults.set(self.uiHttpAddr, forKey: "http_addr")
        defaults.set(self.uiEnableHttpProxy, forKey: "enable_http_proxy")
        defaults.set(self.uiSkipTSLerify, forKey: "skip_tsl_verify")

        defaults.set(self.uiVPNEnable, forKey: "vpn_enable")
        defaults.set(self.uiVPNForceLogout, forKey: "vpn_force_logout")
        defaults.set(self.uiVPNHostEncrypt, forKey: "vpn_host_encrypt")
        defaults.set(self.uiVPNHost, forKey: "vpn_host")
        defaults.set(self.uiVPNUsername, forKey: "vpn_username")
        defaults.set(self.uiVPNPassword, forKey: "vpn_passwd")
    }

    func LoadUserDefaults(defaults: UserDefaults) {
        let _has_preference = defaults.bool(forKey: "has_preference")
        if !_has_preference {
            return
        }

        // if iit falls into default value, we does not load the defaults
        let _socks5_addr = defaults.string(forKey: "socks5_addr") ?? ""
        let _remote_addr = defaults.string(forKey: "remote_addr") ?? ""
        let _http_addr = defaults.string(forKey: "http_addr") ?? ""
        self.uiEnableHttpProxy = defaults.bool(forKey: "enable_http_proxy")
        self.uiSkipTSLerify = defaults.bool(forKey: "skip_tsl_verify")

        self.uiVPNEnable = defaults.bool(forKey: "vpn_enable")
        self.uiVPNForceLogout = defaults.bool(forKey: "vpn_force_logout")
        self.uiVPNHostEncrypt = defaults.bool(forKey: "vpn_host_encrypt")
        let _vpn_host = defaults.string(forKey: "vpn_host") ?? ""
        let _vpn_username = defaults.string(forKey: "vpn_username") ?? ""
        let _vpn_paswd = defaults.string(forKey: "vpn_passwd") ?? ""

        if _socks5_addr.trimmingCharacters(in: .whitespaces) != "" {
            self.uiSocks5Addr = _socks5_addr
        }
        if _remote_addr.trimmingCharacters(in: .whitespaces) != "" {
            self.uiRemoteAddr = _remote_addr
        }
        if _http_addr.trimmingCharacters(in: .whitespaces) != "" {
            self.uiHttpAddr = _http_addr
        }

        if _vpn_host.trimmingCharacters(in: .whitespaces) != "" {
            self.uiVPNHost = _vpn_host
        }
        if _vpn_username.trimmingCharacters(in: .whitespaces) != "" {
            self.uiVPNUsername = _vpn_username
        }
        if _vpn_paswd != "" {
            self.uiVPNPassword = _vpn_paswd
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
        Group {
            ContentView()
        }
    }
}
