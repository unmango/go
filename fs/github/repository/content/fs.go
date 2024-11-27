package content

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/google/go-github/v67/github"
	"github.com/spf13/afero"
)

type Fs struct {
	afero.ReadOnlyFs
	client *github.Client
	owner  string
	repo   string
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("https://github.com/%s/%s", f.owner, f.repo)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	file, dir, err := f.content(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}

	if file != nil {
		return &File{
			client:  f.client,
			content: file,
		}, nil
	} else {
		return &Directory{
			client:   f.client,
			name:     name,
			contents: dir,
		}, nil
	}
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	file, dir, err := f.content(name)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", name, err)
	}

	if file != nil {
		return &File{
			client:  f.client,
			content: file,
		}, nil
	} else {
		return &Directory{
			client:   f.client,
			name:     name,
			contents: dir,
		}, nil
	}
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	file, dir, err := f.content(name)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", name, err)
	}

	if file != nil {
		return &FileInfo{content: file}, nil
	} else {
		return &DirectoryInfo{
			name:     name,
			contents: dir,
		}, nil
	}
}

func (f *Fs) content(name string) (*github.RepositoryContent, []*github.RepositoryContent, error) {
	file, dir, _, err := f.client.Repositories.GetContents(context.TODO(), f.owner, f.repo, name, nil)
	if err != nil {
		return nil, nil, err
	}

	return file, dir, nil
}

func NewFs(gh *github.Client) afero.Fs {
	return &Fs{client: gh}
}
