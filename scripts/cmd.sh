#!/bin/bash
path=$1
cd $path

shift # Drop 1st argument
exec $@

cd -
