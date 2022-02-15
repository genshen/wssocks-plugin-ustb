//
//  Configs.swift
//  wssocks-ustb-client
//
//  Created by genshen on 2022/2/14.
//  Copyright © 2020 - present genshen. All rights reserved.
//

import Foundation

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

    init() {
        // load preference
        let defaults = UserDefaults.standard
        self.LoadUserDefaults(defaults: defaults)
    }

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
