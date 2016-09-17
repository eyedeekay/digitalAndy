
export HAND_WIDTH=4
export HAND_HEIGHT=5
export HAND_X=4
export HAND_Y=2

drawhand(){
        export STARTX=$HAND_X
        export STARTY=$HAND_Y
        roundrectangle $HAND_WIDTH $$HAND_HEIGHT $COLORS
}
