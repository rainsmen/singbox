---
name: sing-box-migration
description: Migration of MatsuriDayo sing-box customization to official 1.13.x for NekoBoxForAndroid compatibility
type: project
---

## Overview
User needed to update MatsuriDayo/sing-box (a fork customized for NekoBoxForAndroid) to get latest features from official SagerNet/sing-box v1.13.x while preserving NekoBox-specific customizations.

**Status**: Completed on 2026-05-22

**Migration result**: https://github.com/rainsmen/singbox (branch: 1.13.x-neko)

## Customization modules migrated

### nekoutils package
- `callback.go`: Selector_OnProxySelected callback for UI update notification
- `srs.go`: GetGeoIPHeadlessRules/GetGeoSiteHeadlessRules for dynamic rule loading

### boxapi package
- `go_net.go`: DialContext for proxy network connections
- `go_http.go`: CreateProxyHttpClient for HTTP client through proxy
- `v2ray_server.go`: SbV2rayServer wrapper
- `v2ray_stats_service.go`: Traffic statistics implementation

### Modified files
- `protocol/group/selector.go`: Added callback invocation in SelectOutbound method
- `route/rule/rule_set_local.go`: Added geoip: and geosite: path prefix support

## Technical notes
- Base version: official sing-box v1.13.12 (stable branch)
- sing library upgraded from v0.7.18 to v0.8.10 (bufio functions compatible)
- Go version required: 1.24.7+ (tested with Go 1.26.3)

## NekoBoxForAndroid integration
The libcore module needs to implement nekoutils callback functions:
- Selector_OnProxySelected: Update UI when proxy selection changes
- GetGeoIPHeadlessRules: Load built-in GeoIP data
- GetGeoSiteHeadlessRules: Load built-in GeoSite data