# go-shell
Simple shell in Go language
## Table of Contents
* [Introduction](#introductiono)
* [Purpose](#purpose)
* [Usage and Examples](#usage-and-examples)
* [Future Improvements](#future-improvements)
***
## Introduction
This project is a short, simple shell program for UNIX-like systems. The program parses a command and executes
it using standard packages. 
***
## Purpose
This project is similar to the shell program I previously wrote in C. I wanted to write a shell in Go language
as an exercise to practice using the language and use these packages and functions. However, I still intend
for this shell to work smoothly and be comparable to other shells.
***
Usage and Examples
This shell is used the same way as any other standard shell.
````
>> ls
README.md	main.go
>> ls -al
total 16
drwxr-xr-x   5 andrewkraynak  staff  160 Dec 10 17:01 .
drwxr-xr-x   5 andrewkraynak  staff  160 Dec 11 12:52 ..
drwxr-xr-x  13 andrewkraynak  staff  416 Dec 12 12:45 .git
-rw-r--r--   1 andrewkraynak  staff  954 Dec 12 12:58 README.md
-rw-r--r--   1 andrewkraynak  staff  591 Dec 12 12:43 main.go
````
Enter commands and the output is displayed on your terminal.
````
>> exit
````
Enter "exit" to end the shell.
***
## Future Improvements
Some improvements I would like to make to this program is to create custom implementations of some commands,
like `<ls>` and `<set>`. I would also make a few more changes to get every command to work, for example the `<cd>`
needs a little more work to run correctly.
