package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func setLogLevel(cmd *cobra.Command) {
	isDebug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}

	if isDebug {
		log.Infoln("Start as debug mode")
		log.SetLevel(log.DebugLevel)
	}
}
