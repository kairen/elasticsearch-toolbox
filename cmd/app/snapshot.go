// Copyright © 2020 Kyle Bai <k2r2.bai@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/kairen/elasticsearch-toolbox/pkg/elastic"
	"github.com/kairen/elasticsearch-toolbox/pkg/util"
	"github.com/spf13/cobra"
)

type snapshot struct {
	repository    string
	prefix        string
	dateFormat    string
	regexPatterns []string
}

func newSnapshotCmd() *cobra.Command {
	snapshot := snapshot{}
	cmd := &cobra.Command{
		Use:   "snapshot [flags]",
		Short: "Create a snapshot for backuping indices",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(snapshot.repository) == 0 {
				cmd.Usage()
				os.Exit(1)
			}

			err := util.Retry(time.Second*5, retryCount, func() error {
				return snapshot.createSnapshot()
			})
			if err != nil {
				glog.Errorf("Failed to create \"%s\" snapshot: %s.\n", snapshot.repository, err)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&snapshot.repository, "repository", "", "", "Name of repository.")
	cmd.Flags().StringVarP(&snapshot.prefix, "prefix", "", "snapshot", "Prefix name for snapshot(Full name is '{prefix}_{date}')")
	cmd.Flags().StringVarP(&snapshot.dateFormat, "date-format", "", "2006.01.02-15.04.05", "Format template for parsing date")
	cmd.Flags().StringSliceVarP(&snapshot.regexPatterns, "index-regex-patterns", "", nil, "Index's regex patterns.")
	return cmd
}

func (d *snapshot) createSnapshot() error {
	client, err := elastic.NewClient(*cfg)
	if err != nil {
		return err
	}

	snapshot := fmt.Sprintf("%s_%s", d.prefix, time.Now().Format(d.dateFormat))
	_, err = client.CreateSnapshot(d.repository, snapshot, strings.Join(d.regexPatterns, ","))
	if err != nil {
		return err
	}
	glog.Infof("Successfully create a snapshot '%s' in %s.", snapshot, d.repository)
	return nil
}
