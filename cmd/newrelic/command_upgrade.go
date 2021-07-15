package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Masterminds/semver/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/newrelic/newrelic-cli/internal/output"
	"github.com/newrelic/newrelic-cli/internal/utils"
)

const newRelicCLILatestTagURL = "https://api.github.com/repos/newrelic/newrelic-cli/releases/latest"
const prereleaseEnvironmentMessageFormat = `
  It appears you're in a development environment using prerelease version %s.
  To upgrade the New Relic CLI, you must be using non-prerelease version.
`

var (
	forceCLIUpgrade         bool
	targetCLIUpgradeVersion string
)

var cmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade to a newer version of the New Relic CLI",
	Long: `Use the upgrade command to upgrade the New Relic CLI to a newer version.
Default behavior upgrades to the latest version. Use the --version flag to upgrade
to a specific version.
`,
	Example: "newrelic version",
	Run:     upgradeCLI,
}

func upgradeCLI(cmd *cobra.Command, args []string) {
	currentVersion, err := semver.StrictNewVersion(version)
	if err != nil {
		// Log it, but do nothing (this logic should live in the upgrade service)
		log.Printf("error parsing current CLI version %s", version)
		return
	}

	isDevEnvironment := currentVersion.Prerelease() != ""
	if isDevEnvironment && !forceCLIUpgrade {
		log.WithFields(log.Fields{
			"version": version,
		}).Debug("cannot upgrade New Relic CLI version while in a development environment")

		output.Printf(prereleaseEnvironmentMessageFormat, version)
		return
	}

	ctx := context.Background()
	latestTag, err := fetchLatestTag(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		output.Print(latestTag)
	}
}

type GitHubRepositoryTag struct {
	TagName string `json:"tag_name"`
}

func fetchLatestTag(ctx context.Context) (*GitHubRepositoryTag, error) {
	client := utils.NewHTTPClient()

	respBytes, err := client.Get(ctx, newRelicCLILatestTagURL)

	log.Print("\n\n **************************** \n")
	log.Printf("\n fetchLatestTag - err:  %+v \n", err)

	if err != nil {
		return nil, err
	}

	repoTag := GitHubRepositoryTag{}

	err = json.Unmarshal(respBytes, &repoTag)
	if err != nil {
		return nil, err
	}

	log.Printf("\n fetchLatestTag - tag:  %+v \n", repoTag)
	log.Printf("\n fetchLatestTag - version:  %+v \n", version)
	log.Print("\n **************************** \n\n")
	time.Sleep(1 * time.Second)

	return &repoTag, nil
}

func init() {
	Command.AddCommand(cmdUpgrade)

	cmdUpgrade.Flags().BoolVarP(&forceCLIUpgrade, "force", "f", false, "Force upgrade the New Relic CLI (use with caution).")
	cmdUpgrade.Flags().StringVarP(&targetCLIUpgradeVersion, "version", "v", "", "The version to upgrade to. Must be newer than your current New Relic CLI version.")
}
