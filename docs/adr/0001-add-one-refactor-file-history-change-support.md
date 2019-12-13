# 1. add one refactor file history change support

Date: 2019-12-13

## Status

2019-12-13 proposed

## Context

Context here...

## Decision

similar to :

```
git log --follow -- main.go
```

for one file changes:

```
git log --follow --pretty='format:[%h] %aN %ad %s' --date=iso  -- README.md
```

## Consequences

Consequences here...
