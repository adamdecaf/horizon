#!/bin/bash

INFILE=$1
OUTFILE=$INFILE"-with-uuids"

if [[ ! -e $INFILE ]];
then
    echo "No file $INFILE was found"
    exit 1
fi

echo '' > $OUTFILE

cat $INFILE | while read line
do
    UUID=`uuidgen | tr '[:upper:]' '[:lower:]'`
    echo "('$UUID', '$line')," >> $OUTFILE
done
