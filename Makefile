# Copyright 2012, Google Inc. All rights reserved.
# Use of this source code is governed by a BSD-style license that can
# be found in the LICENSE file.

MAKEFLAGS = -s

warc.go: warc.y
	goyacc -o warc.go warc.y
	gofmt -w warc.go

clean:
	rm -f y.output warc.go