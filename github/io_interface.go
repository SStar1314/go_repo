package io

type Reader interface {
    Read(p []byte)  (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadWriter interface {
    Reader
    Writer
}

type ReadWriterCloser interface {
    Reader
    Writer
    Closer
}
