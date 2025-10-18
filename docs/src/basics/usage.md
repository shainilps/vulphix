---
title: Simple Usage
---
you can get stared with vulipx right after you install, vulphix support two commands for now

build: this command will build the static site for your markdowns
```bash
vulphix build
```

preview: this command will serve the build
```bash
vulphix preview
```

well inorder to build you need config file,
example config  file
```yaml
title: vulphix
domain: doc.vulphix.com # (optional)
description: this is the docs of your vulphix
handle: https://x.com/codeshaine
source: doc/src
build: dist
sidebar:
  - title: Basics
    pages:
      - - Getting started
        - /getting_started/installation
```
