package wasnfilesystem

import (
	_ "embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"github.com/louisbuchbinder/core/lib/util"
	"github.com/louisbuchbinder/core/wasm"
	"github.com/louisbuchbinder/core/wasm/js"
)

//go:embed filesystem.html
var FILESYSTEM_TEMPLATE []byte

var filesystemTemplate *template.Template

type filesystemTemplateInput struct {
	Root       string
	Dir        string
	DirParts   []string
	Target     string
	ReadOnly   bool
	EntryInput *filesystemTemplateEntryInput
}

type filesystemTemplateEntryInput struct {
	Name     string                          `json:"name,omitempty"`
	Type     string                          `json:"type,omitempty"`
	Size     string                          `json:"size,omitempty"`
	Modified string                          `json:"modified,omitempty"`
	Children []*filesystemTemplateEntryInput `json:"children,omitempty"`
}

func mkdirAll(root, dirpath string) error {
	tree, err := js.NewOpfsFS(root)
	if err != nil {
		return err
	}
	return tree.MkdirAll(dirpath)
}

func removeAll(root, dirpath string) error {
	tree, err := js.NewOpfsFS(root)
	if err != nil {
		return err
	}
	return tree.RemoveAll(dirpath)
}

func newFilesystemTemplateEntryInput(root, dir string) (*filesystemTemplateEntryInput, error) {
	tree, err := js.NewOpfsFS(root)
	if err != nil {
		return nil, err
	}
	workingDir, err := tree.Open(dir)
	if err != nil {
		return nil, err
	}
	stat, err := workingDir.Stat()
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}
	dirEntries, err := workingDir.(fs.ReadDirFile).ReadDir(-1)
	if err != nil {
		return nil, err
	}
	children, err := util.MapOrError(dirEntries, func(_ int, e fs.DirEntry) (*filesystemTemplateEntryInput, error) {
		if e.IsDir() {
			return &filesystemTemplateEntryInput{
				Name:     e.Name(),
				Type:     "folder",
				Size:     "-",
				Modified: "-",
			}, nil
		}
		stat, err := e.Info()
		if err != nil {
			return nil, err
		}
		return &filesystemTemplateEntryInput{
			Name:     e.Name(),
			Type:     "file",
			Size:     fmt.Sprint(stat.Size()),
			Modified: stat.ModTime().Format("2006-01-02"),
		}, nil
	})

	input := &filesystemTemplateEntryInput{
		Name:     stat.Name(),
		Type:     "folder",
		Size:     "-",
		Modified: "-",
		Children: children,
	}
	return input, nil
}

var funcMap = template.FuncMap{
	"minus": minus,
}

func minus(a, b int) int {
	return a - b
}

func init() {
	filesystemTemplate = template.Must(template.New("FILESYSTEM_TEMPLATE").Funcs(funcMap).Parse(string(FILESYSTEM_TEMPLATE)))
}

func AsyncHandler(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	resolve, reject := args[0], args[1]
	_ = reject
	if len(args) < 3 {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   "missing expected req argument",
			Status: http.StatusBadRequest,
		}))
	}
	req, err := args[2].Request()
	if err != nil {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}))
	}

	target := util.OrVal(req.Url().GetSearchParam("target"), "")
	if target == "" {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   "unexpected missing target in query",
			Status: http.StatusBadRequest,
		}))
	}

	root := util.OrVal(req.Url().GetSearchParam("root"), "/")
	dir := util.OrVal(req.Url().GetSearchParam("dir"), "")

	entryInput, err := newFilesystemTemplateEntryInput(root, dir)
	if err != nil {
		fmt.Println(err)
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}))
	}

	input := &filesystemTemplateInput{
		Root:       root,
		Dir:        dir,
		DirParts:   util.Filter(strings.Split(dir, "/"), func(_ int, d string) bool { return d != "" }),
		Target:     target,
		ReadOnly:   req.Url().GetSearchParam("read-only") != nil,
		EntryInput: entryInput,
	}

	content, err := util.ExecuteTemplate(filesystemTemplate, input)
	if err != nil {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusInternalServerError,
		}))
	}
	return resolve.Invoke(js.NewJsResponse(&js.Response{
		Body:   string(content),
		Status: http.StatusOK,
	}))
}

func AsyncMkdirAllHandler(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	resolve, reject := args[0], args[1]
	_ = reject
	if len(args) < 3 {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   "missing expected req argument",
			Status: http.StatusBadRequest,
		}))
	}
	req, err := args[2].Request()
	if err != nil {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}))
	}

	root := util.OrVal(req.Url().GetSearchParam("root"), "/")
	dir := util.OrVal(req.Url().GetSearchParam("dir"), "")
	folder := util.OrVal(req.Url().GetSearchParam("folder"), "")

	if err := mkdirAll(root, strings.Join([]string{dir, folder}, "/")); err != nil {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}))
	}

	return AsyncHandler(args)
}

func AsyncRemoveAllHandler(args []wasm.Value) any {
	if len(args) < 2 {
		return nil
	}
	resolve, reject := args[0], args[1]
	_ = reject
	if len(args) < 3 {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   "missing expected req argument",
			Status: http.StatusBadRequest,
		}))
	}
	req, err := args[2].Request()
	if err != nil {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}))
	}

	root := util.OrVal(req.Url().GetSearchParam("root"), "/")
	dir := util.OrVal(req.Url().GetSearchParam("dir"), "")
	path := util.OrVal(req.Url().GetSearchParam("path"), "")

	if path == "" {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   "unexpected missing path in query",
			Status: http.StatusBadRequest,
		}))
	}

	parts := util.Filter(strings.Split(path, "/"), func(_ int, s string) bool { return s != "" })

	if parts[0] == "" {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   "unexpected missing path-part in query",
			Status: http.StatusBadRequest,
		}))
	}

	filepath := parts[0]
	if dir != "" {
		filepath = strings.Join([]string{dir, parts[0]}, "/")
	}

	if err := removeAll(root, filepath); err != nil {
		return resolve.Invoke(js.NewJsResponse(&js.Response{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}))
	}

	return AsyncHandler(args)
}
