package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var l16 bool

var passgen = &cobra.Command{
	Use:   "passgen",
	Short: "Generate a random password",
	Long:  `Passgen is a tool that generates a random password`,
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword(l16)
		fmt.Println(password)
	},
}

func generatePassword(l16 bool) string {
	var length int

	if l16 {
		length = 16
	} else {
		length = rand.Intn(7) + 10
	}

	splchars := "!@#$_"
	upchars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowchars := "abcdefghijklmnopqrstuvwxyz"

	numSplChars := rand.Intn(3) + 2
	numUpChars := rand.Intn(3) + 3
	numLowChars := length - numSplChars - numUpChars

	var passwordBuilder strings.Builder

	for i := 0; i < numSplChars; i++ {
		passwordBuilder.WriteByte(splchars[rand.Intn(len(splchars))])
	}

	for i := 0; i < numUpChars; i++ {
		passwordBuilder.WriteByte(upchars[rand.Intn(len(upchars))])
	}

	for i := 0; i < numLowChars; i++ {
		passwordBuilder.WriteByte(lowchars[rand.Intn(len(lowchars))])
	}

	password := passwordBuilder.String()
	shuffledPassword := []rune(password)
	rand.Shuffle(len(shuffledPassword), func(i, j int) {
		shuffledPassword[i], shuffledPassword[j] = shuffledPassword[j], shuffledPassword[i]
	})
	return string(shuffledPassword)
}

func Execute() {
	passgen.Flags().BoolVarP(&l16, "l16", "l", false, "Generate a 16-character password (default is 10)")
	if err := passgen.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
