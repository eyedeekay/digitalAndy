#! /usr/bin/env /bin/sh
for file in $(find /usr/share/digitalandy/skel/ -name *.txt); do
        if [ $file != "colors.txt" ]; then
                echo $file
                # dandy -incl=/usr/share/digitalandy/skel/colors.txt \
                #   -desc="$file" \
                #   -dir=
                #   -
        fi
done
