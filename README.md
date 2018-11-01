[![Build Status](https://travis-ci.com/gekalogiros/Doo.svg?branch=master)](https://travis-ci.com/gekalogiros/Doo)
[![Release Number](https://img.shields.io/github/release/gekalogiros/Doo.svg)](https://travis-ci.com/gekalogiros/Doo)
 
# Doo

Doo aims to be a dead simple and super quick cli tool to manage non-sophisticated task lists on your OSX machine

The idea of crafting this app came after attended the "7 Principles of Productive Software Developers" talk given by Java Rockstart and champion [Sebastian Daschner](https://github.com/sdaschner).   

***IMPORTANT:***

This project is under development

## Usage

- Adding a Task
```
$ doo add "I am a todo task added in today's task list"
$ doo add -t "I am a todo task. Due date is in 10 days" -d 10d
$ doo add -t "I am a todo task. Due date is in 1 month" -d 1m
$ doo add -t "I am a todo task. Due date is tomorrow" -d 1
```

- Lookup Task List
```
$ doo ls
$ doo ls today
$ doo ls -d 1m
```

- Moving a Task between lists
```
$ doo mv -f today -t tomorrow -id 53f4
```

- Remove Task List
```
$ doo rm -d 0
$ doo rm -d today
$ doo rm -d -1
$ doo rm -d -1m
$ doo rm -past
```

## Date Formats Allowed (-d flag or parameter)

Assuming today is 2nd of January 2018 and you want to add a todo app 1 year from now, you have the following options when defining the `-d` (date) flag:

```
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

For operations requiring dates in the past you can use the minus (-) symbol at the start (mainly for removing task lists or moving tasks from past dates to today or future)

```
-5
-5d
-5m
-5y
```

Also you can use the following keywords

```
today
tomorrow
yesterday
```

# License

The content of this project itself is licensed under the [Creative Commons Attribution 3.0 license](http://creativecommons.org/licenses/by/3.0/us/deed.en_US), and the underlying source code used to format and display that content is licensed under the [MIT license](http://opensource.org/licenses/mit-license.php).
