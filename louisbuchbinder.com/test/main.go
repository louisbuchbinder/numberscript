package test

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/louisbuchbinder/core/lib/util"
)

var serverExeRlocation string

func initServer(ctx context.Context) error {
	exe, err := runfiles.Rlocation(serverExeRlocation)
	if err != nil {
		return err
	}
	p, err := util.FreePort()
	if err != nil {
		return err
	}
	h.ServerPort = p
	cmd := exec.CommandContext(ctx, exe, "-port", fmt.Sprintf("%d", h.ServerPort))
	if err := cmd.Start(); err != nil {
		return err
	}
	h.CleanupFunctionWrapper.Add(cmd.Cancel)
	ticker := time.NewTicker(time.Second * 3)
LOOP:
	for {
		select {
		case <-ticker.C:
			return fmt.Errorf("failed to start the server")
		default:
			resp, err := http.Get(fmt.Sprintf("http://localhost:%d/", h.ServerPort))
			if err == nil && resp.StatusCode == http.StatusOK {
				// nil error with 200 status code means connection successful so break
				break LOOP
			}
		}
	}
	return nil
}

func initChromedp(ctx context.Context) error {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[3:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	)
	ectx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	h.CleanupFunctionWrapper.AddCancel(cancel)

	cdpParentContext, cancel := chromedp.NewContext(
		ectx,
		chromedp.WithBrowserOption(chromedp.WithDialTimeout(time.Second)),
		// chromedp.WithDebugf(log.Printf),
	)
	h.CleanupFunctionWrapper.AddCancel(cancel)
	h.ChromedpContext = cdpParentContext

	err := chromedp.Run(h.ChromedpContext, chromedp.ActionFunc(func(ctx context.Context) error {
		c := chromedp.FromContext(ctx)
		_, err := target.CreateBrowserContext().Do(cdp.WithExecutor(ctx, c.Browser))
		return err
	}))
	if err != nil {
		return err
	}

	return nil
}

type harness struct {
	ServerPort             int
	ChromedpContext        context.Context
	CleanupFunctionWrapper util.CleanupFunctionWrapper
}

var h harness

func GetElementById(ctx context.Context, id string) (*cdp.Node, error) {
	nodes := make([]*cdp.Node, 1)
	if err := chromedp.Nodes(id, &nodes, chromedp.ByID).Do(ctx); err != nil {
		return nil, err
	}

	if len(nodes) < 1 {
		return nil, fmt.Errorf("%s not found", id)
	}
	return nodes[0], nil
}
