# srt2fcpxml

srt 字幕文件转为final cut pro 字幕文件(fcpxml)

本软件使用 final cut pro X 10.4.6 版本的 fcpxml 文件作为模版开发，如果有问题请升级到对应版本


## 编译

在项目目录下执行`make`命令后在`build`目录下生成`srt2fcpxml`执行文件。

## 下载

不想编译的用户可以直接下载[执行文件](https://github.com/GanymedeNil/srt2fcpxml/releases)。

## 使用
首先需要赋予程序执行权限 `chmod +x ./srt2fcpxml`

```bash
$ ./srt2fcpxml
  -fd int
    	帧率目前只支持整数24、25、30、50、60 (default 25)
  -srt string
    	srt 字幕文件
  -w int
        分辨率宽 (default 1920)
  -h int 
        分辨率高 (default 1080)
  -line string
        每行字幕连接的符号 当传\\n字符时会进行字幕的换行（default - ）
```

执行

```bash
$ ./srt2fcpxml -srt /tmp/test.srt
```

在 srt 文件的目录中会自动生成以srt文件名命名的`fcpxml`文件。
