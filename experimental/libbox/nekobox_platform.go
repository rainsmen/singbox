package libbox

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing/service"
)

// RegisterPlatformInterface wraps the gomobile-facing PlatformInterface into the
// internal adapter.PlatformInterface and registers it on ctx.
//
// NekoBox's libcore creates the box via box.New directly instead of going through
// NewCommandServer, so it must perform this registration itself; otherwise the
// sing-box 1.13 network manager finds no platform interface and falls back to a
// netlink-based interface monitor, which Android bans
// ("netlink socket in Android is banned by Google").
func RegisterPlatformInterface(ctx context.Context, platformInterface PlatformInterface) {
	service.MustRegister[adapter.PlatformInterface](ctx, &platformInterfaceWrapper{
		iif:       platformInterface,
		useProcFS: platformInterface.UseProcFS(),
	})
}
