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

package main

import (
	"os"

	"github.com/golang/glog"
	"github.com/kairen/elasticsearch-toolbox/cmd/app"
	"github.com/spf13/pflag"
)

func main() {
	defer glog.Flush()
	cmd := app.New()

	pflag.CommandLine.Set("logtostderr", "true")
	flags := cmd.PersistentFlags()
	flags.ParseErrorsWhitelist.UnknownFlags = true
	flags.Parse(os.Args[1:])

	if err := cmd.Execute(); err != nil {
		glog.Errorln(err)
		os.Exit(1)
	}
}
