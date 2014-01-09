# Higher

> Higher order functions in Go (really unsafe)
> You probably shouldn't use this

``` go
s1 := []int{1, 2, 3, 4, 5}
s2 := higher.Filter(s1, func(x int) bool { return x > 2 })
s3 := higher.Map(s, strconv.Itoa).([]string)
```

You can chain them too.

``` go
s := higher.Wrap([]int{1, 2, 3, 4, 5}).
        Filter(s1, func(x int) bool { return x > 2 }).
        Map(s, strconv.Itoa).
        Val().([]string)
```

