# vulphix

a markdown SSG generator for documentation site. Simple, Minimal with yaml config.

## Usage

```bash
vulphix init #TODO
vulphix build
vulphix preview
```

vulphix.config.yaml

```yaml
# customization
title: site-name
domain: site-domain.com # (optional)
description: your doc site descriptio
handle: https://x.com/your-username #any social media handle link
source: src
build: dist
sidebar:
  - title: Basics
    pages:
      - - Getting started
        - /basic/setup #(setup == setup.md)
  - title: Links
    pages:
      - - Github
        - https://github.com/username
```

Note:

- index.md which reside in your source root will be your home page
- put favicon.ico in root of source for chooseing it as favicon of your site

## Feature

- [x] fixed template (for now)
- [ ] syntax highlight support for (Go, JavaScript, TypeScript)
- [x] build preview
- [x] left side bar navigation for site
- [ ] go releaser
- [ ] vulphix init
- [ ] right side bar navigation for page
- [ ] deploy to some cloud
- [ ] ci/cd generator for project (Future plan)
- [ ] Hot reload (Future plan)

## Contributing

open for contribution. you can fork repo, create a issue and make a PR. (that's all)

<center>happy coding<3 <center/>
