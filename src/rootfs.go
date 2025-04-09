package src

import (
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

	err := os.Mkdir(targetDirectory, os.ModeDir)
	if err != nil {
		log.Err(err).Msg("Error while calling os.Mkdir")
		return err
	}

	cmd := exec.Command("mkfs.ext4", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Err(err).Msg("Error running mkfs.ext4:")
		log.Info().Msgf("Output: %s", string(output))
		return err
	}

	log.Info().Msgf("Output of mkfs.ext4 : %s", string(output))

	cmd = exec.Command("mount", path, targetDirectory)
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Err(err).Msg("Error running mount command")
		log.Info().Msgf("Mount errored status: %s", string(output))
		return err
	}

	log.Info().Msgf("Mount status: %s", string(output))
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
