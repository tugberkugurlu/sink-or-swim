# sink-or-swim

Me learning Go programming language by implementing concepts from real world problems

## Scenarios

### 00: line-count

Implement a command-line application which takes a path to a file and calculates the number of lines in that file. Also accept optional command-line paratemer to skip counting the empty lines.

**Main learning objectives**: Learn about receiving and validating command-line parameters and reading from the file system.

### 01: Uniq

Implement a function which takes a collection of integers and returns a new collection of integers without duplications.

**Main learning objectives**: Learn about Hashtable and Hashset in Golang.

### 02: Implement Stack w/o Using Collection Data Structures

Implement the Stack data structure w/o using collection data structure by making `push` and `pop` operations perform in O(1) time complexity.

**Main learning objectives**: Learn about data structure implementation within defined constraints.

### 03: Binary-tree Level Order Traversal

See the question [here](https://leetcode.com/problems/binary-tree-level-order-traversal).

**Main learning objectives**: Learn about how you would implement a Tree object in Go. How to implement BFS by using a Queue data structure.

### 04: LRU Cache

Implement Least-recently Used Cache with both get and put operations perfomed in O(1) time complexity. See the question [here](https://leetcode.com/problems/lru-cache/).

**Main learning objectives**: Learn about using doubly linked list in Go.

## Ideas

 - JSON serialization/deserialization
 - HTTP GET request handling
 - HTTP POST request handling
 - CSRF protection
 - Logging in HTTP requests
 - LRU cache through HTTP transport layer
 - LRU cache through HTTP transport layer, with Redis as the data storage system

## Resources

 - [Go by Example](https://gobyexample.com/)
 - [GopherCon 2018: Filippo Valsorda- Asynchronous Networking Patterns](https://www.youtube.com/watch?v=afSiVelXDTQ)
