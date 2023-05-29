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

const (
	CHECK_SUM_SHA1 = iota
	CHECK_SUM_SHA256
	CHECK_SUM_SHA512
	CHECK_SUM_MD5
	CHECK_SUM_ALL
)

type Checker struct {
	Path      string
	CheckSum  string
	IsCorrect bool
	Flag      int
	Size      int64
	SumsUp    [CHECK_SUM_ALL]string
	SumsLow   [CHECK_SUM_ALL]string
}

func (c *Checker) Sum() error {
	var wg sync.WaitGroup
	fi, err := os.Stat(c.Path)
	if err != nil {
		return err
	}
	c.Size = fi.Size()
	switch c.Flag {
	case CHECK_SUM_SHA1:
		wg.Add(1)
		go c.Sha1(&wg)
	case CHECK_SUM_SHA256:
		wg.Add(1)
		go c.Sha256(&wg)
	case CHECK_SUM_SHA512:
		wg.Add(1)
		go c.Sha512(&wg)
	case CHECK_SUM_MD5:
		wg.Add(1)
		go c.Md5(&wg)
	case CHECK_SUM_ALL:
		wg.Add(CHECK_SUM_ALL)
		go c.Sha1(&wg)
		go c.Sha256(&wg)
		go c.Sha512(&wg)
		go c.Md5(&wg)
	default:
		return fmt.Errorf("bad flag : %d", c.Flag)
	}
	wg.Wait()
	return nil
}

func (c *Checker) Sha1(wg *sync.WaitGroup) {
	fp, err := os.Open(c.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := sha1.New()
	_, err = io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}

	c.SumsUp[CHECK_SUM_SHA1] = fmt.Sprintf("%X", hasher.Sum(nil))
	c.SumsLow[CHECK_SUM_SHA1] = fmt.Sprintf("%x", hasher.Sum(nil))

	wg.Done()
}

func (c *Checker) Sha256(wg *sync.WaitGroup) {
	fp, err := os.Open(c.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := sha256.New()
	_, err = io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}

	c.SumsUp[CHECK_SUM_SHA256] = fmt.Sprintf("%X", hasher.Sum(nil))
	c.SumsLow[CHECK_SUM_SHA256] = fmt.Sprintf("%x", hasher.Sum(nil))

	wg.Done()
}

func (c *Checker) Sha512(wg *sync.WaitGroup) {
	fp, err := os.Open(c.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := sha512.New()
	_, err = io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}

	c.SumsUp[CHECK_SUM_SHA512] = fmt.Sprintf("%X", hasher.Sum(nil))
	c.SumsLow[CHECK_SUM_SHA512] = fmt.Sprintf("%x", hasher.Sum(nil))

	wg.Done()
}

func (c *Checker) Md5(wg *sync.WaitGroup) {
	fp, err := os.Open(c.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	hasher := md5.New()
	_, err = io.Copy(hasher, fp)
	if err != nil {
		log.Fatal(err)
	}

	c.SumsUp[CHECK_SUM_MD5] = fmt.Sprintf("%X", hasher.Sum(nil))
	c.SumsLow[CHECK_SUM_MD5] = fmt.Sprintf("%x", hasher.Sum(nil))

	wg.Done()
}

func (c *Checker) String() string {
	var res string
	res = fmt.Sprintf("File(%s) checksum result:\n", c.Path)
	switch c.Flag {
	case CHECK_SUM_SHA1:
		res += fmt.Sprintf("  Size : %d bytes\n  SHA1 : %s\n  sha1 : %s", c.Size, c.SumsUp[CHECK_SUM_SHA1], c.SumsLow[CHECK_SUM_SHA1])
	case CHECK_SUM_SHA256:
		res += fmt.Sprintf("  Size : %d bytes\n  SHA256 : %s\n  sha256 : %s", c.Size, c.SumsUp[CHECK_SUM_SHA256], c.SumsLow[CHECK_SUM_SHA256])
	case CHECK_SUM_SHA512:
		res += fmt.Sprintf("  Size : %d bytes\n  SHA512 : %s\n  sha512 : %s", c.Size, c.SumsUp[CHECK_SUM_SHA512], c.SumsLow[CHECK_SUM_SHA512])
	case CHECK_SUM_MD5:
		res += fmt.Sprintf("  Size : %d bytes\n  MD5 : %s\n  md5 : %s", c.Size, c.SumsUp[CHECK_SUM_MD5], c.SumsLow[CHECK_SUM_MD5])
	case CHECK_SUM_ALL:
		res += fmt.Sprintf("  Size : %d bytes\n  SHA1 : %s\n  sha1 : %s\n", c.Size, c.SumsUp[CHECK_SUM_SHA1], c.SumsLow[CHECK_SUM_SHA1])
		res += fmt.Sprintf("  SHA256 : %s\n  sha256 : %s\n", c.SumsUp[CHECK_SUM_SHA256], c.SumsLow[CHECK_SUM_SHA256])
		res += fmt.Sprintf("  SHA512 : %s\n  sha512 : %s\n", c.SumsUp[CHECK_SUM_SHA512], c.SumsLow[CHECK_SUM_SHA512])
		res += fmt.Sprintf("  MD5 : %s\n  md5 : %s", c.SumsUp[CHECK_SUM_MD5], c.SumsLow[CHECK_SUM_MD5])
	default:
		res = "Empty"
	}

	return res
}

func CompareCheckSum(dst *Checker, src *Checker, sum string) bool {
	res := false

	if sum != "" {
		for i := 0; i < CHECK_SUM_ALL; i++ {
			if src.SumsUp[i] == sum {
				res = true
			}
		}
	} else {
		res = true
	}

	if !res || dst == nil {
		return res
	}

	if dst.Flag != CHECK_SUM_ALL {
		return dst.SumsUp[dst.Flag] == src.SumsUp[dst.Flag]
	} else {
		for i := 0; i < CHECK_SUM_ALL; i++ {
			if dst.SumsUp[i] != src.SumsUp[i] {
				fmt.Printf("Dst: name(%s), sum(%s), size(%d)\nSrc: name(%s), sum(%s), size(%d)\n", dst.Path, dst.SumsUp[i], dst.Size, src.Path, src.SumsUp[i], src.Size)
				return false
			}
		}
		return true
	}
}
