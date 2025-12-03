package load

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/louisbuchbinder/core/lib/build"
	"github.com/louisbuchbinder/core/lib/util"
)

var initializers = []func(){}

func Register(initializer func()) int {
	initializers = append(initializers, initializer)
	return 0
}

func Load() {
	for _, fn := range initializers {
		fn()
	}
}

var (
	manifestPath string
	ManifestData = map[string]build.ManifestData{}
)

func init() {
	flag.StringVar(&manifestPath, "manifest", "", "[required] the manifest path")
	flag.Parse()
	if manifestPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	b := util.Must(os.ReadFile(manifestPath))
	util.Must0(json.Unmarshal(b, &ManifestData))
}

func Sha256Version(f string) string {
	d, ok := ManifestData[f]
	if !ok {
		panic(fmt.Errorf("missing expected manifest entry for '%s'", f))
	}
	return d.Sha256Filename
}
