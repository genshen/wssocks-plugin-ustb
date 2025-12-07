package webview

var presetChromiumExecPath = []string{
	"/usr/bin/chromium-browser", // Linux
	"/usr/bin/chromium",         // Linux
	"C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",        // Windows Chrome
	"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",  // Windows Chrome (32 bit)
	"C:\\Program Files\\Chromium\\Application\\chrome.exe",              // Windows Chromium
	"C:\\Program Files (x86)\\Chromium\\Application\\chrome.exe",        // Windows Chromium (32 bit)
	"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe", // MS Edge
	// for macOS, the path is copied from:
	// https://docs.rs/headless_chrome/latest/x86_64-pc-windows-msvc/src/headless_chrome/browser/mod.rs.html#497-56
	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", // MacOS
	"/Applications/Google Chrome Beta.app/Contents/MacOS/Google Chrome Beta",
	"/Applications/Google Chrome Dev.app/Contents/MacOS/Google Chrome Dev",
	"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
	"/Applications/Chromium.app/Contents/MacOS/Chromium",
	"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
	"/Applications/Microsoft Edge Beta.app/Contents/MacOS/Microsoft Edge Beta",
	"/Applications/Microsoft Edge Dev.app/Contents/MacOS/Microsoft Edge Dev",
	"/Applications/Microsoft Edge Canary.app/Contents/MacOS/Microsoft Edge Canary",
}
