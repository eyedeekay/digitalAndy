#! /bin/bash
rm -rf output
mkdir output
for files in $(find skel -name *.txt); do
  mkdir output
  for ((x = 0 ; x <= 3 ; x++)); do
    echo "path=skel/colors.txt" > config.txg
    echo "path=$files" >> config.txg
    echo $x
    ./digitalandy &> /dev/null
  done
  mv output $files.folder
done
mkdir output

echo "path=skel/colors.txt" > config.txg
for i in $(find skel/ -name dark*.txt); do
	echo "path=$i" >> config.txg
done
./digitalandy