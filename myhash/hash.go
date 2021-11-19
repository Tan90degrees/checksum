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
)

type Checker struct {
	Path string
	Sum  string
}

func Sha1(checker *Checker, cFlag bool) bool {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := sha1.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Printf("Size : %d bytes\nSHA1 : %s\nsha1 : %s\n\n", size, sumUp, sumLow)
	if cFlag {
		if sumUp == checker.Sum {
			return true
		} else {
			return false
		}
	}
	return true
}

func Sha256(checker *Checker, cFlag bool) bool {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := sha256.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Printf("Size : %d bytes\nSHA256 : %s\nsha256 : %s\n\n", size, sumUp, sumLow)
	if cFlag {
		if sumUp == checker.Sum {
			return true
		} else {
			return false
		}
	}
	return true
}

func Sha512(checker *Checker, cFlag bool) bool {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := sha512.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Printf("Size : %d bytes\nSHA512 : %s\nsha512 : %s\n\n", size, sumUp, sumLow)
	if cFlag {
		if sumUp == checker.Sum {
			return true
		} else {
			return false
		}
	}
	return true
}

func Md5(checker *Checker, cFlag bool) bool {
	fp, err := os.Open(checker.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := md5.New()
	size, err := io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}
	sumUp := fmt.Sprintf("%X", hasher.Sum(nil))
	sumLow := fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Printf("Size : %d bytes\nMD5 : %s\nmd5 : %s\n\n", size, sumUp, sumLow)
	if cFlag {
		if sumUp == checker.Sum {
			return true
		} else {
			return false
		}
	}
	return true
}
