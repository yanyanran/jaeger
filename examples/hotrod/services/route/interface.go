// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package route

import (
	"context"
	"time"
)

// Route describes a route between Pickup and Dropoff locations and expected time to arrival.
// 上车点/下车点/时间
type Route struct {
	Pickup  string
	Dropoff string
	ETA     time.Duration
}

// Interface exposed by the Driver service.
type Interface interface {
	FindRoute(ctx context.Context, pickup, dropoff string) (*Route, error)
}
