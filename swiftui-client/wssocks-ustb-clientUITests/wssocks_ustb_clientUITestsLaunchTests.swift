//
//  wssocks_ustb_clientUITestsLaunchTests.swift
//  wssocks-ustb-clientUITests
//
//  Created by genshen on 2022/2/14.
//  Copyright © 2020-present genshen. All rights reserved.
//

import XCTest

class wssocks_ustb_clientUITestsLaunchTests: XCTestCase {

    override class var runsForEachTargetApplicationUIConfiguration: Bool {
        true
    }

    override func setUpWithError() throws {
        continueAfterFailure = false
    }

    func testLaunch() throws {
        let app = XCUIApplication()
        app.launch()

        // Insert steps here to perform after app launch but before taking a screenshot,
        // such as logging into a test account or navigating somewhere in the app

        let attachment = XCTAttachment(screenshot: app.screenshot())
        attachment.name = "Launch Screen"
        attachment.lifetime = .keepAlways
        add(attachment)
    }
}
