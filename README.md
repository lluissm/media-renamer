# media-renamer

![Header](images/demo.gif)

Multiplatform cli tool to rename photos and videos according to its real creation date (the creation date of the picture, not the file).

Highly useful to unify the naming convention of photos and videos taken from different devices as well as being able to browse media in a chronological order from a file explorer by just sorting by date.

_DISCLAIMERS:_

- Only tested in Mac OS and linux so far.
- Only .jpeg and .mov file types have been configured so far (exported from Apple Photos Mac OS app)

## How to configure

The configuration is done via [config.yml](cmd/media-renamer/config.yml) present in the source code.

```yml
- extension: ".mov"
  dateFields:
    - name: "CreationDate"
      dateFormat: "2006:01:02 15:04:05-07:00"
- extension: ".jpeg"
  dateFields:
    - name: "CreateDate"
      dateFormat: "2006:01:02 15:04:05"
```

It consists of a list of fileTypes with its extension and an array of dateFields from which the date could be obtained. The date is in [golang date format](https://go.dev/src/time/format.go).

There can be more than one per fileType and the first one that matches will be used to rename the file. In case of no match, the file name will not be modified.

## How to install

### Dependencies

The tool relies on [exiftool](https://exiftool.org) being installed on the device and added to the path.

An easy way to install in Mac OS is via homebrew:

```bash
$ brew install exiftool
```

An easy way to install in Ubuntu is:

```bash
$ sudo apt install exiftool
```

Instructions to install in all the platforms can be found on the official [website](https://exiftool.org) of the tool.

### Install script for media-renamer

For linux and MacOS systems you can use the install script.

To install the latest version:

```bash
$ curl -s https://raw.githubusercontent.com/lluissm/media-renamer/master/install.sh | bash
```

To install a specific version (e.g., v.1.3.0):

```bash
$ curl -s https://raw.githubusercontent.com/lluissm/media-renamer/master/install.sh | bash -s v1.3.0
```

### Binary packages

The binary packages for Linux, Windows and macOS are uploaded for each release and can be downloaded from the [releases](https://github.com/lluissm/media-renamer/releases) page.
