#!/bin/env bash
set -euxo pipefail

go install golang.org/dl/go1.23.2@latest
export PATH=$PATH:$HOME/go/bin
go1.23.2 download
cp $(which go1.23.2) $(which go)

apt install wget --yes
wget https://github.com/pulumi/pulumictl/releases/download/v0.0.46/pulumictl-v0.0.46-linux-amd64.tar.gz
mkdir /usr/local/pulumi
tar --directory=/usr/local/pulumi --extract --gzip --file=pulumictl-v0.0.46-linux-amd64.tar.gz
export PATH=$PATH:/usr/local/pulumi
cd provider && go mod tidy && cd -
make tfgen
make provider
make build_sdks
# Avoid error: Access to the path '/home/runner/work/pulumi-deploy/pulumi-deploy/pulumi-bitlaunch/sdk/dotnet/obj/d3c3f3c5-9946-497b-8a7e-17f0b6501f6f.tmp' is denied. [/home/runner/work/pulumi-deploy/pulumi-deploy/GithubRunner/GithubRunner.fsproj]
chmod --recursive 0777 sdk/dotnet

mkdir output && mkdir output/bin && mkdir --parents output/sdk/dotnet
cp bin/pulumi-resource-bitlaunch ./output/bin
cp --recursive sdk/dotnet ./output/sdk/dotnet

ls -l ./output

zip --recurse-paths pulumi-bitlaunch.zip ./output
