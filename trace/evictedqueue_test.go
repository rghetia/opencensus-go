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

package trace

import (
	"reflect"
	"testing"
)

func init() {
}

func TestAdd(t *testing.T) {
	q := newEvictedQueue(3)
	q.add("value1")
	q.add("value2")
	if wantLen, gotLen := 2, len(q.queue); wantLen != gotLen {
		t.Errorf("got queue length %d want %d", gotLen, wantLen)
	}
}

func (eq *evictedQueue) queueToArray() []string {
	arr := make([]string, 0)
	for _, value := range eq.queue {
		arr = append(arr, value.(string))
	}
	return arr
}

func TestDropCount(t *testing.T) {
	q := newEvictedQueue(3)
	q.add("value1")
	q.add("value2")
	q.add("value3")
	q.add("value1")
	q.add("value4")
	if wantLen, gotLen := 3, len(q.queue); wantLen != gotLen {
		t.Errorf("got queue length %d want %d", gotLen, wantLen)
	}
	if wantDropCount, gotDropCount := 2, q.droppedCount; wantDropCount != gotDropCount {
		t.Errorf("got drop count %d want %d", gotDropCount, wantDropCount)
	}
	wantArr := []string{"value3", "value1", "value4"}
	gotArr := q.queueToArray()

	if wantLen, gotLen := len(wantArr), len(gotArr); gotLen != wantLen {
		t.Errorf("got array len %d want %d", gotLen, wantLen)
	}

	if !reflect.DeepEqual(gotArr, wantArr) {
		t.Errorf("got array = %#v; want %#v", gotArr, wantArr)
	}
}

func BenchmarkDropCount(b *testing.B) {
	testcases := []struct {
		name string
		cap  int
		add  int
	}{
		{
			name: "cap-5-add-10",
			cap:  5,
			add:  10,
		},
		{
			name: "cap-10-add-20",
			cap:  10,
			add:  20,
		},
		{
			name: "cap-30-add-40",
			cap:  30,
			add:  40,
		},
		{
			name: "cap-40-add-50",
			cap:  40,
			add:  50,
		},
	}

	for _, tc := range testcases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			q := newEvictedQueue(tc.cap)
			for i := 0; i < b.N; i++ {
				for i := 0; i < tc.add; i++ {
					q.add("value1")
				}
			}
		})
	}
}
