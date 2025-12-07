package webview

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
)

type ChromedpWebview struct {
}

func (chr *ChromedpWebview) ShowWebviewAndSetCookies(client *http.Client, loginUrlStr string) ([]*http.Cookie, error) {
	loginUrl, err := url.Parse(loginUrlStr)
	if err != nil {
		return nil, err
	}
	if vpnCookies, err := chr.launchWebview(loginUrl, presetChromiumExecPath); err != nil {
		return nil, err
	} else {
		return vpnCookies, nil
	}
}

// launchWebview launches a chromium view for debug. Which can get the cookie of the visited website.
func (chr *ChromedpWebview) launchWebview(vpnLoginUrl *url.URL, suggestedChromiumPaths []string) ([]*http.Cookie, error) {
	var chromiumPath string
	execFind := false
	for _, path := range suggestedChromiumPaths {
		if _, err := os.Stat(path); err == nil {
			chromiumPath = path
			execFind = true
			break
		}
	}
	if !execFind {
		// fallback to the preset install path
		for _, path := range presetChromiumExecPath {
			if _, err := os.Stat(path); err == nil {
				chromiumPath = path
				execFind = true
				break
			}
		}
	}
	if !execFind {
		return nil, fmt.Errorf("could not find chromium based browser")
	}

	log.Printf("use Chromium: %s\n", chromiumPath)

	userDataDir := filepath.Join(os.TempDir(), "chromedp-wssocks-cookies")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(chromiumPath),
		chromedp.UserDataDir(userDataDir),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-web-security", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Println(vpnLoginUrl.String()) // // ///
	var vpnCookies []*http.Cookie
	var noticeResult bool
	// create task list
	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return network.Enable().Do(ctx)
		}),
		chromedp.Navigate(vpnLoginUrl.String()),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Evaluate(addNotice(), &noticeResult),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			var allCookies []*network.Cookie
			allCookies, err = storage.GetCookies().Do(ctx)
			if err != nil {
				return fmt.Errorf("get cookies failed with %w", err)
			}

			// filter HttpOnly cookies
			for _, cookie := range allCookies {
				if strings.Contains(cookie.Domain, vpnLoginUrl.Host) {
					vpnCookies = append(vpnCookies, convertChromedpCookieToHTTPCookie(cookie))
				}
			}
			return nil
		}),
	)

	if err != nil {
		return nil, err
	}

	<-ctx.Done()

	log.Println("browser window closed.")

	return vpnCookies, nil
}

func addNotice() string {
	jsCode := fmt.Sprintf(`
    (function(m) {
        const oldToast = document.getElementById('chromedp-toast');
        if (oldToast) oldToast.remove();
        
        const toast = document.createElement('div');
        toast.id = 'chromedp-toast';
        toast.textContent = m;
        Object.assign(toast.style, {
            position: 'fixed',
            top: '20px',
            right: '20px',
            background: '#323232',
            color: '#fff',
            padding: '12px 24px',
            borderRadius: '4px',
            boxShadow: '0 3px 10px rgba(0,0,0,0.3)',
            zIndex: '2147483647',
            fontFamily: 'system-ui, sans-serif',
            fontSize: '14px',
            transition: 'opacity 0.3s'
        });
        document.body.appendChild(toast);
        return true;
    })("%s");
    `, "登录VPN后，关闭浏览器窗口或页面即可返回 wssocks-ustb，实现自动登录")
	return jsCode
}
