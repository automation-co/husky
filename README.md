<img src="https://user-images.githubusercontent.com/64161383/155763268-e09d9613-a53f-4ec7-a943-aab93ef2ffa6.png" width="150px" alt="logo"  align="right" />

<div align="left">


# Husky

[![Build Status](https://github.com/automation-co/husky/workflows/Go/badge.svg?branch=main)](https://github.com/automation-co/husky/actions?query=branch%3Amain)
[![Release](https://img.shields.io/github/release/automation-co/husky.svg)](https://github.com/automation-co/husky/releases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/automation-co/husky)
[![Go Report Card](https://goreportcard.com/badge/github.com/automation-co/husky)](https://goreportcard.com/report/github.com/automation-co/husky)
![GitHub](https://img.shields.io/github/license/automation-co/husky)
![GitHub issues](https://img.shields.io/github/issues/automation-co/husky)

 </div>

<!-- --- -->

**Make githooks easy!**

Inspired from the [husky.js](https://github.com/typicode/husky)


---
<a href="https://www.producthunt.com/posts/husky-4?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-husky&#0045;4" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=346044&theme=light" align="right" alt="Husky - Git&#0032;hooks&#0032;made&#0032;easy&#0032;on&#0032;go | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>


## Docs

### Installation

```
go install github.com/automation-co/husky@latest
```

### Getting Started

You can initialise husky by `$ husky init`

> Make sure you have git initialised

This will make the .husky folder with the hooks folder and a sample pre-commit hook

You can add hooks using

```bash
$ husky add <hook> "
  <your commands for that hook>
"
```

### Example

```bash
$ husky add pre-commit "
  go build -v ./... 
  go test -v ./...
"
```

If you have made any other changes in the hooks you can apply them by using `$ husky install`

---

## Blogs and Resources
- [ Get Started with Husky for go ](https://dev.to/devnull03/get-started-with-husky-for-go-31pa)
- [ Git Hooks for your Golang project ](https://dev.to/aarushgoyal/git-hooks-for-your-golang-project-1168)

---

## Get Familiar with Git Hooks

Learn more about git hooks from these useful resources:
- [ Customizing Git - Git Hooks ](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [ Atlassian Blog on Git Hooks ](https://www.atlassian.com/git/tutorials/git-hooks)
- [ Fei's Blog | Get Started with Git Hooks ](https://medium.com/@f3igao/get-started-with-git-hooks-5a489725c639)

---

### Other Alternatives

If you feel husky does not fulfill your needs you can also check out:
- https://github.com/typicode/husky
- https://pre-commit.com/

---

<div align="center">

Developed by [@automation-co](https://github.com/automation-co)

</div>
