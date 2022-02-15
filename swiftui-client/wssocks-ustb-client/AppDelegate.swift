//
//  wssocks_ustb_clientApp.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2020/4/30.
//  Copyright © 2020 genshen. All rights reserved.
//  Updated by genshen on 2021/8/19 and 2022/1/14.
//

import SwiftUI

@main
struct wssocks_ustb_clientApp: App {
    @ObservedObject var config = Configs()

    @NSApplicationDelegateAdaptor(AppDelegate.self) var delegate

    init() {
        self.delegate.config = config
    }

    var body: some Scene {
        // to hiden unnecessary view,
        // see https://khorbushko.github.io/article/2021/04/30/minimal-macOS-menu-bar-extra%27s-app-with-SwiftUI.html
        Settings {
            ContentView(conf: config).frame(width: 450)
        }
    }
}

class AppDelegate: NSObject, NSApplicationDelegate {
    var config: Configs!
    var eventMonitor: EventMonitor?

    var popOver = NSPopover()

    private var statusBar: NSStatusBar!
    private var statusItem: NSStatusItem!

    @objc func onStatusBarToggle(_ sender: Any?) {
        if popOver.isShown {
            popOver.performClose(sender)
            self.eventMonitor?.stop()
        }else {
            if let statusBarButton = statusItem.button {
                self.popOver.show(relativeTo: statusBarButton.bounds, of: statusBarButton, preferredEdge: NSRectEdge.minY)
                self.eventMonitor?.start()
            }
        }
    }

    func applicationDidFinishLaunching(_ aNotification: Notification) {
        // set status bar
        statusBar = NSStatusBar.system
        statusItem = statusBar.statusItem(withLength: NSStatusItem.squareLength)
        if let statusBarButton = statusItem.button {
            statusBarButton.image = NSImage(named: "StatusIcon")
            statusBarButton.image?.isTemplate = true
            statusBarButton.action = #selector(onStatusBarToggle(_:))
        }

        // set main popover
        popOver.behavior = .transient
        popOver.animates = true

        var menuView = MenuBarView()
        menuView.statusItem = statusItem
        menuView.configsInView = config

        popOver.contentViewController = NSViewController()
        popOver.contentViewController?.view = NSHostingView(rootView: menuView)

        // add event listerner to auto close popover
        eventMonitor = EventMonitor(mask: [.leftMouseDown, .rightMouseDown]) { [weak self] event in
            if let strongSelf = self, strongSelf.popOver.isShown {
                self!.popOver.performClose(event!)
                self!.eventMonitor?.stop()
            }
        }
    }

    func applicationWillTerminate(_ aNotification: Notification) {
        // Insert code here to tear down your application
        config.StoreUserDefaults(defaults: UserDefaults.standard)
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        false
    }
}
