package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/lib/util/testutil"
)

type testcaseInput struct {
	content  string
	operator string
}

type testcase struct {
	id          uuid.UUID
	module      string
	pkg         string
	tab         string
	argOperator string
	in          []testcaseInput
	out         []string
}

func testStd(t *testing.T, testcases []testcase) {
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprintf("%s-%s", t.Name(), tc.id.String()), func(tt *testing.T) {
			ctx, cancel := context.WithTimeout(h.ChromedpContext, time.Second*3)
			defer cancel()
			ctx, cancel = chromedp.NewContext(ctx)
			defer cancel()
			toggleId := fmt.Sprintf("%s-toggle", tc.tab)

			err := chromedp.Run(
				ctx,
				chromedp.Navigate(fmt.Sprintf("http://localhost:%d/%s/%s/", h.ServerPort, tc.module, tc.pkg)),
				chromedp.Sleep(time.Millisecond*10),
				chromedp.WaitReady("body", chromedp.ByQuery),
				chromedp.WaitReady(toggleId, chromedp.ByID),
				chromedp.Click(toggleId, chromedp.ByID),
				chromedp.ActionFunc(func(actx context.Context) error {
					for i, in := range tc.in {
						operatorToggleId := fmt.Sprintf("%s-data-%s-toggle", tc.tab, in.operator)
						if err := chromedp.WaitVisible(operatorToggleId, chromedp.ByID).Do(actx); err != nil {
							return err
						}
						if err := chromedp.Click(operatorToggleId, chromedp.ByID).Do(actx); err != nil {
							return err
						}
						if err := chromedp.WaitVisible(fmt.Sprintf("%s-input-%d", tc.tab, i), chromedp.ByID).Do(actx); err != nil {
							return err
						}
					}
					return nil
				}),
				chromedp.ActionFunc(func(actx context.Context) error {
					for i := range tc.out {
						if err := chromedp.WaitVisible(fmt.Sprintf("%s-result-%d", tc.tab, i), chromedp.ByID).Do(actx); err != nil {
							return err
						}
					}
					return nil
				}),
				chromedp.ActionFunc(func(actx context.Context) error {
					for i, input := range tc.in {
						in, err := GetElementById(actx, fmt.Sprintf("%s-input-%d", tc.tab, i))
						if err != nil {
							return err
						}
						if err := chromedp.KeyEventNode(in, input.content).Do(actx); err != nil {
							return err
						}
					}
					return nil
				}),
				chromedp.Sleep(time.Millisecond*10),
				chromedp.ActionFunc(func(actx context.Context) error {
					for i, output := range tc.out {
						var value string
						if err := chromedp.TextContent(fmt.Sprintf("%s-result-%d", tc.tab, i), &value, chromedp.ByID).Do(actx); err != nil {
							return err
						}
						if output != value {
							return fmt.Errorf(
								"%s %s failed, input=[%s] expected='%s', actual='%s'\n",
								tc.pkg,
								tc.tab,
								strings.Join(util.Map(tc.in, func(_ int, in testcaseInput) string { return in.content }), ", "),
								output,
								value,
							)
						}
					}
					return nil
				}),
			)
			testutil.AssertNilError(tt, err)
		})
	}
}
