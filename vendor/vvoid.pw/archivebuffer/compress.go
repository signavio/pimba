package archivebuffer

import (
	"bytes"
	"compress/gzip"
	"io"
)

func NewGzipBuffer(source io.Reader) (*bytes.Buffer, error) {
	gzipBuf := &bytes.Buffer{}
	archiver := gzip.NewWriter(gzipBuf)
	defer archiver.Close()

	_, err := io.Copy(archiver, source)
	return gzipBuf, err
}

func UngzipToBuffer(source io.Reader) (*bytes.Buffer, error) {
	archive, err := gzip.NewReader(source)
	if err != nil {
		return nil, err
	}
	defer archive.Close()

	writer := &bytes.Buffer{}
	_, err = io.Copy(writer, archive)
	return writer, err
}
