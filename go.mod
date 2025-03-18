module firecracker-vmbuilder

replace github.com/WoodProgrammer/firecracker-vmbuilder/src => ./src

go 1.24.1

require (
	github.com/WoodProgrammer/firecracker-vmbuilder/src v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.33.0
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/sys v0.12.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
