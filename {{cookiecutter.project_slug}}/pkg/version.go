package pkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("MyConsoleApp v%s\n", version)
	},
}
