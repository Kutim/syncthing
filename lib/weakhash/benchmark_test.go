// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package weakhash

import (
	"os"
	"testing"

	"github.com/syncthing/syncthing/lib/protocol"
)

const testFile = "../model/testdata/~syncthing~file.tmp"

func BenchmarkWeakHash(b *testing.B) {
	const size = 16384
	data := make([]byte, size)
	hf := NewHash(protocol.BlockSize)

	for i := 0; i < b.N; i++ {
		hf.Write(data)
	}

	_ = hf.Sum32()
	b.SetBytes(size)
}

func BenchmarkFind1MFile(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1 << 20)
	for i := 0; i < b.N; i++ {
		fd, err := os.Open(testFile)
		if err != nil {
			b.Fatal(err)
		}
		_, err = Find(fd, []uint32{0, 1, 2}, 128<<10)
		if err != nil {
			b.Fatal(err)
		}
		fd.Close()
	}
}
