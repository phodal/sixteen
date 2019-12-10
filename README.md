# Sixteen


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
