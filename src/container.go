package src

import (
	"os/exec"

	"github.com/rs/zerolog/log"
)

type Builder interface {
	BuildExportDockerImage(context, dockerFile, targetDirectory string) error
}

type BuildHandler struct{}

func (buildClient *BuildHandler) BuildExportDockerImage(context, dockerFile, targetDirectory string) error {
	cmd := exec.Command("docker", "buildx", "build", "-f", dockerFile, "--output", targetDirectory, context)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Err(err).Msgf("docker buildx error : %s", output)
		return err
	}
	log.Info().Msgf("Build Result %s", string(output))
	return nil
}
