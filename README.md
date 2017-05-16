# semmy: Semantic Versions from the Command Line

To get the project into your go workspace:
```
go get github.com/technodabble/semmy
```

To install it:
```
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
