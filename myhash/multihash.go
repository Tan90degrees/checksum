package myhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func MSha1(checker *Checker, wg *sync.WaitGroup, mychan *chan string) {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer fp.Close()
	hasher := sha1.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	*mychan <- fmt.Sprintf("Size : %d bytes\nSHA1 : %s\nSHA1 : %s\n\n", size, sumUp, sumLow)
	wg.Done()
}

func MSha256(checker *Checker, wg *sync.WaitGroup, mychan *chan string) {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer fp.Close()
	hasher := sha256.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	*mychan <- fmt.Sprintf("Size : %d bytes\nSHA256 : %s\nSHA256 : %s\n\n", size, sumUp, sumLow)
	wg.Done()
}

func MSha512(checker *Checker, wg *sync.WaitGroup, mychan *chan string) {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer fp.Close()
	hasher := sha512.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	*mychan <- fmt.Sprintf("Size : %d bytes\nSHA512 : %s\nSHA512 : %s\n\n", size, sumUp, sumLow)
	wg.Done()
}

func MMd5(checker *Checker, wg *sync.WaitGroup, mychan *chan string) {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer fp.Close()
	hasher := md5.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	*mychan <- fmt.Sprintf("Size : %d bytes\nMD5 : %s\nmd5 : %s\n\n", size, sumUp, sumLow)
	wg.Done()
}
