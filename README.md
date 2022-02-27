# Oh My Scan

Download french scans from scan-op.cc.

## Table of contents

- [Oh My Scan](#oh-my-scan)
  - [Table of contents](#table-of-contents)
  - [Overview](#overview)
    - [About the project](#about-the-project)
    - [Buit with](#buit-with)
    - [Version](#version)
  - [Getting started](#getting-started)
    - [Downloading](#downloading)
    - [Installing](#installing)
    - [How to use](#how-to-use)
      - [Download a scan](#download-a-scan)
      - [Merge a scan](#merge-a-scan)
  - [Authors](#authors)
  - [License](#license)

## Overview

### About the project

Oh My Scan is a script developed in Go that allows you to download locally your favorite french manga scans. To carry out this project, I was inspired by the initial [go-scan-op.com-scrapper](https://github.com/SegFault42/go-scan-op.com-scrapper) project of [https://github.com/SegFault42](https://github.com/SegFault42).

### Buit with

- [Go](https://golang.org/) - The language used.

### Version

**1.0.0**

## Getting started

**WARNING: If you are on Windows, your antivirus may be panicking. But don't worry, just ignore it. If you have any question about this script, don't hesitate to contact me by email or on Twitter.**

### Downloading

To download this application, please do `git clone https://github.com/tun43p/ohmyscan.git`.

### Installing

To install the project, please run `./build`.

### How to use

_If you're are on a windows system, add `.exe` after `ohmyscan`._

#### Download a scan

To download a scan in multiple images, please run this command:

```sh
ohmyscan download --name MANGA_NAME --number NUMBER
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan download --name doubt --number 1
```

If you want to merge directly the scans you just downloaded, please run :

```sh
ohmyscan download --name MANGA_NAME --number NUMBER --merge BOOLEAN
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan download --name doubt --number 1 --merge true
```

#### Merge a scan

After that, if you want to merge all images of a scan in one PDF file, please do:

```sh
ohmyscan merge --name MANGA_NAME --number NUMBER
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan merge --name doubt --number 1
```

## Authors

**tun43p** - _Initial work_ - [tun43p](https://github.com/tun43p).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
