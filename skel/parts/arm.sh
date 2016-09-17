
export ARM_WIDTH=4
export ARM_HEIGHT=20
export ARM_X=4
export ARM_Y=2

drawarm(){
        export STARTX=$ARM_X
        export STARTY=$ARM_Y
        roundrectangle $ARM_WIDTH $ARM_HEIGHT $COLORS
}
