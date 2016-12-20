reg:
	go build ./
droid:
	gomobile build ./android/

all:
	make reg
	make droid

clean:
	rm digitalandy

cleandroid:
	rm digitalandy.apk

distclean:
	make clean
	make cleandroid
