// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package header

import (
	"fmt"
	"net/http"
)

const STATICCONTENT_CACHEDURATION_SECONDS = 86400
const DYNAMICCONTENT_CACHEDURATION_SECONDS = 120

func Cache(w http.ResponseWriter, r *http.Request, seconds int) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
}

func NoCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
}

func ContentType(w http.ResponseWriter, r *http.Request, contentType string) {
	w.Header().Set("Content-Type", contentType)
}