# 2019.03.25-2019.03.31

## Algorithm

见`algo`目录下`add-two-numbers`
这个最开始写理解错题目意思了，导致走了冤枉路，两数相加的时候，链表本来就是逆向的，正好按照顺序两两相加就好了，我理解成需要先倒序然后再加再倒序。下次注意理解题意。

## Review

1. [下个项目我应该用Golang还是Node.js?](https://medium.com/@faith.chikwekwe/should-i-use-golang-or-node-js-for-my-next-app-e15d9c71358e)

   先介绍了`Node.js`是非常适合后端开发入手学习的，还鄙视一下`PHP`。然后写了一些`Golang`的一些优点：高性能、高并发等，还有受谷歌代码文档的影响使`Golang`的代码格式化特别优雅

2. [如何使用日志节省调试的时间](https://medium.freecodecamp.org/how-to-save-hours-of-debugging-with-logs-6989cc533370)

   好的日志可以在需要的时候帮助我们，他对于我们处理错误的能力是非常重要的。

## Tip

由于项目需要，使用`multiprocessing` 多进程并行计算

## Share

在使用`multiprocessing`的进程池 `Pool` 的时候，测试确实产生了多个进程，但是多个核还是没有利用上，目前还比较困惑，后面再深入探究。