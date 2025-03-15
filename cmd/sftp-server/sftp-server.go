package main

import (
	"flag"
	"io"
	"os"

	"github.com/costinm/sftp"
)

// sftp is a standalone SFTP server, like /usr/lib/openssh/sftp-server but
// statically linked and using the go codebase.
//
// Stripped size with tinygo is 373k - which is very reasonable, openssh is
// smaller but depends on libc or musl.
//
// Only problem: tinygo doesn't support Chown yet.
func main() {
	s := &rwc{
		Reader:      os.Stdin,
		WriteCloser: os.Stdout,
	}
	ro := flag.Bool("R", false, "read-only")
	startDir := flag.String("d", "", "start dir")
	debug := flag.Bool("e", false, "print logs on stderr")
	baseDir := os.Getenv("SFTP_ROOT")
	opts := []sftp.ServerOption{}
	if *debug {
		opts = append(opts, sftp.WithDebug(os.Stderr))
	}
	if *startDir != "" {
		opts = append(opts, sftp.WithServerWorkingDirectory(*startDir))
	}
	if baseDir != "" {
		opts = append(opts, sftp.WithBasedir(baseDir))
	}
	if *ro {
		opts = append(opts, sftp.ReadOnly())
	}
	opts = append(opts, sftp.WithAllocator())

	srv, err := sftp.NewServer(s, opts...)
	if err != nil {
		panic(err)
	}

	err = srv.Serve()
	if err != nil {
		panic(err)
	}
	srv.Close()
}

type rwc struct {
	io.Reader
	io.WriteCloser
}
