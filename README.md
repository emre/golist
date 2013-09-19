golist [![Build Status](https://travis-ci.org/emre/golist.png)](https://travis-ci.org/emre/golist) [![Clone in Koding](http://kbutton.org/clone.png?v99)](http://kbutton.org/emre/golist)
======

python's list type implementation on go.

#### installation
```bash
$ go get github.com/emre/golist
```

#### Initialization of a new List 
```go
my_list := golist.New()
my_list := golist.New("Galatasaray", "Real Madrid")
```

#### Appending items
```go
my_list.Append("Drogba")
my_list.Append("Bruma")
```


#### Extending list with another list
```go
other_list := golist.New("Ronaldo", "Bale")
my_list.Extend(other_list)
```

#### Get an index for the element
```go
index, err := my_list.Index("Drogba")
if err != nil {
	fmt.Println(error)
}
fmt.Println(index)
```
#### Counting items
```go
my_list.Append("Drogba", "Drogba")
drogba_count := my_list.Count("Drogba")
fmt.Println(drogba_count)
```

#### Deleting items by value/index
```go
// by value
my_list.Remove("Drogba")

// by index
my_list.Delete(0)
```

#### Pop items

```go
// pops the item which of index zero.
value, err := my_list.Pop(0)

// pops latest item
value, err := my_list.Pop()
```

#### Reversing a List
```go
goals_list := golist.New("18:Drogba", "56:Bale", "90+2:Sabri Sarioglu")
goals_list.Reverse()
```

#### Get the size of list
```go
list_size := goals_list.Len()
```

#### Further documentation
- <a href="http://godoc.org/github.com/emre/golist">godoc golist</a>
- <a href="https://github.com/emre/golist/blob/master/list_test.go">usage examples</a>.

