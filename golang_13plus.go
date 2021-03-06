// Copyright 2015 ENDOH takanao.
// <https://github.com/MiCHiLU/go-gob-http>
//
// Copyright 2014 Orchestrate, Inc.
// <https://github.com/orchestrate-io/dvr>
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

// +build go1.3

package gobhttp

import (
	"net/http"
)

// This file contains functions calls that will be put in place with golang
// 1.3 or higher.

// This call wraps copying the TLS value since it only showed up in golang
// 1.3 and higher.
func newGobRequestVS(req *http.Request, r *GobRequest) {
	r.TLS = req.TLS
}

// This call wraps copying the TLS value since it only showed up in golang
// 1.3 and higher.
func newGobResponseVS(resp *http.Response, r *GobResponse) {
	r.TLS = resp.TLS
}
