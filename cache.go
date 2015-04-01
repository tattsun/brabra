package main

import (
	sha "crypto/sha256"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

type CacheId struct {
	cacheId string
}

func (self *CacheId) getId() string {
	return self.cacheId
}

func GetCacheWriter(basedir string, cacheid *CacheId) io.WriteCloser {
	file, err := os.Create(basedir + "/" + cacheid.getId() + ".png")
	if err != nil {
		panic(err)
	}
	return file
}

func GenCacheId(uri string, level float64) *CacheId {
	data := []byte(fmt.Sprintf("%s%f", uri, level))
	checkSum := sha.Sum256(data)
	return &CacheId{cacheId: fmt.Sprintf("%x", checkSum)}
}

func GetCachedImage(basedir string, cacheid *CacheId) (image.Image, error) {
	file, err := os.Open(basedir + "/" + cacheid.getId() + ".png")
	defer file.Close()
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	return img, nil
}
