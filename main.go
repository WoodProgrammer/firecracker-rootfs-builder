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
func main() {
	parserCli := newParserClient()
	buildCLi := newBuildClient()

	result, err := parserCli.ParseYamlFile("config.yaml")
	if err != nil {
		log.Err(err).Msg("Error while running parserCli.ParseYamlFile()")
	}

	buildCLi.BuildExportDockerImage(result.Context, result.DockerfilePath, result.TargetDirectory)

}
