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

package util

import (
	"strings"
	"time"

	"github.com/golang/glog"
)

func ParseName(index string) string {
	array := strings.Split(index, "-")
	if len(array) > 2 {
		return strings.Join(array[:len(array)-1], "-")
	}
	return array[0]
}

func ParseDate(index string) string {
	array := strings.Split(index, "-")
	if len(array) >= 2 {
		return array[len(array)-1]
	}
	return ""
}

func IsExpired(expiredDay int, date, format string) (bool, error) {
	indexTime, err := time.Parse(format, date)
	if err != nil {
		return false, err
	}

	if int(time.Now().Sub(indexTime).Hours()/24) >= expiredDay {
		return true, nil
	}
	return false, nil
}

func Retry(d time.Duration, attempts int, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		err = f()
		if err == nil {
			return nil
		}
		glog.Errorf("Error: %s, Retrying in %s. %d Retries remaining.", err, d, attempts-i)
		time.Sleep(d)
	}
	return err
}
