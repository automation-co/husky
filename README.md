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

If you have made any other changes in the hooks you can appply them by using `$ husky install`

---

## Get Familiar with Git Hooks

Learn more about git hooks from these useful resources:
- https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks
- https://www.atlassian.com/git/tutorials/git-hooks
- https://medium.com/@f3igao/get-started-with-git-hooks-5a489725c639

---

### Other Alternatives

If you feel husky does not fulfill your needs you can also check out:
- https://github.com/typicode/husky
- https://pre-commit.com/

---

<div align="center">

Developed by [@automation-co](https://github.com/automation-co)

</div>
