# archivebuffer

Utilities to help creating tarballs, untar, gzip and ungzip.

## NewTarballBuffer()

Creates a tarball from a source file or directory and returns a `*bytes.Buffer`.

### Usage

```
tarBuf, err := NewTarballBuffer("/tmp/foobar")
if != err {
    // handle error
}
```

## UntarToFile()

Receives a tar wrapped in an `io.Reader` and unarchives it to a determined path.

### Usage

```
err := UntarToFile(tarBuf, "/tmp")
if != err {
    // handle error
}
```

## NewGzipBuffer()

Creates a new Gzip from an `io.Reader` and returns a `*bytes.Buffer`. 

### Usage

```
gzipBuf, err := NewGzipBuffer(tarBuf)
if != err {
    // handle error
}
```

## UngzipToBuffer()

Receives a Gzip wrapped in an `io.Reader` and returns a `*bytes.Buffer`.

### Usage

```
ungzipBuf, err := UngzipToBuffer(gzipBuf)
if != err {
    // handle error
}
```

## Maintainers

Stephano Zanzin <sz@shitty.pizza>

## License

Please, refer to the [LICENSE](LICENSE) file.
