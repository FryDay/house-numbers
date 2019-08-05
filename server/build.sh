#!/bin/bash
sudo docker build --tag rpi-ws281x-go-builder .
sudo docker run --rm -ti -v "$(pwd)":/go/src/house-numbers rpi-ws281x-go-builder /usr/bin/qemu-arm-static /bin/sh -c "go build -o src/house-numbers/house-numbers -v house-numbers"