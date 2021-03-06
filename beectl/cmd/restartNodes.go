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
	"github.com/marvin9002/swarm-install/beectl/nodes"
	"github.com/spf13/cobra"
)

// restartNodesCmd represents the restartNodes command
var restartNodesCmd = &cobra.Command{
	Use:   "restart-nodes",
	Short: "重启bee节点",
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}
		bee := nodes.NewSwarm()
		return bee.Restart(port)
	},
}

func init() {
	rootCmd.AddCommand(restartNodesCmd)

	restartNodesCmd.Flags().StringP("port", "", "", "[可选] 对应debug-api-addr 的端口,如果不输入重启全部")
	//addNodesCmd.Flags().StringP("show", "", "", "显示节点 debug-api-addr")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restartNodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restartNodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
