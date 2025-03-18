package src

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

type RootFS interface {
	CreateFileDD(size int64, fileName string) error
	FormatandMountFileSystem(path, targetDirectory string) error
}

type RootFSHandler struct{}

func (rootfs *RootFSHandler) FormatandMountFileSystem(path, targetDirectory string) error {
	cmd := exec.Command("mkfs.ext4", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Err(err).Msg("Error running mkfs.ext4:")
		log.Info().Msgf("Output: %s", string(output))
		return err
	}

	log.Info().Msgf("Output: %s", string(output))

	cmd = exec.Command(fmt.Sprintf("mount", path, targetDirectory))
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Err(err).Msg("Error running mount")
		log.Info().Msgf("Mount exec output: %s", string(output))
		return err
	}

	log.Info().Msgf("Mount exec output: %s", string(output))
	return nil
}

func (rootfs *RootFSHandler) CreateFileDD(size int64, fileName string) error {

	size = int64(size * 1024 * 1024) // in MiB files

	file, err := os.Create(fileName)
	if err != nil {
		log.Err(err).Msg("Error creating file:")
		return err
	}
	defer file.Close()

	zeroBlock := make([]byte, 1024*1024) // 1 MiB bufferred files in there
	for i := 0; i < 50; i++ {
		_, err := file.Write(zeroBlock)
		if err != nil {
			log.Err(err).Msg("Error writing to file:")
			return err
		}
	}

	log.Info().Msg("rootfs.ext4 created successfully")
	return nil
}
