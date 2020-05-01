//
//  AppDelegate.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2020/4/30.
//  Copyright © 2020 genshen. All rights reserved.
//

import Cocoa
import SwiftUI

@NSApplicationMain
class AppDelegate: NSObject, NSApplicationDelegate {

    var window: NSWindow!
    var contentView: ContentView!

    func applicationDidFinishLaunching(_ aNotification: Notification) {
        // Create the SwiftUI view that provides the window contents.
        contentView = ContentView()
        contentView.LoadUserDefaults()
        let rootView = contentView.frame(minWidth: 300)

        // Create the window and set the content view. 
        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: 320, height: 480),
            styleMask: [.titled, .closable, .miniaturizable, /*.resizable,*/ .fullSizeContentView],
            backing: .buffered, defer: false)
        window.title = "wssocks Client"
        window.center()
        window.setFrameAutosaveName("Main Window")
        window.contentView = NSHostingView(rootView: rootView)
        window.makeKeyAndOrderFront(nil)
    }

    func applicationWillTerminate(_ aNotification: Notification) {
        // Insert code here to tear down your application
        contentView.StoreUserDefaults()
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        true
    }

}


struct AppDelegate_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
