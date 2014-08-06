gedd
====

gedd searches for matching files in a fedd index.

####Usage
```
gedd "pattern"
```
The pattern is not case sensitive. The shell pattern matching is used featuring the following symbols: ```
*?[^-]
```.

####Installation

You can either use
```
go get https://github.com/hirsch/gedd
```
or download it and navigate to the directory followed by
```
go build
```
Copy the compiled file to /usr/bin/ (Linux) or C:\Windows\System32\ (Windows)
to be able to execute it in every directory.
Alternatively you may set a PATH variable both on Windows and Linux.
