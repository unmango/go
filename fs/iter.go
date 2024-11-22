package aferox

type iterOptions struct {
	skipDirs bool
}

type IterOption func(*iterOptions)
