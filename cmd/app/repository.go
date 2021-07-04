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
	"github.com/kairen/elasticsearch-toolbox/pkg/objconfig"
	"github.com/kairen/elasticsearch-toolbox/pkg/objconfig/s3"
	"github.com/kairen/elasticsearch-toolbox/pkg/util"
	"github.com/spf13/cobra"
)

type repository struct {
	provider    string
	bucket      string
	s3AccessKey string
	s3SecretKey string
	s3Region    string
	s3RoleARN   string
	compress    bool
}

func newRepositoryCmd() *cobra.Command {
	repo := repository{}
	cmd := &cobra.Command{
		Use:   "repository [flags]",
		Short: "Create a repository to place snapshots",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(repo.bucket) == 0 || len(repo.provider) == 0 {
				cmd.Usage()
				os.Exit(1)
			}

			err := util.Retry(time.Second*2, retryCount, func() error {
				return repo.createRepository()
			})
			if err != nil {
				glog.Errorf("Failed to create \"%s\" snapshot repository: %s.\n", repo.bucket, err)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&repo.bucket, "bucket", "", "", "Name of bucket for objstore")
	cmd.Flags().BoolVarP(&repo.compress, "compress", "", false, "Enable for compressing file")
	cmd.Flags().StringVarP(&repo.provider, "provider", "", "", "Object sotre provider('s3', 'gcs', 'azure')")
	cmd.Flags().StringVarP(&repo.s3AccessKey, "s3.access-key", "", "", "AWS access key for S3")
	cmd.Flags().StringVarP(&repo.s3SecretKey, "s3.secret-key", "", "", "AWS secret key for S3")
	cmd.Flags().StringVarP(&repo.s3Region, "s3.regoin", "", "", "AWS region")
	cmd.Flags().StringVarP(&repo.s3RoleARN, "s3.role-arn", "", "", "AWS IAM role arn for snapshotting")
	return cmd
}

func (d *repository) createRepository() error {
	var objcfg objconfig.Interface
	switch objconfig.Provider(d.provider) {
	case objconfig.S3:
		objcfg = &s3.Config{
			Bucket:               d.bucket,
			AccessKey:            d.s3AccessKey,
			SecretKey:            d.s3SecretKey,
			Region:               d.s3Region,
			RoleARN:              d.s3RoleARN,
			Compress:             d.compress,
			ServerSideEncryption: true,
		}
	case objconfig.GCS:
		// TODO: Implement google cloud storage
	case objconfig.Azure:
		// TODO: Implement azure object service
	default:
		glog.Fatalf("The provider '%s' is not supported.", d.provider)
	}

	client, err := elastic.NewClient(*cfg)
	if err != nil {
		return err
	}

	_, err = client.CreateSnapshotRepository(objcfg)
	if err != nil {
		return err
	}
	glog.Infof("Successfully create a snapshot repostory %s.", d.bucket)
	return nil
}
