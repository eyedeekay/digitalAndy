digitalAndy
===========

Rules-based pixel-art generator for procedural generation of game content, named
very affectionately after my old buddy Andy who created pixel art for me in
college. You take a plain text file containing the rules required to generate
your png file, and you feed it to a program which picks it apart and creates an
image within some more-or-less tunable parameters.

Background:
-----------

I've been working on a video game for a long time which makes use of procedural
content generation.

How it works:
-------------

Much more cleanly than before. Here's how it works. You create a file named
config.txg, that's *.txg* not .txt, and that file can contain one of a few
directives. These directives are:

  * path="*file_to_include*" : This is an optional path to a file which you want
  to include in the final product.
  * lcolor *colorname* *red green blue alpha* : This is a description of a color
  to load and prepare the program to draw in. It will not be used until paired
  with an scolor directive.
  * point *X* *Y* : This is the coordinates of a simple point. When paired with
  an scolor directive on the same line, a point of the scolor will be drawn at
  the described coordinates.
  * rect *X* *Y* *W* *H* : This draws a rectangle when paired with an scolor on
  the same line.
  * round *X* *Y* *W* *H* : This draws a round-ish graphic when paired with an
  scolor on the same line.
  * scolor *list of loaded colors* : This is a list of colors, one of which will
  be selected per defined region. The color that is selected will be picked
  completely at random.

It's still a little confusing to use. Basically dandy.go produces the executable
dandy, which takes config.txg and uses the rules to generate a single piece of
art in the output/ folder with a random filename. It's then possible to write
scripts which generate large numbers of pieces of art from a single rule like
in the example scripts(build.sh and rename.sh).

To-Do List/Ongoing Improvements:
--------------------------------

More shapes(Triangles, then Ellipses)

Per-OS wrapper Scripts, launchers, interfaces as appropriate(Run on Windows,
OSX, Android)

Add into Jenkins

Integrate into web site
