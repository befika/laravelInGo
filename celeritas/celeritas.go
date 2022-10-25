package celeritas

type Celeritas struct {
	AppName    string
	AppVersion string
	Debug      bool
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderName: []string{
			"cmd",
			"internal/adapter",
			"internal/adapter/storage",
			"internal/adapter/persistence",
			"internal/adapter/http",
			"internal/adapter/http/server",
			"internal/adapter/http/middleware",
			"internal/glue",
			"internal/constant",
			"internal/model",
			"platform",
			"platform/client",
			"platform/logger",
			"assets",
			"config",
		},
	}

	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil

}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, folder := range p.folderName {
		err := c.CreateDirIfNotExist(root + "/" + folder)
		if err != nil {
			return err
		}
	}

	return nil

}
