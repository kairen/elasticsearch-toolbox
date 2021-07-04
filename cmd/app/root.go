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
	goflag "flag"
	"os"

	"github.com/kairen/elasticsearch-toolbox/pkg/elastic"
	"github.com/spf13/cobra"
)

var (
	cfg        = &elastic.Config{TLS: elastic.TLSConfig{}}
	retryCount int
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "elasticsearch-toolbox [command] [flags]",
		Short: "Maintaining and operating an Elasticsearch cluster.",
	}

	cmd.AddCommand(newRotateCmd())
	cmd.AddCommand(newSnapshotCmd())
	cmd.AddCommand(newRepositoryCmd())
	cmd.AddCommand(newVersionCmd())

	cmd.PersistentFlags().StringSliceVarP(&cfg.Servers, "endpoints", "", []string{"http://elasticsearch:9200"}, "Endpoints of elasticsearch.")
	cmd.PersistentFlags().StringVarP(&cfg.Username, "username", "", os.Getenv("ELASTIC_USERNAME"), "Username for basic auth.")
	cmd.PersistentFlags().StringVarP(&cfg.Password, "password", "", os.Getenv("ELASTIC_PASSWORD"), "Password for basic auth.")
	cmd.PersistentFlags().BoolVarP(&cfg.Sniffer, "sniffer", "", false, "Enable client to use a sniffing process for finding all nodes of your cluster.")
	cmd.PersistentFlags().StringVarP(&cfg.TLS.CaPath, "tls.ca", "", "", "SSL Certificate Authority file used to secure elasticsearch communication.")
	cmd.PersistentFlags().StringVarP(&cfg.TLS.CertPath, "tls.cert", "", "", "SSL certification file used to secure elasticsearch communication.")
	cmd.PersistentFlags().StringVarP(&cfg.TLS.KeyPath, "tls.key", "", "", "SSL key file used to secure elasticsearch communication.")
	cmd.PersistentFlags().BoolVarP(&cfg.TLS.SkipHostVerify, "tls.skip-host-verify", "", false, "(insecure) Skip server's certificate chain and host name verification")
	cmd.PersistentFlags().IntVarP(&retryCount, "retry-count", "", 5, "The number of retry for deleting request.")
	cmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	return cmd
}
