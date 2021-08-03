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
	"fmt"


	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card <name>",
	Short: "This command create awesome greetings",
	Long: `For example:
greetctl create card eva -n="Eva Green"
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")

		fmt.Println("Value of the flag name :" + name)
		fmt.Println("Value of the flag language :" + language)
		//fmt.Println("Here are the arguments: " + strings.Join(args, ","))
	},
}

func init() {
	createCmd.AddCommand(cardCmd) // making card a sub-command of create. eg, greetctl create card

	// Here you will define your flags and configuration settings.

	cardCmd.PersistentFlags().StringP("occaasion", "o", "", "Possible values: newyear, xmas, thanksgiving")

	cardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")

	cardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	cardCmd.MarkPersistentFlagRequired("name")
	cardCmd.MarkPersistentFlagRequired("occasion")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
