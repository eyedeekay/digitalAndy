include ../config.mk
PREFIX = /
MANPREFIX = $(PREFIX)/share/man

app:
	go build dandy.go

clean:
	rm *.out *.png output/*.png dandy skel/*.txt.folder/ output *-pak digitalandy.apk -rf

install:
	cp dandy $(DESTDIR)$(PREFIX)/usr/bin/
	cp digitalandy $(DESTDIR)$(PREFIX)/usr/bin/
	mkdir -p $(DESTDIR)$(PREFIX)/usr/share/digitalandy/skel
	cp dandy.cfg $(DESTDIR)$(PREFIX)/usr/share/digitalandy/dandy.cfg
	cp config.txg $(DESTDIR)$(PREFIX)/usr/share/digitalandy
	cp skel/*.txt $(DESTDIR)$(PREFIX)/usr/share/digitalandy/skel/

deb-pkg:
	make
	checkinstall --install=no \
		--default \
		--pkgname=digitalandy \
		--pkgversion=1 \
		--deldoc=yes \
		--deldesc=yes \
		--backup=no \
		--pakdir=../

apk-pkg:
	gomobile build ./
