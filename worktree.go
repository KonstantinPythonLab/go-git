package git

import (
	"strings"
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func (w *Worktree) isIgnoreCase() bool {
	cfg, err := w.r.Config()
	if err != nil {
		return false
	}
	return cfg.Core.IgnoreCase
}

func (w *Worktree) pathsMatch(a, b string) bool {
	if w.isIgnoreCase() {
		return strings.EqualFold(a, b)
	}
	return a == b
}