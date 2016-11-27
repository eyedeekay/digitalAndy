#! /bin/sh
FIND=$(find skel/ -name *.txt.folder)
for folder in $FIND; do
	FILE=$(echo $folder | sed 's|skel/||' | sed 's|.txt.folder||')
	echo $FILE
	COUNT=0
	for file in $(find $folder -name *.png); do
		COUNT=$((COUNT+1))
		mv $file skel/$FILE.txt.folder/$FILE-$COUNT.png
		echo $file skel/$FILE.txt.folder/$FILE-$COUNT.png
	done
	mv $folder skel/$FILE
	echo $folder skel/$FILE
done