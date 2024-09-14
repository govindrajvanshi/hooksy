


# Hooksy


**Make githooks easy!**

---

## Docs

### Installation

```
go install github.com/govindrajvanshi/hooksy@latest
```

### Getting Started

You can initialise hooksy by `$ hooksy init`

> Make sure you have git initialised

This will make the .hooksy folder with the hooks folder and a sample pre-commit hook

You can add hooks using

```bash
$ hooksy add <hook> "
  <your commands for that hook>
"
```

### Example

```bash
$ hooksy add pre-commit "
  go build -v ./... 
  go test -v ./...
"
```

If you have made any other changes in the hooks you can apply them by using `$ hooksy install`

---



## Get Familiar with Git Hooks

Learn more about git hooks from these useful resources:
- [ Customizing Git - Git Hooks ](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [ Atlassian Blog on Git Hooks ](https://www.atlassian.com/git/tutorials/git-hooks)
- [ Fei's Blog | Get Started with Git Hooks ](https://medium.com/@f3igao/get-started-with-git-hooks-5a489725c639)

---

### Other Alternatives

If you feel hooksy does not fulfill your needs you can also check out:

- https://pre-commit.com/

---

<div align="center">

Developed by [@govindrajvanshi](https://github.com/govindrajvanshi)

</div>
