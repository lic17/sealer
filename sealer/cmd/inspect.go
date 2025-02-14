// Copyright © 2021 Alibaba Group Holding Ltd.
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

	"github.com/alibaba/sealer/image"
	"github.com/spf13/cobra"
)

var clusterFilePrint bool

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "print the image information or clusterFile",
	Long: `sealer inspect kubernetes:v1.18.3 to print image information
sealer inspect -c kubernetes:v1.18.3 to print image Clusterfile`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if clusterFilePrint {
			cluster, err := image.GetClusterFileFromImageManifest(args[0])
			if err != nil {
				return fmt.Errorf("failed to find Clusterfile by image %s: %v", args[0], err)
			}
			fmt.Println(cluster)
		} else {
			file, err := image.GetYamlByImage(args[0])
			if err != nil {
				return fmt.Errorf("failed to find information by image %s: %v", args[0], err)
			}
			fmt.Println(file)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
	inspectCmd.Flags().BoolVarP(&clusterFilePrint, "Clusterfile", "c", false, "print the clusterFile")
}
