// +build server
// +build !linux

package ldcache

type LDCache struct {
}

func Open() (*LDCache, error) {
	return &LDCache{}, nil
}

func (c *LDCache) Close() error {
	return nil
}

func (c *LDCache) Lookup(libs ...string) (paths32, paths64 []string) {
	return []string{}, []string{}
}
