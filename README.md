digitalAndy
===========

Rules-based pixel-art generator for procedural generation of game content, named
very affectionately after my old buddy Andy who created pixel art for me in
college. You take a plain text file containing the rules required to generate
your png file, and you feed it to a program which picks it apart and creates an
image within some more-or-less tunable parameters.

How it works:
-------------
Much more cleanly than before. Here's how it works. You create a file named
config.txg, that's *.txg* not .txt, which can contain one of a few directives.
These directives are:

  * path="*file_to_include*" : This is an optional path to a file which you want
  to include in the final product.
  * lcolor *colorname* *red green blue alpha* : This is a description of a color
  to load and prepare the program to draw in. It will not be used until paired
  with an scolor directive.
  * point *X* *Y* : 
  * rect *X* *Y* *W* *H* : 
  * round *X* *Y* *W* *H* : 
  * scolor *list of loaded colors* :
