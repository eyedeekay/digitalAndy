reg:
	go build ./
droid:
	gomobile build -o digitalandy.apk ./android/

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

install:
	go build ./

droidinstall:
	gomobile build -o digitalandy.apk ./android/
	adb install digitalandy.apk

droidrun:
	adb logcat -c
	nohup adb logcat &
	adb shell am start -n org.golang.todo.android/org.golang.app.GoNativeActivity
	sleep 3
	killall adb
	mv nohup.out debug.log

droidrm:
	adb uninstall org.golang.todo.android

droidall:
	make droidrm
	make droid
	make droidinstall
	make droidrun
