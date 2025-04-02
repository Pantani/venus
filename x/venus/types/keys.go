package types

const (
	// ModuleName defines the module name
	ModuleName = "venus"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_venus"
)

var (
	ParamsKey = []byte("p_venus")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
