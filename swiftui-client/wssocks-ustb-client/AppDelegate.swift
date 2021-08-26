//
//  AppDelegate.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2020/4/30.
//  Copyright © 2020 genshen. All rights reserved.
//  Updated by genshen on 2021/8/19.
//

import Cocoa
import SwiftUI

@main
class AppDelegate: NSObject, NSApplicationDelegate {

    private var menuExtrasConfigurator: MacExtrasConfigurator?

    var window: NSWindow!
    var contentView: ContentView!

    func applicationDidFinishLaunching(_ aNotification: Notification) {
        // Create the SwiftUI view that provides the window contents.
        menuExtrasConfigurator = .init()
        
        contentView = ContentView()

        // Create the window and set the content view.
        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: 480, height: 300),
            styleMask: [.titled, .closable, .miniaturizable, .resizable, .fullSizeContentView],
            backing: .buffered, defer: false)
        menuExtrasConfigurator?.window = window
        menuExtrasConfigurator?.configsInView = contentView.config

        window.title = "Preference"
        window.isReleasedWhenClosed = false
        window.center()
        window.setFrameAutosaveName("Preference")
        window.contentView = NSHostingView(rootView: contentView)
        // window.makeKeyAndOrderFront(nil)
    }

    func applicationWillTerminate(_ aNotification: Notification) {
        // Insert code here to tear down your application
        contentView.config.StoreUserDefaults(defaults: UserDefaults.standard)
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        false
    }
    
}
