package entities

import (
	"io/fs"
	"time"
)

// Commit is a snapshot of the repository at a given time
type Commit struct {
	ID string
	Message string
	Parent []string // parent commits, can be multiple for merge commits
	Author string
	Date time.Time
	TreeID string
}

// FileLog this represents the state of a file, can be referenced by multiple tree entries
type FileLog struct {
	ID string
	Content []byte
}

// TreeEntry is a single entry in a tree
type TreeEntry struct {
	ID string
	Hash string
	FileLogID string
	Mode fs.FileMode
}

// Tree is a directory structure at a given commit
type Tree struct {
	ID string
	Entries []TreeEntry
}

// Index is a flat list of files that are going to be committed
type Index struct {
	ID string
	Entries []TreeEntry
}

// IndexEntry is a single entry in the index
// it represents a file that is going to be committed
type IndexEntry struct {
	Path string
	Date time.Time
	FileLogID string
}