_My personal and very naive take on git_

Some terminalogies that helped me understand the problem clearly

Commit = Snapshot node (metadata + pointer to tree)
Tree = Directory structure at that snapshot
FileLog = Raw content of files at that moment
Branch = Pointer to the latest commit
History = Commit graph via parents (not stored in blob or index)
Index = Staged candidate snapshot (flat, not graph. A way to build said tree)

More importantly an easy way to understand the structure is:

- FileLog: what
- TreeEntry: where
- Commit: when + why
- Index: what’s next
