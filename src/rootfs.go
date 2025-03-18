package src

import (
	"os"

	"github.com/rs/zerolog/log"
)

type RootFS interface {
	CreateFileDD(size int64, fileName string) error
}

type RootFSHandler struct{}

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
