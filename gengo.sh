#!/bin/bash

fullfile=$1
filepath=$2
pkgname=$(basename $filepath)
filename=$(basename -- "$fullfile")
extension="${filename##*.}"
filename="${filename%.*}"

if [ $extension == "vy" ]
then
    vyper -f bytecode $fullfile > $filepath/$filename.bin
    vyper -f abi $fullfile > $filepath/$filename.abi
    vyper -f external_interface $fullfile > $filepath/${filename}_interface.vy

    abigen -abi $filepath/$filename.abi -bin $filepath/$filename.bin -pkg $pkgname -out $filepath/$filename.go
else
    abigen -sol $fullfile -pkg $pkgname -out $filepath/$pkgname.go
fi

