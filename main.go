package main

import (
	"checksum/myhash"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()

	src := flag.String("src", "", "Path to check.")
	dst := flag.String("dst", "", "Path to compare with.")
	way := flag.String("way", "", "The way to check.\nPlease choose between SHA1, SHA256, SHA512, MD5 or ALL.")
	sum := flag.String("sum", "", "Checksum should be.")

	flag.Parse()

	checker := &myhash.Checker{Path: *src}

	if *sum != "" {
		checker.CheckSum = strings.ToUpper(*sum)
	}

	switch *way {
	case "sha1", "SHA1":
		checker.Flag = myhash.CHECK_SUM_SHA1

	case "sha256", "SHA256":
		checker.Flag = myhash.CHECK_SUM_SHA256

	case "sha512", "SHA512":
		checker.Flag = myhash.CHECK_SUM_SHA512

	case "md5", "MD5":
		checker.Flag = myhash.CHECK_SUM_MD5

	case "all", "ALL":
		checker.Flag = myhash.CHECK_SUM_ALL

	default:
		fmt.Println("Please choose between SHA1, SHA256, SHA512, MD5 or ALL.")
		os.Exit(0)
	}

	err := checker.Sum()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(checker)

	var dstChecker *myhash.Checker

	if *dst != "" {
		dstChecker = &myhash.Checker{Path: *dst, Flag: checker.Flag}
		err = dstChecker.Sum()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *sum != "" || *dst != "" {
		checker.IsCorrect = myhash.CompareCheckSum(dstChecker, checker, checker.CheckSum)
		if checker.IsCorrect {
			fmt.Print("\nCompare result: Right!!!\n\n")
		} else {
			fmt.Print("\nCompare result: False!!!\n\n")
		}
	}

	fmt.Println("Time used:", time.Since(timeStart))
	os.Exit(0)
}
