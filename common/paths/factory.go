// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"github.com/andreaskoch/allmark2/common/route"
)

type PatherFactory interface {
	Absolute(prefix string) Pather
	Relative(baseRoute *route.Route) Pather
}