package main

import (
	"context"
	"fmt"

	"github.com/newrelic/newrelic-cli/internal/cli"
	"github.com/spf13/cobra"
)

var (
	forceCLIUpgrade         bool
	targetCLIUpgradeVersion string
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Show the version of the New Relic CLI",
	Long: `Use the version command to print out the version of this command.
`,
	Example: "newrelic version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("newrelic version %s\n", version)
	},
}

var cmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade to a newer version of the New Relic CLI",
	Long: `Use the upgrade command to upgrade the New Relic CLI to a newer version.
Default behavior upgrades to the latest version. Use the --version flag to upgrade
to a specific version.
`,
	Example: "newrelic version",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli.UpgradeCLIVersion(ctx, version, targetCLIUpgradeVersion, forceCLIUpgrade)
	},
}

func init() {
	Command.AddCommand(cmdVersion)

	cmdVersion.AddCommand(cmdUpgrade)
	cmdUpgrade.Flags().BoolVarP(&forceCLIUpgrade, "force", "f", false, "Force upgrade the New Relic CLI (use with caution).")
	cmdUpgrade.Flags().StringVarP(&targetCLIUpgradeVersion, "version", "v", "", "The version to upgrade to. Must be newer than your current New Relic CLI version.")
}
