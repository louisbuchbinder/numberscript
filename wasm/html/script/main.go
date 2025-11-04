package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
)

var (
	file         string
	module       string
	output       string
	outputSHA256 string
)

func init() {
	flag.StringVar(&file, "file", "", "[required] the input file for the html script src")
	flag.StringVar(&module, "module", "", "[required] module path for the script src")
	flag.StringVar(&output, "output", "", "[required] the output file path")
	flag.StringVar(&outputSHA256, "output-sha256", "", "[required] the output sha256 file path")
	flag.Parse()
}

func main() {
	f := util.Must(os.Open(file))
	defer f.Close()
	sha256 := util.Must(util.Sha256HexOfFile(f))
	filename := strings.TrimPrefix(path.Base(file), "sha256.")
	filenameSHA256 := sha256 + "." + filename
	src := fmt.Sprintf("/%s/pkg/%s", module, filename)
	srcSHA256 := fmt.Sprintf("/%s/pkg/%s", module, filenameSHA256)
	target := `<script src="%s"></script>`
	content := fmt.Sprintf(target, src)
	contentSHA256 := fmt.Sprintf(target, srcSHA256)
	util.Must0(os.WriteFile(output, []byte(content), 0o644))
	util.Must0(os.WriteFile(outputSHA256, []byte(contentSHA256), 0o644))
}
