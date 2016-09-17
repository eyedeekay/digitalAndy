#/bin/bash
circle(){
        for i in `seq 1 $CIRCLE_DIAMETER`; do
                echo "point;X $POINT_X;Y $POINT_Y; $POINT_COLORS"
        done
}
