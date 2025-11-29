package internal

import (
	"errors"
	"os"
	"path/filepath"
)

const GitterDirName = ".gitter"

type Repo struct {
	WorkTree  string
	GitterDir string
}

// InitRepo initializes a new gitter repository at the given path
// It creates the necessary directories and files for a new repository to store its objects and references
func InitRepo(path string) (*Repo, error) {
	gdir := filepath.Join(path, GitterDirName)

	if _, err := os.Stat(gdir); err == nil {
		return nil, errors.New("gitter repository already exists")
	}

	objects := filepath.Join(gdir, "objects")
	refsHeads := filepath.Join(gdir, "refs", "heads")

	if err := os.MkdirAll(objects, 0755); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(refsHeads, 0755); err != nil {
		return nil, err
	}

	headPath := filepath.Join(gdir, "HEAD")
	if err := os.WriteFile(headPath, []byte("ref: refs/heads/main\n"), 0644); err != nil {
		return nil, err
	}

	mainRef := filepath.Join(refsHeads, "main")
	if err := os.WriteFile(mainRef, []byte(""), 0644); err != nil {
		return nil, err
	}

	indexPath := filepath.Join(gdir, "index")
	if err := os.WriteFile(indexPath, []byte(""), 0644); err != nil {
		return nil, err
	}

	return &Repo{
		WorkTree:  path,
		GitterDir: gdir,
	}, nil
}

func DiscoverRepo(start string) (*Repo, error) {
	path, err := filepath.Abs(start)
	if err != nil {
		return nil, err
	}

	for {
		gdir := filepath.Join(path, GitterDirName)
		if st, err := os.Stat(gdir); err == nil && st.IsDir() {
			return &Repo{
				WorkTree:  path,
				GitterDir: gdir,
			}, nil
		}

		parent := filepath.Dir(path)
		if parent == path {
			return nil, errors.New("no gitter repository found")
		}
		path = parent
	}
}
