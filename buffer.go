package gohashtree

type ChunkContainer interface {
	// Expects a pointer to the first element of a contiguous array of Bytes
	// The length of the array should be a multiple of 32
	Ref() *byte
	// WordCount should return the amount of 32 byte chunks in the container
	WordCount() int
}

// HashBuffer is a ChunkContainer that also supports falling back to a generic go implementation
type HashBuffer interface {
	ChunkContainer

	// ChunkRef should return a 32 byte slice to the kth chunk
	ChunkRef(k int) []byte
	GrabWord1(k, j int) uint32
	GrabWord2(k, j int) uint32
}

func ToBuf(v interface{}) HashBuffer {
	switch k := v.(type) {
	case HashBuffer:
		return k
	case []byte:
		return (Bytes)(k)
	case [][32]byte:
		return (ArraySlice)(k)
	default:
		panic("unsupported type for HashBuffer")
	}
}

type Bytes []byte

func (p Bytes) Ref() *byte {
	return &p[0]
}
func (p Bytes) WordCount() int {
	return len(p) / 32
}

func (p Bytes) GrabWord1(k, j int) uint32 {
	return uint32(p[2*k*32+j])<<24 | uint32(p[2*k*32+j+1])<<16 | uint32(p[2*k*32+j+2])<<8 | uint32(p[2*k*32+j+3])
}
func (p Bytes) GrabWord2(k, j int) uint32 {
	return uint32(p[(2*k+1)*32+j])<<24 | uint32(p[(2*k+1)*32+j+1])<<16 | uint32(p[(2*k+1)*32+j+2])<<8 | uint32(p[(2*k+1)*32+j+3])
}
func (p Bytes) ChunkRef(k int) []byte {
	return p[k*32 : k*32+32]
}

type ArraySlice [][32]byte

func (p ArraySlice) Ref() *byte {
	return &p[0][0]
}
func (p ArraySlice) WordCount() int {
	return len(p)
}

func (p ArraySlice) GrabWord1(k, j int) uint32 {
	return uint32(p[2*k][j])<<24 | uint32(p[2*k][j+1])<<16 | uint32(p[2*k][j+2])<<8 | uint32(p[2*k][j+3])
}
func (p ArraySlice) GrabWord2(k, j int) uint32 {
	return uint32(p[(2*k + 1)][j])<<24 | uint32(p[(2*k + 1)][j+1])<<16 | uint32(p[(2*k + 1)][j+2])<<8 | uint32(p[(2*k + 1)][j+3])
}
func (p ArraySlice) ChunkRef(k int) []byte {
	return p[k][:]
}
