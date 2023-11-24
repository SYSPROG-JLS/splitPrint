# splitPrint README

An add-on to mvs-tk5 to create separate print files from jobs printed on JES printers

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

## Installation

```sh

Download splitPrint.zip

unzip to any folder of your choice

copy the appropriate binary to the mvs-tk5 root
and rename to splitPrint (Windows users just 
copy splitPrint.exe to the mvs-tk5 root)

edit mvs-tk5/conf/tk5.cnf and make the following change:
000E 1403 prt/prt00e.txt ${TK5CRLF}
to
000E 1403 "|splitPrint" ${TK5CRLF}       for Windows users
or 
000E 1403 "|./splitPrint" ${TK5CRLF}     for all other OSs

In mvs-tk5 cuu 00E corresponds to PRT1 in JES
and 00F corresponds to PRT2. You can use either printer -
I just prefer PRT1. To use PRT2 instead just change the
line in tk5.cnf that defines 000F.

```

## Usage

- Out of the box MVS-TK5 is configured for all output 
  from a JES printer ending up in prt00e.txt or prt00f.txt. 
  I prefer having the spooled print output from each job 
  end up in a file of its own. splitPrint will do this for you.
  When installed properly, your print output will still
  go to the prt folder but each job will have a unique name.

  For example:

  TSO-STC00165-12_53_32.txt
  MF1-STC00182-10_08_12.txt

- Please note:
  splitPrint will only work properly if you have NOT
  changed the JES2 configuration. In particular 
  the parameters that control how the print file
  banners are formatted. This is because the
  splitPrint Go program has certain tests for end
  of report hardcoded, and changes to how the banners
  are formatted will prevent splitPrint from breaking
  the print output properly. 


## Support

Please [open an issue](https://github.com/SYSPROG-JLS/splitPrint/issues) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/SYSPROG-JLS/splitPrint/pulls).
