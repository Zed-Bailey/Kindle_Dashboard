#!/bin/sh

if [ -f "bundle.zip" ]; then
    echo "removing old bundle"
    rm -f bundle.zip
fi


# bundles all the required files into a single zip
zip -r bundle.zip dashboard