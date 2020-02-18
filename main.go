package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		usage()
	}
	var secret string
	if strings.HasPrefix(os.Args[1], "otpauth://") {
		key, err := otp.NewKeyFromURL(os.Args[1])
		if err == nil {
			secret = key.Secret()
		} else {
			usage()
		}
	} else {
		secret = os.Args[1]
	}
	difficulty := 100000
	if len(os.Args) > 2 {
		diff, err := strconv.Atoi(os.Args[2])
		if err != nil {
			usage()
		}
		if diff < 1 || diff > 6 {
			usage()
		}
		difficulty = int(math.Pow10(6 - diff))
	}
	t := time.Now().UTC()
	step := 30 * time.Second
	for {
		code, _ := totp.GenerateCode(secret, t)
		numCode, _ := strconv.Atoi(code)
		if numCode < difficulty {
			fmt.Printf("code will be %s at %s\n", code, t)
			os.Exit(0)
		}
		t = t.Add(step)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <OTP secret|OTP URL> [difficulty (1-6)]\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}
