package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"

	"flag"
	"log"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

type AocRoot struct {
	fs.Inode

	aocTree *Inode
}

func recursiveMap(r *fs.Inode, i *Inode, ctx context.Context) {
	for objName, obj := range i.Children {
		mode := uint32(0)
		if obj.IsDir {
			mode = fuse.S_IFDIR
		}
		ch := r.NewPersistentInode(ctx, obj, fs.StableAttr{Mode: mode})
		r.AddChild(objName, ch, false)
		if obj.IsDir {
			recursiveMap(ch, obj, ctx)
		}
	}
}

func (r *AocRoot) OnAdd(ctx context.Context) {
	recursiveMap(&r.Inode, r.aocTree, ctx)
}

func (r *AocRoot) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Mode = 0755
	return 0
}

var _ = (fs.NodeGetattrer)((*AocRoot)(nil))
var _ = (fs.NodeOnAdder)((*AocRoot)(nil))

func Atoi64OrPanic(input string) uint64 {
	entry, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return entry
}

var OK syscall.Errno

type Inode struct {
	fs.Inode
	Name     string
	Size     uint64
	IsDir    bool
	Children map[string]*Inode
	Parent   *Inode
}

var _ = (fs.NodeGetattrer)((*Inode)(nil))

func (f *Inode) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	mode := uint32(0644)
	if f.IsDir {
		mode = 0755
	}
	out.Attr = fuse.Attr{Mode: mode}
	out.Size = f.Size
	return OK
}

func ParseFsTree(input io.Reader) *Inode {
	scanner := bufio.NewScanner(input)

	root := Inode{
		Name:     "/",
		Children: make(map[string]*Inode),
		IsDir:    true,
	}

	currentNode := &root
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			// CD
			dirName := strings.TrimPrefix(line, "$ cd ")
			if dirName == "/" {
				currentNode = &root
			} else if dirName == ".." {
				currentNode = currentNode.Parent
			} else {
				currentNode = currentNode.Children[dirName]
			}
		} else if strings.HasPrefix(line, "dir ") {
			// dir
			dirName := strings.TrimPrefix(line, "dir ")
			newNode := Inode{
				Name:     dirName,
				IsDir:    true,
				Children: make(map[string]*Inode),
				Parent:   currentNode,
			}
			currentNode.Children[dirName] = &newNode
		} else if !strings.HasPrefix(line, "$ ls") {
			// file
			spaceIdx := strings.Index(line, " ")
			size := Atoi64OrPanic(line[0:spaceIdx])
			fileName := line[spaceIdx+1:]
			newNode := Inode{
				Name:   fileName,
				Size:   uint64(size),
				IsDir:  false,
				Parent: currentNode,
			}
			currentNode.Children[fileName] = &newNode
		}
	}
	return &root
}

func main() {
	debug := flag.Bool("debug", false, "print debug data")
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatal("Usage:\n  day7 input.txt MOUNTPOINT")
	}
	opts := &fs.Options{}
	opts.Debug = *debug
	inputFile, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("Input file fail: %v\n", err)
	}
	aocTree := ParseFsTree(inputFile)

	server, err := fs.Mount(flag.Arg(1), &AocRoot{aocTree: aocTree}, opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Wait()
}
