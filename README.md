# semmy: Semantic Versions from the Command Line

## Get the software

Here you have two options. If you just want to run it:
```
$> go get github.com/technodabble/semmy
```

But, if you want to develop it, clone it from github:
```
$> cd $GOPATH
$> mkdir -p src/github.com/technodabble
$> cd src/github.com/technodabble
$> git clone git@github.com:technodabble/semmy.git
```

To install it:
```
go get -u github.com/kardianos/govendor
govendor sync github.com/technodabble/semmy
go install github.com/technodabble/semmy
```

To run it:
```
semmy check range [version]
```
`range` is as [documented here](https://github.com/blang/semver#ranges).

`version` is a [Semantic Version](https://semver.org) If it's omitted, it will
be read from the first line of `STDIN`

### Examples
```
$> semmy check '>1.0.10 <3.0.4' 2.4.5-alpha.13+10211979
2.4.5-alpha.13+10211979 is in the range: >1.0.10 <3.0.4

$> echo $?
0

$> echo 3.5.6 | semmy check '>1.0.10 <3.0.4'
3.5.6 is NOT in the range: >1.0.10 <3.0.4

$> echo $?
1
```
