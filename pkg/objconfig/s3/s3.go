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

package s3

import (
	"github.com/kairen/elasticsearch-toolbox/pkg/objconfig"
)

type Config struct {
	Bucket               string `json:"bucket"`
	AccessKey            string `json:"region"`
	SecretKey            string `json:"access_key,omitempty"`
	Region               string `json:"secret_key,omitempty"`
	RoleARN              string `json:"role_arn,omitempty"`
	ServerSideEncryption bool   `json:"server_side_encryption,omitempty"`
	Compress             bool   `json:"compress,omitempty"`
}

func NewS3Config(bucket, accessKey, secretKey, region string) *Config {
	return &Config{
		Bucket:               bucket,
		AccessKey:            accessKey,
		SecretKey:            secretKey,
		Region:               region,
		ServerSideEncryption: true,
		Compress:             false,
	}
}

func (c *Config) Type() string {
	return string(objconfig.S3)
}

func (c *Config) BucketName() string {
	return c.Bucket
}

func (c *Config) Settings() (map[string]interface{}, error) {
	return objconfig.ToMap(c)
}
