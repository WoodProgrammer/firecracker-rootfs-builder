package main

import (
	"os"

	src "github.com/WoodProgrammer/firecracker-vmbuilder/src"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	configFile string
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

func HandleRootFS() {
	parserClient := newParserClient()
	buildCLient := newBuildClient()
	rootFsClient := newRootFSClient()

	result, err := parserClient.ParseYamlFile(configFile)
	if err != nil {
		log.Err(err).Msg("Error while running parserCli.ParseYamlFile()")
	}

	err = rootFsClient.CreateFileDD(10, "ops")
	if err != nil {
		log.Err(err).Msg("Error while running rootFsClient.CreateFileDD()")
	}

	fsErr := rootFsClient.FormatandMountFileSystem("ops", result.TargetDirectory)
	if fsErr != nil {
		log.Err(fsErr).Msg("Error while running rootFsClient.FormatFileSystem()")
	}

	err = buildCLient.BuildExportDockerImage(result.Context, result.DockerfilePath, result.TargetDirectory)
	if err != nil {
		log.Err(err).Msg("Error while running buildCLient.BuildExportDockerImage()")
	}

}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "rootfsCreator",
		Short: "CLI tool to manage RootFS for firecracker micro VMs",
		Run: func(cmd *cobra.Command, args []string) {
			HandleRootFS()
		},
	}
	rootCmd.Flags().StringVarP(&configFile, "config", "C", "config.yaml", "Config file of RootFS creation")

	rootCmd.MarkFlagRequired("config")

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("rootfsCreator execution failed")
		os.Exit(1)
	}
}
