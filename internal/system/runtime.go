package system

import "github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/logger"

func Runtime(args []string) {
	log := logger.Get()

	if len(args) < 3 {
		log.Panic("Not enough arguments")
		panic("Not enough arguments")
	}

	switch args[1] {
	case "install":
		DownloadAllAssets(args[2])
	case "uninstall":
		RemoveAllAssets(args[2])
	case "update":
		UpdateAssetCache(args[2])
	default:
		log.Error("Unsupported command")
	}
}
