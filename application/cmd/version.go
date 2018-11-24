package cmd

import (
	"fmt"
	"github.com/duck8823/duci/application"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = createCmd("version", "Display version", displayVersion)

func displayVersion(cmd *cobra.Command, _ []string) {
	readConfiguration(cmd)

	println(fmt.Sprintf("Version: %s", application.VersionString()))
	if application.IsLatestVersion() {
		os.Exit(0)
		return
	}

	println(fmt.Sprintf(
		"%s is not latest, you should upgrade to v%s",
		application.VersionStringShort(),
		application.CurrentVersion(),
	))
}
