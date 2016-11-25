#! /bin/sh
FILES_LIST=""
CWD=$(pwd)
STARTX=0
STARTY=0
for filesh in $(find "$CWD/skel/parts/" -name '*.sh'); do
        . $filesh
        echo $filesh
        FILES_LIST="$FILES_LIST $filesh"
done
echo "Init at X: $STARTX"
echo "Init at Y: $STARTY"
