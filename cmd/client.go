// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"media/api"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("client called")
		setLogLevel(cmd)
		ServerAddr, err := cmd.Flags().GetString("ServerAddr")
		if err != nil {
			log.Fatal(err)
		}

		if len(strings.TrimSpace(ServerAddr)) == 0 {
			log.Fatal("empty server address!")
		}

		typ, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}

		c := &api.RPCClient{ServerAddr: ServerAddr}
		// if err := c.Init(); err != nil {
		// 	log.Fatal("==> Init", err)
		// }

		switch typ {
		case "heartBeat":
			log.Infoln("Start HeartBeat test")
			if err = c.HeartBeat(); err != nil {
				log.Fatal("==> HeartBeat", err)
			}
		case "one":
			var ds []*api.HeartBeatRequest
			for i := 0; i < 20; i++ {
				ds = append(ds, &api.HeartBeatRequest{ClientID: "=====>>1", ClientIP: fmt.Sprintf("192.168.0.%d", i)})
			}
			if err := c.Stream(ds); err != nil {
				log.Fatal("==> Stream one || ", err)
			}
		case "two":
			var ds []*api.HeartBeatRequest
			for i := 50; i < 70; i++ {
				ds = append(ds, &api.HeartBeatRequest{ClientID: "=====>>2", ClientIP: fmt.Sprintf("172.18.0.%d", i)})
			}
			if err := c.Stream(ds); err != nil {
				log.Fatal("==> Stream two || ", err)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clientCmd.PersistentFlags().String("ServerAddr", "", "A help for foo")
	clientCmd.PersistentFlags().String("type", "", "A help for foo")
}
