# How unbuffered channels work

![](./_media/go-unbuffered-channels.gif)

## Resources

 - [Unbuffered and buffered channels](https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/unbuffered-and-buffered-channels.html)
 - [Go (programming language): What happens when a goroutine blocks?](https://www.quora.com/Go-programming-language-What-happens-when-a-goroutine-blocks) This’s one of the key insights about Go: "A single IO thread isn't necessary since a blocked goroutine does not imply a blocked thread -- in other words, goroutines multiplexed onto the same thread are not blocked."
 - [GopherCon 2017: Kavya Joshi - Understanding Channels](https://www.youtube.com/watch?v=KBZlN0izeiY&feature=youtu.be)
