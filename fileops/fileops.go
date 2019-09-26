package fileops

import (
	"fmt"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/lang"
	"github.com/jfixby/pin/str"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const DIRECT_CHILDREN = true
const ALL_CHILDREN = !DIRECT_CHILDREN

var FoldersOnly = func(file string) bool {
	return IsFolder(file)
}

type FileFilter func(file string) bool

func ExtensionIs(file string, ext string) bool {
	return IsFile(file) && str.EndsWith(file, "."+ext)
}

func Abs(target string) string {
	if target == "" {
		dir, err := os.Getwd()
		lang.CheckErr(err)
		return dir
	}
	f, e := filepath.Abs(target)
	lang.CheckErr(e)
	return f
}

func ListFiles(target string, filter FileFilter, directChildren bool) []string {
	if IsFile(target) {
		lang.ReportErr("This is not a folder: %v", target)
	}

	files, err := ioutil.ReadDir(target)
	lang.CheckErr(err)
	result := []string{}
	for _, f := range files {
		fileName := f.Name()
		filePath := filepath.Join(target, fileName)

		if filter(filePath) {
			result = append(result, filePath)
		}

		if IsFolder(filePath) && directChildren == ALL_CHILDREN {
			children := ListFiles(filePath, filter, ALL_CHILDREN)
			result = append(result, children...)
			continue
		}
	}
	return result
}

func NameWithoutExtention(file string) string {
	p := PathToArray(file)
	fileName := p[len(p)-1]
	i := str.IndexOf(fileName, ".", 0)
	lang.AssertNot("IndexOf(.)", i, -1)
	name := fileName[:i]
	return name
}

func SplitPath(fullPath, prefixPath string) string {
	if strings.Index(fullPath, prefixPath) != 0 {
		lang.ReportErr("Incorrect path prefix: <%v>\n expected <%v>",
			prefixPath,
			fullPath,
		)
	}
	list := PathToArray(fullPath)
	prefixList := PathToArray(prefixPath)
	return filepath.Join(list[len(prefixList):]...)
}

func PathToArray(path string) []string {
	// The same implementation is used in LookPath in os/exec;
	// consider changing os/exec when changing this.
	sep := os.PathSeparator
	ListSeparator := uint8(sep)
	if path == "" {
		return []string{}
	}

	// Split path, respecting but preserving quotes.
	list := []string{}
	start := 0
	quo := false
	for i := 0; i < len(path); i++ {
		c := path[i]
		split := c == ListSeparator
		if c == '"' {
			quo = !quo
		} else if split && !quo {
			list = append(list, path[start:i])
			start = i + 1
		}
	}
	list = append(list, path[start:])

	// Remove quotes.
	for i, s := range list {
		list[i] = strings.Replace(s, `"`, ``, -1)
	}

	return list
}

func IsFolder(target string) bool {
	fileInfo, err := os.Stat(target)
	lang.CheckErr(err)
	return fileInfo.IsDir()
}

func IsFile(target string) bool {
	return !IsFolder(target) && FileExists(target)
}
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func Delete(filename string) {
	pin.D("delete", filename)
	if protectFromDeletion {
		return
	}
	err := os.RemoveAll(filename)
	lang.CheckErr(err)
}

func AppendStringToFile(path, text string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	lang.CheckErr(err)
	defer f.Close()
	_, err = f.WriteString(text)
	lang.CheckErr(err)
}

func ReadFileToString(file string) string {
	b, err := ioutil.ReadFile(file) // just pass the file name
	lang.CheckErr(err)
	str := string(b) // convert content to a 'string'
	str = strings.Replace(str, "\r", "", -1)
	return str
}

func EngageDeleteSafeLock(safeLock bool) {
	protectFromDeletion = safeLock
}

var protectFromDeletion = false

func WriteStringToFile(file, data string) {
	ensureParent(file)

	fo, err := os.Create(file)
	lang.CheckErr(err)
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(data))
	lang.CheckErr(err)
}

func Parent(file string) string {
	return filepath.Dir(file)
}

func ensureParent(file string) {
	parent := filepath.Dir(file)
	os.MkdirAll(parent, 0700)
}

func Copy(from string, to string) {
	pin.D("copy", from)
	pin.D("to", to)
	ensureParent(to)
	_, err := copy(from, to)
	lang.CheckErr(err)
}

func CopyFolderContentToFolder(from string, to string, filter FileFilter, directChildren bool) {
	pin.D("copy", from)
	pin.D("  to", to)
	inputFiles := ListFiles(from, filter, directChildren)
	//pin.D("inputFiles", inputFiles)
	for _, f := range inputFiles {
		postfix := strings.TrimPrefix(f, from)
		//pin.D("postfix", postfix)
		newpath := filepath.Join(to, postfix)
		//pin.D("newpath", newpath)
		if IsFolder(f) {
			err := os.MkdirAll(newpath, 0700)
			lang.CheckErr(err)
			pin.D("make", newpath)
			continue
		}
		if IsFile(f) {
			if IsFolder(f) {
				lang.ReportErr("This is not a file: %v", from)
			}
			Copy(f, newpath)
			continue
		}
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
