[![Build Status](https://travis-ci.com/gekalogiros/Doo.svg?branch=master)](https://travis-ci.com/gekalogiros/Doo)

# Doo

Doo aims to be a dead simple and super quick cli tool to manage non-sophisticated task lists on your OSX machine

The idea of crafting this app came after attended the "7 Principles of Productive Software Developers" talk given by Java Rockstart and champion [Sebastian Daschner](https://github.com/sdaschner).   

***IMPORTANT:***

This project is under development

## Motivation

The terminal is the control hub for many dev systems. 

Doo aims to be a good companion to developers and system administrators or just people who want to control stuff from the command line. 

Doo aims to be an effective tool that stores the points you'd like to discuss in your daily stand-ups, light tasks that you have to complete or phone calls that you need to make. It actually uses the filesystem to do it. It is as easy as it gets.

Doo has not been designed to be dead simple. It is supposed to be something quick and dirty for tracking tasks of minor importance. 

There are plenty of amazing and rich featured task management apps out there. All of them are definitely better than Doo. However, if your lack of discipline prevented you from making them an integral part of your daily routine, give Doo a go.

## Usage

- Add Note
```
$ doo add -d "I am a todo task" -dd 10d
```

- Remove Notes for a certain date
```
$ doo rm -dt 17-09-2018 -f
```

## Due Date Formats Allowed (--dd flag)

Assuming today is 2nd of January 2018 and you want to add a todo app 1 year from now, you have the following options when defining the `-dd` (due date) flag:

```
1y
1Y
12m
12M
365d
365D
365
02/01/2019
2/1/2019
02/01/19
2/1/19
02-01-2019
2-1-2019
02-01-19
2-1-19
```

# License

The content of this project itself is licensed under the [Creative Commons Attribution 3.0 license](http://creativecommons.org/licenses/by/3.0/us/deed.en_US), and the underlying source code used to format and display that content is licensed under the [MIT license](http://opensource.org/licenses/mit-license.php).
