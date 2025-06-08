# vulpix
a  markdown SSG generator for documentation site. Simple, Minimal with yaml config.

## Usage
```bash
vulpix init #TODO
vulpix build
vulpix preview
```
```yaml
# customization
title: site-name
domain: site-domain.com # (optional)
description: your doc site description
handle: https://x.com/your-username #any social media handle link
source: src
build: dist
sidebar:
  - title: Basics
    pages:
      - - Getting started
        -  /basic/setup  #(setup == setup.md)
  - title: Links
    pages:
      - - Github
        - https://github.com/username
```

## Features
- [ ] fixed template (for now)
- [ ] syntax highlight support for (Go, JavaScript, TypeScript)
- [ ] build preview
- [ ] vulpix init
- [ ]  deploy to some cloud
- [ ]  ci/cd generator for project (Future plan)
- [ ] Hot reload (Future plan)

## Contributing
open for contribution. you can fork repo, create a issue and  make a PR. (that's all)

<center>happy coding<3</center>
