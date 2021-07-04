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

package elastic

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/kairen/elasticsearch-toolbox/pkg/objconfig"
	"github.com/olivere/elastic/v7"
)

const timeout = time.Second * 5

type Client struct {
	raw *elastic.Client
	ctx context.Context
}

func NewClient(cfg Config) (*Client, error) {
	if len(cfg.Servers) < 1 {
		return nil, errors.New("No servers specified")
	}

	options, err := cfg.getConfigOptions()
	if err != nil {
		return nil, err
	}

	rawClient, err := elastic.NewClient(options...)
	if err != nil {
		return nil, err
	}

	if cfg.Version == 0 {
		// Determine ElasticSearch Version
		pingResult, _, err := rawClient.Ping(cfg.Servers[0]).Do(context.Background())
		if err != nil {
			return nil, err
		}
		esVersion, err := strconv.Atoi(string(pingResult.Version.Number[0]))
		if err != nil {
			return nil, err
		}
		glog.Infof("Elasticsearch version detected: %d", esVersion)
		cfg.Version = uint(esVersion)
	}
	return &Client{
		raw: rawClient,
		ctx: context.Background(),
	}, nil
}

func (c *Client) CatIndices(index string) (elastic.CatIndicesResponse, error) {
	return c.raw.CatIndices().Index(index).Do(c.ctx)
}

func (c *Client) DeleteIndex(index string) (*elastic.IndicesDeleteResponse, error) {
	return c.raw.DeleteIndex(index).Do(c.ctx)
}

func (c *Client) CreateSnapshotRepository(provider objconfig.Interface) (*elastic.SnapshotCreateRepositoryResponse, error) {
	settings, err := provider.Settings()
	if err != nil {
		return nil, err
	}
	return c.raw.SnapshotCreateRepository(provider.BucketName()).
		Type(provider.Type()).
		Settings(settings).
		Do(c.ctx)
}

func (c *Client) CreateSnapshot(repo, snapshot, indices string) (*elastic.SnapshotCreateResponse, error) {
	obj := &struct {
		Indices            string `json:"indices"`
		IgnoreUnavailable  bool   `json:"ignore_unavailable"`
		IncludeGlobalState bool   `json:"include_global_state"`
	}{
		Indices:            indices,
		IgnoreUnavailable:  true,
		IncludeGlobalState: false,
	}
	return c.raw.SnapshotCreate(repo, snapshot).
		BodyJson(obj).
		Do(c.ctx)
}
