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
        Map(strconv.Itoa).
        Tap(fmt.Println).
        Val().([]string)
```

### Available functions:

* `Map`
* `Filter`
* `Reduce`
* `ForEach`
* `Tap`
* `Any`
* `Every`
* `Contains`
* `Find`

### Chaining 

* `Wrap`
* `Val`
