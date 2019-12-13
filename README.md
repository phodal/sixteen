# Sixteen


## Usage

Usage Mode:

 - Interactive Mode
 - Cli Mode

Install

```
go run main.go
```

examples

```
➜  sixteen git:(master) ✗ go run main.go     
Use the arrow keys to navigate: ↓ ↑ → ← 
? Refactoring: 
  ▸ list
    step
    switch
    delete
↓   commit
```

 - commit: use refactoring commit
 - create: create new refactoring task 
 - step: show step of tasks
 - show: show tasks info


### Show Example

```
go run main.go show
```

results:

```
add blablbla blabla 
  2019-12-12 13:16:09 refactoring: make date be long-3NJoo9aWR
phodal-refactoring itfas
  2019-12-11 20:28:28 refactoring: use cli api-4TPJClaZR
blabla 
  2019-12-11 20:32:15 refactoring: remove unused commit-ITolT9aWg
a show case
  2019-12-11 20:21:18 refactoring: remove unsued validate && add commit select tasks-cU6xor-Wg
支持中文符号
  2019-12-12 13:06:38 refactoring: add basic show foramt-iRych9aZg
  2019-12-11 20:23:00 refactoring: extract command method-iRych9aZg
```

## Document

Goals:

 - 旧的不变
 - 新的创建
 - 一步切换
 - 旧的再见
 

```
[refactoring] create new --- targetA  --  具体做了啥-1
[refactoring] create new --- targetA  --  具体做了啥-2
[refactoring] delete old  --- targetA  --  具体做了啥
[refactoring] switch  --- targetA  --  具体做了啥
```

Design:

steps to markdown file

```
sixteen list
sixteen step {id}
sixteen switch {id}
sixteen new {id}
sixteen old {id}
sixteen delete {id}
```

commit log

```
refactoring: blablablablabla [123-12]
```

todo-doing-done

new-switch-old-delete

config file:

```
prefix: 'refactoring'
log-pattern: '{message} {id}' or '[{id}] {message}'  
```

License
---

[![Phodal's Idea](http://brand.phodal.com/shields/idea-small.svg)](http://ideas.phodal.com/)

@ 2019 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MIT license. See `LICENSE` in this directory.
