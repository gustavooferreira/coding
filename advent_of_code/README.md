# README

The problems are stored in this structure so that there is no package name collision.

Create a folder for each problem.

You can copy a template to the folder for an exercise for day X and part Y, by doing:

```sh
mkdir -p 2023/dayX/partY
cp -R template_dayX_partY/ 2023/dayX/partY
```

Open vim in this folder. This is required because the go module is defined from this root folder.

To run tests for a particular package do the following:

```sh
go test -v [packages]
```

For example, to test the code, put yourself inside the exercise folder and run:

```sh
cd 2023/day1/part1
go test -v ./solutions
```

As the golang documentation says (`go help packages`):

```
Usually, [packages] is a list of import paths.
An import path that is a rooted path or that begins with
a . or .. element is interpreted as a file system path and
denotes the package in that directory.
```
