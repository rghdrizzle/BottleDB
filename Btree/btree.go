package btree

type BTree struct{
	root uint64
	get func(uint64) []byte // reads a page from disk
	new func([]byte) // allocates and writes a enw page
	del func(uint64) // delets a page
}