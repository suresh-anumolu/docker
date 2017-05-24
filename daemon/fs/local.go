package fs

import (
	"io"
	"os"
	"runtime"

	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/chrootarchive"
	"github.com/docker/docker/pkg/scopedpath"
)

type localfs struct {
	root string
}

func (lfs *localfs) Remote() bool {
	return false
}

func (lfs *localfs) HostPathName() string {
	return lfs.root
}

func (lfs *localfs) ExtractArchive(input io.Reader, path string, options *archive.TarOptions) error {
	return chrootarchive.Untar(input, path, options)
}

func (lfs *localfs) ArchivePath(path string, options *archive.TarOptions) (io.ReadCloser, error) {
	return archive.TarWithOptions(path, options)
}

func (lfs *localfs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (lfs *localfs) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (lfs *localfs) ResolveFullPath(name string) (string, error) {
	return scopedpath.EvalScopedPath(name, lfs.root)
}

func (*localfs) Platform() string {
	return runtime.GOOS
}
