//
//  wssocksClientApi.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2020/5/1.
//  Copyright © 2020 genshen. All rights reserved.
//

import Foundation
import WssocksGoApi

typealias UintPtr = UInt

struct WssocksClient {
    private var handle: UintPtr

    init() {
        handle = NewClientHandles()
    }

    public func startClient(config: Configs) -> String? {
        let CStrSocks5Addr = config.uiSocks5Addr.cString(using: String.Encoding.utf8)
        let CStrRemoteAddr = config.uiRemoteAddr.cString(using: String.Encoding.utf8)
        let CStrHttpAddr = config.uiHttpAddr.cString(using: String.Encoding.utf8)

        let CStrVPNHost = config.uiVPNHost.cString(using: String.Encoding.utf8)
        let CStrVPNUsername = config.uiVPNUsername.cString(using: String.Encoding.utf8)
        let CStrVPNPasswd = config.uiVPNPassword.cString(using: String.Encoding.utf8)

        let socks5AddrPtr = UnsafeMutablePointer(mutating: CStrSocks5Addr)
        let remoteAddrPtr = UnsafeMutablePointer(mutating: CStrRemoteAddr)
        let httpAddrPtr = UnsafeMutablePointer(mutating: CStrHttpAddr)

        let vpnHostPtr = UnsafeMutablePointer(mutating: CStrVPNHost)
        let vpnUsernamePtr = UnsafeMutablePointer(mutating: CStrVPNUsername)
        let vpnPasswdPtr = UnsafeMutablePointer(mutating: CStrVPNPasswd)
    
        guard let v = StartClientWrapper(self.handle, socks5AddrPtr, remoteAddrPtr, httpAddrPtr,
                                         config.uiEnableHttpProxy, config.uiSkipTSLerify,
                                         config.uiVPNEnable, config.uiVPNForceLogout, config.uiVPNHostEncrypt,
                                         vpnHostPtr, vpnUsernamePtr, vpnPasswdPtr) else { return nil }
        return String(bytesNoCopy: v, length: strlen(v), encoding: .utf8, freeWhenDone: true)
    }

    public func waitClient() -> String? {
        guard let v = WaitClientWrapper(self.handle) else { return nil }
        return String(bytesNoCopy: v, length: strlen(v), encoding: .utf8, freeWhenDone: true)
    }

    public func stopClient() -> String? {
        guard let v = StopClientWrapper(self.handle) else { return nil }
        return String(bytesNoCopy: v, length: strlen(v), encoding: .utf8, freeWhenDone: true)
    }
}
