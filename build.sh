#!/usr/bin/env bash

package_name="dt"
platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/386" "linux/amd64")

rm -r dist
rm -r tmp

mkdir dist

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    mkdir tmp
    cp README.md tmp/

    zip_name=$package_name'-'$GOOS'-'$GOARCH
    output_name=$package_name
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o tmp/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    cd tmp
    tar czf $zip_name.tgz $output_name README.md
    mv $zip_name.tgz ../dist/
    cd ..
    rm -r tmp
done
