# Colorblindor

Colorblindor is a automatic generator to 3 colorblind types, Deuteranopia, Protanopia, and Tritanopia.

# Dependencies

You will need `python3` and `golang` installed in your machine.

# Install

Run `pip3 install daltonize img2pdf`

and `go install github.com/edersonferreira/colorblindor`

# Usage

You can run the command `colorblindor` with the PDF file as a argument. Like:

`colorblindor myPdf.pdf`

And the colorblindor will generate 3 files, `deltan.pdf`, `protan.pdf` and `tritan.pdf`.

You can open these files, and see if the colors are fine to your case. 

The output will be like this:

```
$ colorblindor myPdf.pdf
Converting PDF to images...
PDF converted to images.
deltan : 1 ° Page.
deltan : 2 ° Page.
deltan : 3 ° Page.
deltan : 4 ° Page.
deltan : 5 ° Page.
protan : 1 ° Page.
protan : 2 ° Page.
protan : 3 ° Page.
protan : 4 ° Page.
protan : 5 ° Page.
tritan : 1 ° Page.
tritan : 2 ° Page.
tritan : 3 ° Page.
tritan : 4 ° Page.
tritan : 5 ° Page.
Finish! Let's check the files deltan.pdf, protan.pdf and tritan.pdf
```
