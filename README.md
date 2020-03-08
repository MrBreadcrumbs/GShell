# GShell
basic windows shell written in golang with some enumeration functionality added


# To compile:

(linux)

$ env GOOS=target-OS GOARCH=target-architecture go build -o gshell.exe main.go

(windows)

set GOOS=Windows

set GOARCH=target-arch

go build -o gshell.exe main.go


# supported os/arch:

windows/386

windows/amd64


# TODO:
add functionality for arguments to cd, handle spaces in path

add functionality for arguments to ls, handle spaces in path

add cross-platform support and implement enum for linux

clean up enum output

add option for output to file
