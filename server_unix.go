//go:build !windows && !plan9
// +build !windows,!plan9

package sftp

import (
	"path"
	"path/filepath"
)

// toLocalPath converts a relative path to a local path,
// and add the 'baseDir' prefix.
//
// This works for open, ls and most operations.
// Doesn't work for symlink targets (yet).
// Poor man susbtitute for chroot - when running as regular
// user.
func (s *Server) toLocalPath(p string) string {
	p = filepath.Clean(p)
	if s.workDir != "" && !path.IsAbs(p) {
		p = path.Join(s.workDir, p)
	}
	if s.baseDir != "" {
		p = path.Join(s.baseDir, p)
	}
	return p
}

// toRealpath is used in the 'realpath' command - will add the workdir and
// clean the path, but not add the 'base'
func (s *Server) toRealpath(p string) string {
	p = filepath.Clean(p)
	if s.workDir != "" && !path.IsAbs(p) {
		p = path.Join(s.workDir, p)
	}
	return p
}
