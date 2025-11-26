package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/lib/util/testutil"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	util.Must0(initServer(ctx))
	util.Must0(initChromedp(ctx))
	cleanup := h.CleanupFunctionWrapper.Cleanup()
	code := m.Run()
	if err := cleanup(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if code == 0 {
			code = 1
		}
	}
	os.Exit(code)
}

func TestLoadHomePage(t *testing.T) {
	err := chromedp.Run(h.ChromedpContext,
		chromedp.Navigate(fmt.Sprintf("http://localhost:%d/", h.ServerPort)),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)
	testutil.AssertNilError(t, err)
}

func TestEncodingBase32(t *testing.T) {
	testEncoding(t, encodeTestcasesBase32)
}

func TestEncodingBase64(t *testing.T) {
	testEncoding(t, encodeTestcasesBase64)
}

func TestEncodingHex(t *testing.T) {
	testEncoding(t, encodeTestcasesHex)
}

func TestEncodingHtml(t *testing.T) {
	testEncoding(t, encodeTestcasesHtml)
}

func TestEncodingUri(t *testing.T) {
	testEncoding(t, encodeTestcasesUri)
}

func TestHashAdler32(t *testing.T) {
	testHash(t, hashTestcasesAdler32)
}
