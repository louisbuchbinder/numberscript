package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/louisbuchbinder/core/lib/util/testutil"
)

func testHash(t *testing.T, testcases []hashTestcase) {
	ctx, cancel := chromedp.NewContext(h.ChromedpContext)
	defer cancel()
	for _, tc := range testcases {
		tc := tc
		inputId := fmt.Sprintf("%s-input-0", tc.tab)
		resultId := fmt.Sprintf("%s-result-0", tc.tab)
		toggleId := fmt.Sprintf("%s-toggle", tc.tab)

		err := chromedp.Run(
			ctx,
			chromedp.Navigate(fmt.Sprintf("http://localhost:%d/hash/%s/", h.ServerPort, tc.pkg)),
			chromedp.WaitReady("body", chromedp.ByQuery),
			chromedp.WaitReady(toggleId, chromedp.ByID),
			chromedp.Click(toggleId, chromedp.ByID),
			chromedp.WaitVisible(inputId, chromedp.ByID),
			chromedp.WaitVisible(resultId, chromedp.ByID),
			chromedp.WaitReady(inputId, chromedp.ByID),
			chromedp.WaitReady(resultId, chromedp.ByID),
			chromedp.ActionFunc(func(actx context.Context) error {
				in, err := GetElementById(actx, inputId)
				if err != nil {
					return err
				}
				if err := chromedp.KeyEventNode(in, tc.in).Do(actx); err != nil {
					return err
				}
				return nil
			}),
			chromedp.Sleep(time.Millisecond*10),
			chromedp.ActionFunc(func(actx context.Context) error {
				var output string
				if err := chromedp.TextContent(resultId, &output, chromedp.ByID).Do(actx); err != nil {
					return err
				}
				if tc.out != output {
					return fmt.Errorf("%s %s failed, input='%s' expected='%s', actual='%s'\n", tc.pkg, tc.tab, tc.in, tc.out, output)
				}
				return nil
			}),
		)
		testutil.AssertNilError(t, err)
	}
}
