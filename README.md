# Socer

Socer implements multiple generic data structures for simple use.
This library is designed to create a simple API, following the Go
simlicity philosophy that I love so much.

# Usage

## List

```go
var l List[int]

l.PushBack(32)
l.PushBack(64)
l.PushBack(128)

fmt.Println(l)
// [32 64 128]
```

## Stack

```go
var s Stack[int]

s.Push(32)
s.Push(64)
s.Push(128)

fmt.Println(s)
// [32 64 128]

s.Pop()

fmt.Println(s)
// [32 64]
```

## Queue

```go
var q Queue[int]

q.Enqueue(32)
q.Enqueue(64)
q.Enqueue(128)

fmt.Println(q)
// [128 64 32]

s.Dequeue()

fmt.Println(q)
// [64 32]
```

## Deque

```go
var d Deque[int]

d.PushBack(32)
d.PushBack(64)
d.PushFront(128)

fmt.Println(d)
// [64 32 128]

s.PopBack()

fmt.Println(d)
// [32 128]
```

## Tree

```go
var b Tree[int]

b.Insert(32)
b.Insert(64)
b.Insert(128)

fmt.Println(b.Root().Value)
// 32
fmt.Println(b.Root().Left.Value)
// 64
fmt.Println(b.Root().Right.Value)
// 128
```

# Why?

The only advantage of using Socer over something like
[emirpasic/gods](https://github.com/emirpasic/gods) is the simplicity
of the package. Socer don't focus on creating a big and useful lib,
rather, it focus on creating multiple simple and flexible APIs to
manipulate common data structures.

I took inspiration from [nothings/stb](https://github.com/nothings/stb)
which is a stb library for C. Being a stb library means it consists
of a single header file (.h), making it easy to just copy and paste
the library on your project and start using it. You can basically
do the same thing with these APIs.

I also took inspiration from the very Go stdlib, as I said, I really
love Go's simplicity.

Still, it's essentially just a learning project for me.

# Collaboration and feedback

Feel free to open issues/pull requests for the repository,
I would be grateful to see what I can improve in the project.

# License

Copyright 2026 Bruno Guimarães Souza. All rights reserved.
Use of this source code is governed by a MIT license that
can be find in the LICENSE file.
