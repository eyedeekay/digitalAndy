#/bin/bash
square(){
        for POINT_X in `seq 1 $RECT_W`; do
                echo "point;X $POINT_X;Y $POINT_Y; $POINT_COLORS"
        done
}
