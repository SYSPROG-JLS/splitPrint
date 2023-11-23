# README splitPrint

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

Replace the contents of `README.md` with your project's:

- Name: split
- Description
- Installation instructions
- Usage instructions
- Support instructions
- Contributing instructions
- Licence

Feel free to remove any sections that aren't applicable to your project.

## Support

Please [open an issue](https://github.com/SYSPROG-JLS/splitPrint) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/SYSPROG-JLS/splitPrint/).
