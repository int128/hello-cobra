package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmdEcho := &cobra.Command{
		Use:   "echo",
		Short: "Just say args",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("args=%+v", args)
		},
	}
	cmdEcho.Flags().String("bar", "", "bar (echo)")

	rootCmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "Hello world with cobra",
		Long: fmt.Sprintf(`Hello world with cobra.

Examples:
  # Root
  %[1]s

  # Echo
  %[1]s echo
`, os.Args[0]),
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("hello")
		},
	}
	rootCmd.Flags().String("foo", "", "foo (root)")

	rootCmd.AddCommand(cmdEcho)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
