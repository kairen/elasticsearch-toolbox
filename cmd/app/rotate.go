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
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/kairen/elasticsearch-toolbox/pkg/elastic"
	"github.com/kairen/elasticsearch-toolbox/pkg/util"
	"github.com/spf13/cobra"
)

type rotate struct {
	days          int
	regexPatterns []string
	dateFormat    string
}

func newRotateCmd() *cobra.Command {
	toolbox := &rotate{}
	rotateCmd := &cobra.Command{
		Use:   "rotate [flags]",
		Short: "Delete expired indices",
		Long:  "Automatically delete indices in Elasticsearch older than a given number of days.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(toolbox.regexPatterns) == 0 {
				cmd.Usage()
				os.Exit(1)
			}

			for _, regex := range toolbox.regexPatterns {
				err := util.Retry(time.Second*2, retryCount, func() error {
					return toolbox.deleteExpiredIndices(regex)
				})
				if err != nil {
					glog.Errorf("Failed to remove \"%s\" indices: %s.\n", regex, err)
				}
			}
			return nil
		},
	}

	rotateCmd.Flags().IntVarP(&toolbox.days, "days", "d", 90, "Days to keep.")
	rotateCmd.Flags().StringSliceVarP(&toolbox.regexPatterns, "index-regex-patterns", "", nil, "Index's regex pattern.")
	rotateCmd.Flags().StringVarP(&toolbox.dateFormat, "date-format", "", "2006.1.2", "Format template for parsing date.")
	return rotateCmd
}

func (d *rotate) deleteExpiredIndices(regex string) error {
	client, err := elastic.NewClient(*cfg)
	if err != nil {
		return err
	}

	indices, err := client.CatIndices(regex)
	if err != nil {
		return err
	}

	if len(indices) == 0 {
		glog.V(2).Infof("The indices are not matched with \"%s\" regex.", regex)
		return nil
	}

	for _, i := range indices {
		date := util.ParseDate(i.Index)
		expired, err := util.IsExpired(d.days, date, d.dateFormat)
		if err != nil {
			return err
		}

		if expired {
			glog.V(3).Infof("The index \"%s\" has expired.", i.Index)
			if _, err := client.DeleteIndex(i.Index); err != nil {
				return err
			}
			glog.Infof("Successfully removed index %s.", i.Index)
		}
	}
	return nil
}
