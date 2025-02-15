package main

import (
	"context"

	"github.com/malikbenkirane/natsio/svctl/svctl/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(
		cmd.
			NewCLI().
			ExecuteContext(context.Background()),
	)
}
