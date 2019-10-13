package fpl

import (
	"context"
	"io/ioutil"
	"log"
	"math"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// opts := []chromedp.ExecAllocatorOption{
// 	chromedp.NoFirstRun,
// 	chromedp.NoDefaultBrowserCheck,
// 	chromedp.Headless,
// 	chromedp.DisableGPU,
// 	chromedp.ExecPath("/my/directory/headless-shell/headless-shell"),
// }
// allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...); defer cancel()
// ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf)); defer cancel()

// c, err := cdp.New(ctxt, cdp.WithRunnerOptions(cdpr.Flag("disable-web-security", "1")))

// https://github.com/chromedp/chromedp/issues/109

// FROM chromedp/headless-shell:latest
// COPY myapp /myapp/
// ENV PATH=$PATH:/headless-shell
// WORKDIR /myapp
// EXPOSE 3000
// ENTRYPOINT [ "myapp", "-param1=" ]

func (api *FPL) GetScreenshot(url string, element string) {
	// create context

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
		//chromedp.ExecPath("/my/directory/headless-shell/headless-shell"),
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	//ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	log.Println("STARTING: chromedp")
	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(`https://fantasy.premierleague.com/entry/4719576/event/3`, `div.Pitch-sc-1mctasb-0.kppgUD`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("elementScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://fantasy.premierleague.com/entry/4719576/event/1`, 100, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
