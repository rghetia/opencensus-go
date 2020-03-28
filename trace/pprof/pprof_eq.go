// Copyright 2019, OpenCensus Authors
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

package main

import (
	"fmt"
	_ "net/http/pprof"
	"go.opencensus.io/trace"
	//_ "github.com/pkg/profile"
	"net/http"
)

func init() {
}
var qs []*trace.EQ
func save(q *trace.EQ) {
	qs = append(qs,q)
}
func main() {
	//defer profile.Start(profile.MemProfile).Stop()
	for i := 0 ; i < 1000000 ; i++ {
		c := 30
		q := trace.NewEQ(30)
		save(q)
		for j := 0 ; j < 4*c ; j++ {
			q.Add(j)
		}
	}
	fmt.Println("work done")
	http.ListenAndServe("localhost:8080", nil)
}
