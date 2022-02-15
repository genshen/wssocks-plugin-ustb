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
    @ObservedObject var config: Configs

    @State private var uiEnableSubmitBtn = true
    @State private var uiSubmitBtnLabel = "Start"

    @State private var showingAlert = false
    @State private var alertMessage: String = ""

    init(conf: Configs) {
        self.config = conf
    }

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
    }

    func onSubmit() {
        if !uiEnableSubmitBtn {
            return
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
        ContentView(conf: Configs())
    }
}
