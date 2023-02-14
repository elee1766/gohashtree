/*
MIT License

# Copyright (c) 2021 Prysmatic Labs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package gohashtree

import (
	"fmt"
)

func _hash(digests *byte, p *byte, count uint32)

func HashBuf(digests ChunkContainer, chunks ChunkContainer) error {
	chunkCount := chunks.WordCount()
	digestsCount := digests.WordCount()
	if chunkCount == 0 {
		return nil
	}
	if chunkCount%2 != 0 {
		return fmt.Errorf("odd number of chunks")
	}
	if digestsCount < chunkCount/2 {
		return fmt.Errorf("not enough digest length, need at least %v, got %v", chunks.WordCount()/(2), digests.WordCount())
	}
	if supportedCPU {
		_hash(digests.Ref(), chunks.Ref(), uint32(chunkCount/2))
	} else {
		cast, ok := chunks.(HashBuffer)
		if !ok {
			return fmt.Errorf("chunks does not implement HashBuffer and no cpu features detected")
		}
		cast2, ok := digests.(HashBuffer)
		if !ok {
			return fmt.Errorf("chunks does not implement HashBuffer and no cpu features detected")
		}
		sha256_1_generic(cast2, cast)
	}
	return nil
}

func Hash(digests [][32]byte, chunks [][32]byte) error {
	return HashBuf(ArraySlice(digests), ArraySlice(chunks))
}

func HashFlat(digests []byte, chunks []byte) error {
	return HashBuf(Bytes(digests), Bytes(chunks))
}

func HashChunksBuf(digests ChunkContainer, chunks ChunkContainer) {
	_hash(
		digests.Ref(),
		chunks.Ref(),
		uint32(chunks.WordCount()/2),
	)
}
func HashChunksFlat(digests []byte, chunks []byte) {
	HashChunksBuf(Bytes(digests), Bytes(chunks))
}
func HashChunks(digests [][32]byte, chunks [][32]byte) {
	HashChunksBuf(ArraySlice(digests), ArraySlice(chunks))
}
