### Image-resizer

This is a simple batch image resizer written in Golang. It will take images from a directory, resize them, and save them in another directory.

This program runs on Linux, Windows and Mac. Downloads are available on the releases tab.

#### Usage

##### On Windows

##### On Linux

##### On Mac
I don't have access to a Mac :(

##### Building on Linux

###### For windows:
```bash
GOOS=windows GOARCH=amd64 go build -o image_resizer.exe image_resizer.go
```

###### For Linux
```bash
go build image_resizer.go 
```

###### For Mac
```bash
GOOS=darwin GOARCH=amd64 go build image_resizer.go
```

Built with `go version go1.13.4 linux/amd64`

