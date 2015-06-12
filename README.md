# Higher [![Build Status](https://travis-ci.org/icholy/higher.png?branch=master)](https://travis-ci.org/icholy/higher)

> Higher order functions in Go (really unsafe).
> You probably shouldn't use this.

[godoc](http://godoc.org/github.com/icholy/higher)

### Examples

``` go
s1 := []int{1, 2, 3, 4, 5}
s2 := higher.Filter(s1, func(x int) bool { return x > 2 })
s3 := higher.Map(s, strconv.Itoa).([]string)
```

### You can chain them too.

``` go
s := higher.Wrap([]int{1, 2, 3, 4, 5}).
        Filter(func(x int) bool { return x > 2 }).
        PMap(strconv.Itoa).
        Tap(fmt.Println).
        Val().([]string)
```

### Available functions:

* `Map`
* `PMap`
* `Filter`
* `PFilter`
* `Reduce` reduce can't be parallel
* `ForEach`
* `PForEach`
* `Tap`
* `PTap`
* `Any`
* `PAny`
* `Every`
* `PEvery`
* `Contains`
* `PContains`
* `Find`
* `PFind`

### Chaining 

* `Wrap`
* `Val`

![](higher.jpeg)
