# Morse Images

Encodes a given message into morse which is then rendered as a series of
images. This allows for a simple encoding to be hidden in plain site and may be
used as a gentle introduction to code-based puzzles.

```sh
go get -u github.com/logank/morse_images/cmd/morse_images/...
morse_images "I want to hide this in plain sight"
```

## Building

### Replacing Image Assets

Images assets are named:

* Dah_Dah.png
* Dah_Dit.png
* Dah_Shh.png
* Dit_Dah.png
* Dit_Dit.png
* Dit_Shh.png
* Shh_Dah.png
* Shh_Dit.png
* Shh_Shh.png

They are encoded using go-bindata:

```sh
$ go get -u github.com/shuLhan/go-bindata/...
```

Once bindata is installed, copy the above images into the source directory and
regenerate via `go generate`.
