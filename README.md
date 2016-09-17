digitalAndy
===========

Rules-based pixel-art generator for procedural generation of game content, named
very affectionately after my old buddy Andy who created pixel art for me in
college. You take a plain text file containing the rules required to generate
your png file, and you feed it to a program which picks it apart and creates an
image within some more-or-less tunable parameters.

How it works:
-------------
Short version: Right now, not perfectly, it's a work-in-progress, I'm developing
according to what I need it to do instead of a clear plan. But that's part of
the point, too. What you need to do is compose a set of configuration files
according to what you need the program to do. These are non-executable text
files containing details in the form of colors and coordinates.

For example, some colors might be:

        color; green;R 0;G 255;B 0;T 255
        color; blue;R 0;G 0;B 255;T 255
        color; red;R 255;G 0;B 0;T 255

Whereas, some coordinates might be:

        point;X 5;Y 7; green
        point;X 6;Y 7; green
        point;X 7;Y 7; green
        point;X 8;Y 7; green

But, if you wanted to make images that way you may as well just use a GUI tool.
Instead, if you specify more than one color to the coordinate, one of the colors
will be selected randomly.

        point;X 5;Y 13; red green
        point;X 6;Y 13; red green
        point;X 7;Y 13; red green
        point;X 8;Y 13; red green

This way, you can use a single configuration file to create a whole bunch of
variant images. You can accomplish some shading-like stuff by grouping similar
colors together on different pixels in a sensible way, without having to rely on
creating the graphics by hand. Because I don't have time to do it myself.

Shortcomings and Things to Do:
------------------------------

I don't have a really good way of weighting the chances of a given random color.
Sometimes it still selects the wrong color from the list at mildly unpredictable
intervals(much less than it used to be.)
