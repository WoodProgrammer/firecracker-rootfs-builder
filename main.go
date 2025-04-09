package main

import (
	"os"

	src "github.com/WoodProgrammer/firecracker-vmbuilder/src"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	configFile string
	rootFsName string
	rootFsSize int64
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
		os.Exit(1)
	}

	err = rootFsClient.CreateFileDD(rootFsSize, rootFsName)
	if err != nil {
		log.Err(err).Msg("Error while running rootFsClient.CreateFileDD()")
		os.Exit(1)
	}

	fsErr := rootFsClient.FormatandMountFileSystem(rootFsName, result.TargetDirectory)
	if fsErr != nil {
		log.Err(fsErr).Msg("Error while running rootFsClient.FormatFileSystem()")
		os.Exit(1)
	}

	err = buildCLient.BuildExportDockerImage(result.Context, result.DockerfilePath, result.TargetDirectory)
	if err != nil {
		log.Err(err).Msg("Error while running buildCLient.BuildExportDockerImage()")
		os.Exit(1)
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
	rootCmd.Flags().StringVarP(&rootFsName, "filesystem-name", "F", "rootfs", "Name of rootfs")
	rootCmd.Flags().Int64VarP(&rootFsSize, "filesystem-size", "S", 10, "Size of rootfs")

	rootCmd.MarkFlagRequired("config")

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("rootfsCreator execution failed")
		os.Exit(1)
	}
}
