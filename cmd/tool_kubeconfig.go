// Copyright © 2016 Samsung CNCT
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
	"strings"

	"github.com/spf13/cobra"
)

// kubectlCmd represents the kubectl command
var kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Allows you to merge a kraken cluster admin.kubeconfig file to your local kubeconfig file.",
	Long: "Allows you to merge, remove, and backup a kraken cluster admin.kubeconfig file from your local kubeconfig file to be used with kubectl.",
	PreRunE: preRunGetClusterConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		command := []string{"/kraken/bin/computed_kubectl.sh", ClusterConfigPath}
		for _, element := range args {
			command = append(command, strings.Split(element, " ")...)
		}

		onFailure := func(out []byte) {
			fmt.Printf("%s \n", out)
		}

		onSuccess := func(out []byte) {
			fmt.Printf("%s \n", out)
		}

		ExitCode, err = runKrakenLibCommandNoSpinner(command, ClusterConfigPath, onFailure, onSuccess)

		return err
	},
}

func init() {
	toolCmd.AddCommand(kubeconfigCmd)
}
