package assets

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func AssetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo}
}
