# 2019.03.25-2019.03.31

## Algorithm

见`algo`目录下`two-sum`

## Review

1. [下个项目我应该用Golang还是Node.js?](https://medium.com/@faith.chikwekwe/should-i-use-golang-or-node-js-for-my-next-app-e15d9c71358e)

   先介绍了`Node.js`是非常适合后端开发入手学习的，还鄙视一下`PHP`。然后写了一些`Golang`的一些优点：高性能、高并发等，还有受谷歌代码文档的影响使`Golang`的代码格式化特别优雅

2. [如何使用日志节省调试的时间](https://medium.freecodecamp.org/how-to-save-hours-of-debugging-with-logs-6989cc533370)

   好的日志可以在需要的时候帮助我们，他对于我们处理错误的能力是非常重要的。

2. 第三

## Tip

配置`supervisor`
使用`supervisor`管理`flask`的`uwsgi`, 刚开始怎么都管理不了，启动报错
`FALTAL error ocurred when uwsgi works with superviso`。
最后发现是由于`uwsgi.ini`配置文件中`daemonize = /home/www/logs/uwsgi.log` 日志文件 可能启用了后端进程，把他修改为`logto = /home/www/logs/uwsgi.log`就`ok`了。
还有一点，之前一直一位`uwsgi`不能平滑重启，这次使用中知道了 `uwsgi --reload /var/uwsgi/uwsgi.pid`

## Share

python 获取文件字符编码类型

```
import chardet
# 获取文件编码类型
def get_encoding(file):
    # 二进制方式读取，获取字节数据，检测类型
    with open(file, 'rb') as f:
        return chardet.detect(f.read())['encoding']

file_name = 'data.sav'
encoding = get_encoding(file_name)
print(encoding)
```