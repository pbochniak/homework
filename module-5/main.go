package main

import (
	"fmt"

	"github.com/pbochniak/module-5/blob"
	"github.com/pbochniak/module-5/blob/fs"
	"github.com/pbochniak/module-5/blob/memory"
)

func testMemoryBlobStorage() {
	var storage blob.BlobStorage = memory.NewBlobStorage()
	fmt.Println(storage.Update("a", []byte("abc")))
	fmt.Println(storage.Create("a", []byte("abc")))
	fmt.Println(storage.Create("a", []byte("xyz")))
	fmt.Println(storage.Read("a"))
	fmt.Println(storage.Update("a", []byte("xyz")))
	fmt.Println(storage.Delete("a"))
	fmt.Println(storage.Delete("a"))
	fmt.Println(storage.Read("a"))
}

func testFsBlobStorage() {
	var storage1 blob.BlobStorage = fs.NewBlobStorage("temp")
	var storage2 blob.BlobStorage = fs.NewBlobStorage("temp")
	fmt.Println(storage1.Update("a", []byte("abc")))
	fmt.Println(storage1.Create("a", []byte("abc")))
	fmt.Println(storage2.Create("a", []byte("xyz")))
	fmt.Println(storage2.Read("a"))
	fmt.Println(storage2.Update("a", []byte("xyz")))
	fmt.Println(storage1.Read("a"))
	fmt.Println(storage2.Delete("a"))
	fmt.Println(storage1.Read("a"))
}

func main() {
	testMemoryBlobStorage()
	fmt.Println()
	testFsBlobStorage()
}
