# 2019.03.18-2019.03.24

## Algorithm
见`algo`目录下`two-sum`

## Review
![解决问题的框架：Understand, Design, Build](https://lob.com/blog/understand-design-build-a-framework-for-problem-solving?utm_source=wanqu.co&utm_campaign=Wanqu+Daily&utm_medium=website)
我们平时学到的都是怎么去Build，工程师培养培养解决问题的能力主要还是理解和设计，最后才是简单的写。工作中，有的同事别人把需求的理解的设计方案想好之后，他就等着写代码，这样就很难提高他实际解决问题的能力！

## Tip

配置`supervisor`
使用`supervisor`管理`flask`的`uwsgi`, 刚开始怎么都管理不了，启动报错
`FALTAL error ocurred when uwsgi works with superviso`。最后发现是由于`uwsgi.ini`配置文件中`daemonize = /home/www/logs/uwsgi.log` 日志文件 可能启用了后端进程，把他修改为`logto = /home/www/logs/uwsgi.log`就`ok`了。
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
