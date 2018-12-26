// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package ui

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// TemplateDir statically implements the virtual filesystem provided to vfsgen.
var TemplateDir = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 26, 21, 47, 23, 73468341, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2018, 12, 26, 21, 47, 23, 73468341, time.UTC),
			uncompressedSize: 816,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x51\xbd\x8e\xdb\x3c\x10\xec\xfd\x14\x7b\xaa\x3f\x91\xdf\x39\x08\x12\x18\x94\x9b\xe4\x80\x54\xb1\x81\x4b\x8a\x94\x2b\x72\x15\xd1\x21\x29\x81\xbb\x92\x7d\x6f\x1f\xc8\xb2\xce\x8e\xf3\xd3\x48\x9c\x9d\xc1\xec\x70\x68\x1e\x3e\xee\x3e\x7c\xf9\xb6\x7f\x82\x56\x62\xd8\xae\xcc\xfc\x03\x30\x2d\xa1\x9b\x0e\x00\x26\x92\x20\xd8\x16\x33\x93\x54\xc5\x20\x4d\xf9\xbe\xb8\xa5\x12\x46\xaa\x8a\xd1\xd3\xb1\xef\xb2\x14\x60\xbb\x24\x94\xa4\x2a\x8e\xde\x49\x5b\x39\x1a\xbd\xa5\xf2\x0c\xfe\x03\x9f\xbc\x78\x0c\x25\x5b\x0c\x54\x3d\x2e\x46\xe2\x25\xd0\x76\x8f\xd2\xc2\xae\x81\xa7\x93\x0f\x04\xcf\x82\x32\x30\xec\xf1\x3b\x19\x3d\x0b\x66\x71\xf0\xe9\x07\x64\x0a\x55\xc1\xf2\x12\x88\x5b\x22\x29\xa0\xcd\xd4\x54\x45\x2b\xd2\xf3\x46\x6b\xeb\xd2\x81\x95\x0d\xdd\xe0\x9a\x80\x99\x94\xed\xa2\xc6\x03\x9e\x74\xf0\x35\xeb\x7a\x08\x11\xf5\xff\xea\x9d\x5a\x6b\xcb\x17\xac\xa2\x4f\xca\x32\x2f\xa1\xd8\x66\xdf\x0b\x38\x6a\x28\x03\x67\x7b\x75\x1f\x98\x54\xd3\x25\xc1\x23\x71\x17\x67\xf3\x4c\x81\x90\x89\xf5\xf8\x56\xbd\x51\x8f\xfa\xc0\x1a\x43\x50\x07\x2e\xb6\x46\xcf\x56\xe7\x66\xf5\x52\xad\xa9\x3b\xf7\xb2\xac\x22\x2b\xbe\x4b\x60\x03\x32\x57\xc5\x05\x5e\x82\x00\x18\xe7\xc7\x85\x9b\xea\x45\x9f\x28\xbf\xb2\xf7\x7c\x18\x62\xba\x21\xa7\x7a\xb1\x0e\xb4\x08\x66\xe0\xb9\xac\xbb\xec\x28\x93\xfb\x45\x3b\xa9\xaf\xc1\x6e\xa7\xf9\x7e\x34\x0d\xdd\x56\x7d\xea\x58\x8c\x16\xf7\x17\xfa\x6b\xff\x0f\xf2\xf3\x10\x6b\xca\xbb\xe6\x19\x63\x1f\x88\xff\xa4\x34\xfa\x7e\xb3\xd1\xbf\x25\x34\xfa\x7c\xab\x9b\x46\xb4\xf3\xe3\x6b\x7d\x57\x60\xf4\xa5\xdb\xf9\x2d\x66\x9f\x95\x79\x28\x4b\x18\x7d\xdc\x80\x70\xb5\xde\xf0\xfc\x3d\x56\xeb\x0d\x9d\x7a\x4c\x4e\xb0\x5e\x95\xe5\x76\xf5\x33\x00\x00\xff\xff\x9e\x03\x4e\x2b\x30\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}