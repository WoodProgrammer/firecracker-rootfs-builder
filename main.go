package main

import (
	src "github.com/WoodProgrammer/firecracker-vmbuilder/src"
	"github.com/rs/zerolog/log"
)

func newBuildClient() src.Builder {
	return &src.BuildHandler{}
}

func newParserClient() src.Parser {
	return &src.ParseHandler{}
}

func newRootFSClient() src.RootFS {
	return &src.RootFSHandler{}
}
func main() {
	parserClient := newParserClient()
	buildCLient := newBuildClient()
	rootFsClient := newRootFSClient()

	result, err := parserClient.ParseYamlFile("config.yaml")
	if err != nil {
		log.Err(err).Msg("Error while running parserCli.ParseYamlFile()")
	}

	err = buildCLient.BuildExportDockerImage(result.Context, result.DockerfilePath, result.TargetDirectory)
	if err != nil {
		log.Err(err).Msg("Error while running buildCLient.BuildExportDockerImage()")
	}

	err = rootFsClient.CreateFileDD(10, "ops")
	if err != nil {
		log.Err(err).Msg("Error while running rootFsClient.CreateFileDD()")
	}
	// Extract rootfs components now

}
