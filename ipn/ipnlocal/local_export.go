package ipnlocal

import (
	"sync/atomic"

	"github.com/sagernet/tailscale/net/dns/resolver"
	"github.com/sagernet/tailscale/wgengine"
	"github.com/sagernet/tailscale/wgengine/filter"
)

// ExportFilter 為 upstream/stable 側独有、我方 dev 線沒有的匯出
// （取捨規則：一方有一方無 → 補上），實作沿用 stable 原文。
func (b *LocalBackend) ExportFilter() *atomic.Pointer[filter.Filter] {
	return &b.currentNode().filterAtomic
}

func (b *LocalBackend) ExportEngine() wgengine.Engine {
	return b.e
}

func (b *LocalBackend) ExportMagicDNSHosts() resolver.MagicDNSHosts {
	return magicDNSHosts{b}
}

func (b *LocalBackend) SetExternalSSHHostKeys(keys []string) {
	b.mu.Lock()
	b.externalSSHHostKeys = keys
	if b.hostinfo != nil {
		b.hostinfo.SSH_HostKeys = keys
	}
	b.mu.Unlock()
	b.doSetHostinfoFilterServices()
}
