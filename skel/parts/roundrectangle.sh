#/bin/bash
roundrectangle(){
        for POINT_X in `seq 1 $RECT_W`; do
                for POINT_Y in `seq 1 $RECT_H`; do
                        PPOINT_X=`expr $STARTX + $POINT_X - 2`
                        echo 1
                        PPOINT_Y=`expr $STARTY + $POINT_Y - 2`
                        echo 2
                        echo "point;X $POINT_X;Y $POINT_Y; $POINT_COLORS"
                        echo 3
                done
        done
        RECT_WW=`expr $RECT_W - 1`
        echo 4
        RECT_YY=`expr $RECT_Y - 1`
        echo 5
        for RPOINT_X in `seq 1 $RECT_WW`; do
                for RPOINT_Y in `seq 1 $RECT_YY`; do
                        PPOINT_X=`expr $STARTX + $RPOINT_X`
                        echo 6
                        PPOINT_Y=`expr $STARTY + $RPOINT_Y`
                        echo 7
                        echo "point;X $RPOINT_X;Y $RPOINT_Y; $POINT_COLORS"
                        echo 8
                        #None of this is going to work. Revise tomorrow.
                done
        done
}
