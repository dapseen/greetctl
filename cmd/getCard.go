/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"time"

	"github.com/dapseen/greetctl/models/cards"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// getCardCmd represents the getCard command
var getCardCmd = &cobra.Command{
	Use:     "card <resource_id>",
	Short:   "This command is to read contents from card ",
	Aliases: []string{"cards"},
	Long: ` For example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr
	greetctl get card eva`,
	Run: func(cmd *cobra.Command, args []string) {

		//Connection
		var (
			timeout        = 2 * time.Second
			requestTimeout = 10 * time.Second
		)

		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
			DialTimeout: timeout,
		})

		key := clientv3.NewKV(cli)

		if err != nil {
			println(err)
		}

		defer cli.Close()
		if len(args) > 0 {
			cards.FetchCard(args[0], key, ctx)
		} else {
			cards.FetchAllCards()
		}
	},
}

func init() {
	getCmd.AddCommand(getCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
