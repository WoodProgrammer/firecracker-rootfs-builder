<div><img src="docs/img/img.png" width="100"/><h1>Firecracker RootFS Builder</h1></div>

Firecracker RootFS handler is a simple project to create root-file-system with Docker images in ext4 format and allows you to run them in micro-vms by Firecracker.

## Features
- Convert Docker images to rootFS 
- Store them locally (Registry soon ...)

## Installation

To install OpenMetricMigrator, clone the repository and build the project:

```sh
# Clone the repository
git clone https://github.com/WoodProgrammer/firecracker-rootfs-builder.git

cd firecracker-rootfs-builder

go build -o fco .

mv fco /usr/local/bin
```

## Usage

Run the tool with the required options:

```sh

cat <<EOF>config.yaml

image: alpine
docker_file: Dockerfile
context: "."
target_directory: "alpine-rootfs"

EOF

./fco -C config.yaml
```

Config file contains docker image spec and rootfs details.Apart from these details you can basically pass the rootFS  size and name as well. (Needs to improve)

### Available Flags


```sh

CLI tool to manage RootFS for firecracker micro VMs

Usage:
  rootfsCreator [flags]

Flags:
  -C, --config string            Config file of RootFS creation (default "config.yaml")
  -F, --filesystem-name string   Name of rootfs (default "rootfs")
  -S, --filesystem-size int      Size of rootfs (default 10)
  -h, --help                     help for rootfsCreator

```

## Example

Convert an OpenMetrics file to a Prometheus-compatible format:

```sh
fco -C config.yaml
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the tool.

## Contact

For any questions or feedback, feel free to open an issue on GitHub.

## TO DO ;

* Adjustable command args
* Registry ability
* More use cases on test suits
