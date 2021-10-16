rm -rf ./pack-output
mkdir ./pack-output
mkdir ./pack-output/static
go build main.go
cp ./main ./pack-output/launcher
cp ./config.yaml ./pack-output/config.yaml
