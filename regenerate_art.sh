#! /usr/bin/env /bin/sh

for file in $(find /usr/share/digitalandy/skel/ -name *.txt); do
        if [ $file != "/usr/share/digitalandy/skel/colors.txt" ]; then
                FOLDER=$(echo $file | sed 's|/usr/share/digitalandy/skel/||' | sed 's|.txt||')
                echo skel/$FOLDER
                VAR=0
                mkdir -p "skel/$FOLDER"
                while [ $VAR != 4 ]; do
                        dandy -incl=/usr/share/digitalandy/skel/colors.txt \
                           -desc="$file" \
                           -dir=skel/$FOLDER \
                           -name=$FOLDER-$VAR.png \
                           1> /dev/null
                        VAR=$((VAR + 1))
                        echo $FOLDER-$VAR.png
                done
        fi
done
