package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

var passgen = &cobra.Command{
	Use:   "passgen",
	Short: "Generate a random password",
	Long:  `Passgen is a tool that generates a random password`,
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword()
		fmt.Println(password)
	},
}

func generatePassword() string {
	length := rand.Intn(7) + 10
	chars := "!@#$_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if length < 10 && length > 16 {
		fmt.Println("Invalid password length. Please specify either 10 or 16.")
		os.Exit(1)
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

func Execute() {
	if err := passgen.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
