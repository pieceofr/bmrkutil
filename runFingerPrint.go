package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
	"golang.org/x/crypto/sha3"
)

func runFingerPrint(c *cli.Context) error {
	filepath := c.String("file")
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	digest := sha3.Sum512(file)
	fingerprint := prefix + hex.EncodeToString(digest[:])
	fmt.Println("file:", filepath)
	fmt.Println("fingerprint:", fingerprint)
	return nil
}
