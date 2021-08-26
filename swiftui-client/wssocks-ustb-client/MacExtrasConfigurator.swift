//
//  MacExtrasConfigurator.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2021/8/18.
//

import Foundation
import SwiftUI

final class MacExtrasConfigurator: NSObject {
    
    private var statusBar: NSStatusBar
    private var statusItem: NSStatusItem

    var window: NSWindow! // ref of window
    var configsInView: Configs! // ref of config values in ContentView

    override init() {
        statusBar = NSStatusBar.system
        statusItem = statusBar.statusItem(withLength: NSStatusItem.squareLength)
        super.init()
        
        createMenu()
    }

    private func createMenu() {
        if let statusBarButton = statusItem.button {
            statusBarButton.image = NSImage(named: "StatusIcon")
            statusBarButton.image?.isTemplate = true

            let toggleMenuItem = NSMenuItem(title: "开启wssocks", action: #selector(Self.toggleWssocks(_:)), keyEquivalent: "")
            toggleMenuItem.target = self
            // see: https://hetima.github.io/fucking_nsimage_syntax/
            toggleMenuItem.image = NSImage(named: NSImage.statusNoneName)

            let sysProxyMenuItem = NSMenuItem(title: "打开系统代理设置", action: #selector(Self.openNetworkProxyPreferences(_:)), keyEquivalent: "")
            sysProxyMenuItem.target = self

            let prefMenuItem = NSMenuItem(title: "偏好设置", action: #selector(Self.showPref(_:)), keyEquivalent: ",")
            prefMenuItem.target = self
            
            let githubItem = NSMenuItem(title: "Github", action: #selector(Self.showGithubAction(_:)), keyEquivalent: "")
            githubItem.target = self
            
            let helpItem = NSMenuItem(title: "帮助", action: #selector(Self.showHelpAction(_:)), keyEquivalent: "")
            helpItem.target = self
            
            let aboutItem = NSMenuItem(title: "关于", action: #selector(Self.showAboutAction(_:)), keyEquivalent: "")
            aboutItem.target = self
            
            let quitItem = NSMenuItem(title: "退出", action: #selector(Self.quitApp(_:)), keyEquivalent: "q")
            quitItem.target = self

            let mainMenu = NSMenu()
            mainMenu.addItem(toggleMenuItem)
            mainMenu.addItem(NSMenuItem.separator())
            mainMenu.addItem(sysProxyMenuItem)
            mainMenu.addItem(NSMenuItem.separator())
            mainMenu.addItem(prefMenuItem)
            mainMenu.addItem(githubItem)
            mainMenu.addItem(helpItem)
            mainMenu.addItem(aboutItem)
            mainMenu.addItem(NSMenuItem.separator())
            mainMenu.addItem(quitItem)
            
            statusItem.menu = mainMenu
        }
    }

    @objc private func openNetworkProxyPreferences(_ sender: NSMenuItem) {
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

    private let defaults = UserDefaults.standard
    var wssocksStatus = 0
    private var client = WssocksClient()

    func checkPref() -> Bool {
        if self.configsInView.uiVPNPassword == "" {
            return false
        }
        return true
    }

    private func noticeFailed(message: String) {
        let notification = NSUserNotification()
        notification.identifier = "unique-id" + String(Int(Date().timeIntervalSince1970))
        notification.title = "开启 wssocks 失败"
        notification.informativeText = message
        notification.soundName = NSUserNotificationDefaultSoundName
        let notificationCenter = NSUserNotificationCenter.default
        notificationCenter.deliver(notification)
    }

    @objc private func toggleWssocks(_ sender: NSMenuItem) {
        if wssocksStatus == 0 {
            if(checkPref() == false) {
                noticeFailed(message: "No VPN password")
                return
            }

            sender.title = "正在开启wssocks..."
            sender.target = nil
            DispatchQueue.global().async {
                let msg = self.client.startClient(config: self.configsInView) ?? ""
                DispatchQueue.main.sync {
                    if msg != "" {
                        self.noticeFailed(message: msg)
                        sender.title = "开启wssocks"
                        sender.image = NSImage(named: NSImage.statusNoneName)
                    } else {
                        sender.title = "停止wssocks"
                        sender.image = NSImage(named: NSImage.statusAvailableName)
                        self.statusItem.image = NSImage(named: "LaunchIcon")
                        self.wssocksStatus = 1
                    }
                    sender.target = self
                }
            }
        } else {
            sender.target = nil
            sender.title = "正在停止wssocks..."
            DispatchQueue.global().async {
                let msg = self.client.stopClient() ?? ""
                DispatchQueue.main.sync {
                    if msg != "" {
                        self.noticeFailed(message: msg)
                    }
                    sender.target = self
                    sender.title = "开启wssocks"
                    sender.image = NSImage(named: NSImage.statusNoneName)
                    self.statusItem.image = NSImage(named: "StatusIcon")
                    self.wssocksStatus = 0
                }
            }
        }
    }
    
    @objc private func showGithubAction(_ sender: Any?) {
        guard let url = URL(string: "https://github.com/genshen/wssocks-plugin-ustb") else {
            return
        }
        NSWorkspace.shared.open(url)
    }
    
    @objc private func showHelpAction(_ sender: Any?) {
        guard let url = URL(string: "https://genshen.github.io/wssocks-plugin-ustb") else {
            return
        }
        NSWorkspace.shared.open(url)
    }

    @objc private func showPref(_ sender: Any?) {
        window.makeKeyAndOrderFront(nil)
        NSApp.activate(ignoringOtherApps: true)
    }
    
    @objc private func quitApp(_ sender: NSMenuItem) {
        NSApplication.shared.terminate(self)
    }
    
    @objc private func showAboutAction(_ sender: Any?) {
        NSApp.orderFrontStandardAboutPanel(sender);
        NSApp.activate(ignoringOtherApps: true)
    }
}
