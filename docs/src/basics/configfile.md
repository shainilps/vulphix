---
title: Configuration file
---

As of right now config file of vulpix is very minmal. you don't get lot of features like you get from feature rich ssg like hugo or even the astro starlight. This ssg is simple, minimal and was designed for me

Unlike other SSGs I am using YAML format as the config type, because It is more elegant to look at.

config file:

```yaml
title: Vulpix
domain: doc.vulpix.com # (optional)
description: this is the docs of your vulpix
handle: https://x.com/codeshaine
source: doc/src
build: dist
sidebar:
  - title: Basics
    pages:
      - - Getting started
        - /getting_started/installation
  - title: Links
    pages:
      - - Github
        - https://github.com/codeshaine
      - - Twitter
        - https://x.com/code_shaine
```
The config file is more self  exploratory.
- title, domain, description, handle are used for the meta info of the site.
- note that title used as main logo
- source and build direcory is the path of your source and build that you want to generate
- sidebar is the main feature available here like other SSGs. This follow a list of object so you can have flexibility
these are all the main features and characterstics
