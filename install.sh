#!/bin/sh

FILEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo $FILEDIR

PATHDIRS="
/usr/local
/opt/local
$HOME/bin
/usr/bin"

for dir in $PATHDIRS
do
    if [ "`echo $PATH | grep $dir`" ]; then
        cp -rvf "$FILEDIR/bin" "$FILEDIR/games" $dir
        break
    fi
done
