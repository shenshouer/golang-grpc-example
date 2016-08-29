package cmd

import (
	"fmt"

	"media/api"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("server called")
		setLogLevel(cmd)

		advertise, err := cmd.Flags().GetString("advertise")
		if err != nil {
			log.Fatal(err)
		}

		s := &api.RPCServer{Advertise: advertise}
		err = s.Serve()
		if err != nil{
			log.Fatal()
		}
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.PersistentFlags().StringP("advertise", "a", ":8001", "advertise for rest api server")

}
