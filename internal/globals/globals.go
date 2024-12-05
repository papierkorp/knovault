package globals

import (
    "knovault/internal/types"
)

var manager types.AssetManager

func SetAssetManager(m types.AssetManager) {
    manager = m
}

func GetAssetManager() types.AssetManager {
    return manager
}