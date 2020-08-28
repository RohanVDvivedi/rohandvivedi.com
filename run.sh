#!/bin/sh

# go the application directory
cd ~/go/src/rohandvivedi.com

# build javascript and css bundles, for the react frontend client
npm run build_bundles

# compile and run go backend
go run src/main.go "$@"