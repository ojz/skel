# skel

Copies an entire tree and does a search and replace in directorynames, filenames and file contents. Useful when you need to create multiple projects from a single common base.

```bash
Usage of skel:
  -dest="": Where to save the generated files.
  -name="": The name of the new project (defaults to the basename of the destination).
  -nonce="": The word which should be replaced (defaults to the basename of the source).
  -source="": Where to find the template files.
```

## Example setup

```
mkdir -p ~/.skel/acme-go-angularjs
# create a seed project in the directory above 
alias gong="skel -nonce seedproject -source ~/.skel/acme-go-angularjs -dest"
gong ~/Projects/todoApp -name todo
```

## Todo

* allow source to be a file
