// Code generated by "esc -o std.go -pkg=jsonnet std/std.jsonnet"; DO NOT EDIT.

package jsonnet

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/std/std.jsonnet": {
		local:   "std/std.jsonnet",
		size:    42048,
		modtime: 1509761227,
		compressed: `
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
d/4vAAD//z/DaoBApAAA
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/std": {
		isDir: true,
		local: "std",
	},
}
