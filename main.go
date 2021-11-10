package main

import (
	"checksum/myhash"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	timeStart := time.Now()
	path := flag.String("p", "", "Path to check.")
	way := flag.String("w", "", "The way to check.\nPlease choose between SHA1, SHA256, SHA512, MD5 or ALL.")
	sum := flag.String("s", "", "Checksum should be.")
	flag.Parse()
	var cFlag bool = true
	var isRight bool
	if *sum == "" {
		cFlag = false
	}
	checker := myhash.Checker{Path: *path, Sum: strings.ToUpper(*sum)}
	switch *way {
	case "SHA1":
		isRight = myhash.Sha1(&checker, cFlag)
	case "sha1":
		isRight = myhash.Sha1(&checker, cFlag)
	case "SHA256":
		isRight = myhash.Sha256(&checker, cFlag)
	case "sha256":
		isRight = myhash.Sha256(&checker, cFlag)
	case "SHA512":
		isRight = myhash.Sha512(&checker, cFlag)
	case "sha512":
		isRight = myhash.Sha512(&checker, cFlag)
	case "MD5":
		isRight = myhash.Md5(&checker, cFlag)
	case "md5":
		isRight = myhash.Md5(&checker, cFlag)
	case "ALL":
		var wg sync.WaitGroup
		mychan := make(chan string)
		wg.Add(4)
		go myhash.MSha256(&checker, &wg, &mychan)
		go myhash.MSha512(&checker, &wg, &mychan)
		go myhash.MMd5(&checker, &wg, &mychan)
		go myhash.MSha1(&checker, &wg, &mychan)
		for i := 0; i < 4; i++ {
			fmt.Print(<-mychan)
		}
		wg.Wait()
		close(mychan)
		fmt.Println("Time:", time.Since(timeStart))
		os.Exit(0)
	case "all":
		var wg sync.WaitGroup
		mychan := make(chan string)
		wg.Add(4)
		go myhash.MSha256(&checker, &wg, &mychan)
		go myhash.MSha512(&checker, &wg, &mychan)
		go myhash.MMd5(&checker, &wg, &mychan)
		go myhash.MSha1(&checker, &wg, &mychan)
		for i := 0; i < 4; i++ {
			fmt.Print(<-mychan)
		}
		wg.Wait()
		close(mychan)
		fmt.Println("Time:", time.Since(timeStart))
		os.Exit(0)
	default:
		fmt.Println("Please choose between SHA1, SHA256, SHA512, MD5 or ALL.")
		os.Exit(0)
	}
	if cFlag {
		if isRight {
			fmt.Println("Right!!!")
			fmt.Println("Time:", time.Since(timeStart))
			os.Exit(0)
		} else {
			fmt.Println("False!!!")
			fmt.Println("Time:", time.Since(timeStart))
			os.Exit(0)
		}
	}
	fmt.Println("Time:", time.Since(timeStart))
	os.Exit(0)
}
