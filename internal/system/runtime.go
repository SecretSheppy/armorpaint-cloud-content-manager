package system

func Runtime(args []string) {
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
