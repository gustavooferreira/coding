# README

The problems are stored in this structure so that there is no package name collision.

Create a folder for each problem.

Open vim in this folder. This is required because the go module is defined from this root folder.

To run tests for a particular package do the following:

```
 go test -v [packages]
```

For example, to test the very first Leet Code exercise:

```
 go test -v ./1/solutions
```

As the golang documentation says (`go help packages`):

> Usually, [packages] is a list of import paths.
>
> An import path that is a rooted path or that begins with
> a . or .. element is interpreted as a file system path and
> denotes the package in that directory.
