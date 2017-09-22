package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

const (
	Sep   = "="
	Usage = `
    usage: %s <entry>

      entry: Must match {NAME_OF_THE_ENTRY}={value}
      example:
        DEVIANT__EVENT_STORE__USERNAME=admin
`
)

var (
	ErrNotEnoughParts = fmt.Errorf("Entry should contain %s", Sep)
	ErrTooManyParts   = fmt.Errorf("Entry should contain only one %s", Sep)
	red               = color.New(color.FgRed).SprintFunc()
	green             = color.New(color.FgGreen).SprintFunc()
)

func Hash(value string) string {
	sum := md5.Sum([]byte(value))
	return hex.EncodeToString(sum[:])
}

func Format(entry string) (string, error) {
	parts := strings.Split(entry, Sep)
	if len(parts) < 2 {
		return "", ErrNotEnoughParts
	}
	if len(parts) > 2 {
		return "", ErrTooManyParts
	}
	name := parts[0]
	value := parts[1]
	return fmt.Sprintf("\"%s:%s\": \"%s\"", name, Hash(value), value), nil
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf(Usage, args[0])
		os.Exit(1)
	}

	entry, err := Format(args[1])
	if err != nil {
		fmt.Println(red(err.Error()))
		fmt.Printf(Usage, args[0])
		os.Exit(1)
	}
	fmt.Println(green(entry))
}
