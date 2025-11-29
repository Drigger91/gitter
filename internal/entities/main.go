package entities

import (
	"io/fs"
	"time"
)

// Commit is a snapshot of the repository at a given time
type Commit struct {
	ID       string
	Message  string
	Parents  []string
	Author   string
	Date     time.Time
	TreeID   string
}
// FileLog is a raw content of a file at a given time
type FileLog struct {
	ID      string
	Content []byte
}

type ObjectType int

const (
	ObjectFileLog ObjectType = iota
	ObjectTree
)
// TreeEntry is a single entry in a tree
type TreeEntry struct {
	Name     string
	Mode     fs.FileMode
	ObjectID string
	Type     ObjectType // FileLog or Tree, for clarity. Tree can be in case of a directory.
}

// Tree is a directory structure at a given commit
type Tree struct {
	ID      string
	Entries []TreeEntry
}

// IndexEntry is a single entry in the index
type IndexEntry struct {
	Path      string
	FileLogID string
	Mode      fs.FileMode
	Size      int64
	ModTime   time.Time
}

type Index struct {
	Entries map[string]IndexEntry
}
