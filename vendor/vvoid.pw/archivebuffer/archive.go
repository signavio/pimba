package archivebuffer

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 'source' is the path to the file or directory that you
// would like to archive.
//
// Set 'keepBaseDir' to true or false, in order to define whether the
// base directory should be kept or not when being extracted.
//
// Example:
//     NewTarballBuffer("/tmp/foobar", true)
func NewTarballBuffer(source string, keepBaseDir bool) (*bytes.Buffer, error) {
	tarBuf := &bytes.Buffer{}
	tarball := tar.NewWriter(tarBuf)
	defer tarball.Close()

	sourceInfo, err := os.Stat(source)
	if err != nil {
		return nil, err
	}

	var baseDir string
	if sourceInfo.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		// TODO: Review this feature later and research for a better solution.
		if !keepBaseDir {
			header.Name = strings.TrimPrefix(header.Name, baseDir)
			header.Name = strings.TrimPrefix(header.Name, "/")
		}

		if err := tarball.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tarball, file)
		return err
	})

	return tarBuf, err
}

func UntarToFile(tarball io.Reader, target string) error {
	tarReader := tar.NewReader(tarball)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}
