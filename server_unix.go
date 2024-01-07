//go:build !windows && !plan9
// +build !windows,!plan9

package sftp

import (
	"path"
	"path/filepath"
)

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
