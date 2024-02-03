//
//  MenuBarView.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2022/2/1.
//

import SwiftUI

struct MenuBarView: View {
    @Namespace var animation
    @State private var showingAlert = false
    @State private var alertMessage: String = ""
    @State private var btnInProgress = false

    @State private var statusDesc: String = ""
    @State private var clickI: Int = 0

    var statusItem: NSStatusItem! // ref of status item
    var configsInView: Configs! // ref of config values in ContentView

    @State var ignoreWaitErr = true
    private let defaults = UserDefaults.standard
    @State var wssocksStatus = 0
    var client = WssocksClient()

    var body: some View {
        VStack{
            HStack{
                if #available(macOS 14.0, *) {
                    SettingsLink{
                        Image(systemName: "gear.circle.fill")
                    }
                } else {
                    Button (action:{
                        showPref()
                    }, label: {
                        if #available(macOS 11.0, *) {
                            Image(systemName: "gear.circle.fill")
                        } else {
                            Text("偏好设置")
                        }
                    })
                }

                Button (action:{
                    openNetworkProxyPreferences()
                }, label: {
                    if #available(macOS 11.0, *) {
                        Image(systemName: "network")
                    } else {
                        Text("系统代理设置")
                    }
                })
                Spacer()
                Button (action:{
                        showAboutAction(nil)
                }, label: {
                    if #available(macOS 11.0, *) {
                        Image(systemName: "info.circle")
                    } else {
                        Text("关于")
                    }
                })
            }
            .padding(.top, 8)
            .padding(.horizontal, 8)
            Divider()//.padding(.top, 4)

            HStack{
                Button(action: toggleWssocks, label: {
                    Text(wssocksStatus == 0 ? "Start": "Stop")
                        .font(.callout)
                        .fontWeight(.bold)
                        .foregroundColor(.white)
                        .padding(.vertical, 4)
                        .frame(maxWidth: .infinity)
                }).buttonStyle(.borderedProminent)
                    .controlSize(.large)
                    .background(ZStack{
                        if !btnInProgress {
                            RoundedRectangle(cornerRadius: 4)
                                .fill(Color.accentColor)
                                .matchedGeometryEffect(id: "Start", in: animation)
                        }
                    })
                    .buttonBorderShape(.roundedRectangle)
                    .disabled(btnInProgress)
                    .alert(isPresented: $showingAlert) {
                            Alert(title: Text("Error"), message: Text("\(alertMessage)"), dismissButton: .default(Text("OK")))
                    }
            }
            .padding(.horizontal)
           // .padding(.top)

           // Divider().padding(.top, 4)

            if #available(macOS 11.0, *) {
                Button(action: {
                    self.clickI = clickI + 1
                }) {
                    Image(systemName: wssocksStatus==0 ? "bolt.slash": (clickI % 5 == 4 ? "circle.hexagongrid.fill": "bolt"))
                        .resizable()
                        .aspectRatio(contentMode: .fit)
                        .foregroundColor(wssocksStatus==0 ? Color.primary: Color.accentColor)
                        .padding((wssocksStatus == 1 && clickI % 5 == 4) ? 12: 24)
                        .symbolRenderingMode(.multicolor)
                }.buttonStyle(PlainButtonStyle())
                if wssocksStatus == 1 && clickI % 5 == 4 {
                    Text("Oh! This is an easter egg!")
                }
            } else {
                // Fallback on earlier versions
            }

            // quit button
            if #available(macOS 11.0, *) {
                Button(action: quitApp) {
                    Label("Quit App", systemImage: "xmark.circle.fill")
                        //.foregroundColor(Color.accentColor)
                }.buttonStyle(BorderlessButtonStyle())
            } else {
                // Fallback on earlier versions
            }

            Divider().padding(.top, 4)

            // bottom view
            HStack{
                Button (action:{
                    showGithubAction(nil)
                }, label: {
                    if #available(macOS 11.0, *) {
                        Label("Github", systemImage: "heart.circle.fill")
                            .symbolRenderingMode(.multicolor)
                    } else {
                        Text("Github")
                    }
                })
                Spacer()
                Button (action:{
                    showHelpAction(nil)
                }, label: {
                    if #available(macOS 11.0, *) {
                        Label("帮助", systemImage: "questionmark.circle")
                    } else {
                        Text("帮助")
                    }
                })
            }
            .padding(.bottom, 4)
            .padding(.horizontal, 8)
        }
        .frame(width: 250, height: 250)
    }

    func checkPref() -> Bool {
        if self.configsInView.uiVPNPassword == "" {
            return false
        }
        return true
    }

    private func openNetworkProxyPreferences() {
//        let url = URL(string:"x-apple.systempreferences:com.apple.preference.network?Proxies")!
//        NSWorkspace.shared.open(url)
        if #available(macOS 13, *) {
            // see https://github.com/bvanpeski/SystemPreferences/blob/main/macos_preferencepanes-Ventura.md#network
            let url = URL(string:"x-apple.systempreferences:com.apple.Network-Settings.extension?Proxies")!
            NSWorkspace.shared.open(url)
        }else{
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
   }

    private func showGithubAction(_ sender: Any?) {
       guard let url = URL(string: "https://github.com/genshen/wssocks-plugin-ustb") else {
           return
       }
       NSWorkspace.shared.open(url)
   }

    private func noticeFailed(message: String) {
        showingAlert = true
        // todo: notification.title = "开启 wssocks 失败"
        alertMessage = message
    }

    private func toggleWssocks() {
        if wssocksStatus == 0 {
            if(checkPref() == false) {
                noticeFailed(message: "No VPN password")
                return
            }

            statusDesc = "正在开启wssocks..."
            btnInProgress = true
            DispatchQueue.global().async {
                let msg = self.client.startClient(config: self.configsInView) ?? ""
                DispatchQueue.main.sync {
                    btnInProgress = false
                    if msg != "" {
                        self.noticeFailed(message: msg)
                        self.setWssocksStatusUI(status: 0)
                    } else {
                        self.setWssocksStatusUI(status: 1)
                    }
                }
                if msg != "" { // starting has error, skip waiting
                    return
                }

                self.ignoreWaitErr = false // now we can accept wait error
                let errorMsg = self.client.waitClient() ?? ""
                DispatchQueue.main.sync {
                    // we only log error when it task is finished with error (not button clicking by user).
                    if errorMsg != "" && !ignoreWaitErr {
                        self.noticeFailed(message: errorMsg)
                    }
                    // waiting finished, set the status as "stopped".
                    self.setWssocksStatusUI(status: 0)
                    self.ignoreWaitErr = true
                }
            }
        } else {
            statusDesc = "正在停止wssocks..."
            btnInProgress = true
            self.ignoreWaitErr = true
            DispatchQueue.global().async {
                let msg = self.client.stopClient() ?? ""
                DispatchQueue.main.sync {
                    btnInProgress = false
                    if msg != "" {
                        self.noticeFailed(message: msg)
                    }
                    self.setWssocksStatusUI(status: 0)
                }
            }
        }
    }

    private func setWssocksStatusUI(status: Int) {
        self.wssocksStatus = status
        if status == 0 {
            statusDesc = "点击以开启wssocks"
            statusItem.image = NSImage(named: "StatusIcon")
        } else if status == 1 {
            statusDesc = "点击以停止wssocks"
            statusItem.image = NSImage(named: "LaunchIcon")
        } else {

        }
    }

    private func showHelpAction(_ sender: Any?) {
        guard let url = URL(string: "https://genshen.github.io/wssocks-plugin-ustb") else {
            return
        }
        NSWorkspace.shared.open(url)
    }

    @Environment(\.openURL) var openURL
    private func showPref() {
        // see: https://stackoverflow.com/a/69600396/10068476
        NSApp.activate(ignoringOtherApps: true)
        // see: https://stackoverflow.com/a/65356627/10068476
        if #available(macOS 13, *) {
            NSApp.sendAction(Selector(("showSettingsWindow:")), to: nil, from: nil)
        } else {
            NSApp.sendAction(Selector(("showPreferencesWindow:")), to: nil, from: nil)
        }
        NSApp.windows.first?.orderFrontRegardless()
    }

    private func quitApp() {
       NSApplication.shared.terminate(self)
    }

    private func showAboutAction(_ sender: Any?) {
       NSApp.orderFrontStandardAboutPanel(sender);
       NSApp.activate(ignoringOtherApps: true)
   }
}

struct MenuBarView_Previews: PreviewProvider {
    static var previews: some View {
        MenuBarView()
    }
}

struct TabButton: View {
    var title: String
    var body: some View{
        Button(action: {}, label: {
            Text(title)
                .font(.callout)
                .fontWeight(.bold)
                .foregroundColor(.primary)
                .padding(.vertical, 4)
                .frame(maxWidth: .infinity)
                .background(
                    ZStack{
                        RoundedRectangle(cornerRadius: 4)
                            .stroke(Color.primary)
//                            .fill(Color.blue)
                    }
                )
        })
            .buttonStyle(PlainButtonStyle())
    }
}
