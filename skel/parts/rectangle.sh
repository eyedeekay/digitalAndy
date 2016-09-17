#/bin/bash
rectangle(){
        for POINT_X in `seq 1 $RECT_W`; do
                for POINT_Y in `seq 1 $RECT_H`; do
                        PPOINT_X=`expr $STARTX + $POINT_X`
                        PPOINT_Y=`expr $STARTY + $POINT_Y`
                        echo "point;X $POINT_X;Y $POINT_Y; $POINT_COLORS"
                done
        done
}
