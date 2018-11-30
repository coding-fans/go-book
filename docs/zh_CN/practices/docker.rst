.. Docker化部署
    FileName:   docker.rst
    Author:     Fasion Chan
    Created:    2018-11-29 19:50:57
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
        本文介绍如何制作一个紧凑的Docker镜像用于部署Go应用，要点包括：
        以scratch为母版、静态链接、添加时区信息、添加CA证书。
    :keywords: golang, docker, zoneinfo, ssl ca cert, GOOS, CGO_ENABLED

============
Docker化部署
============

*Go* 语言应用部署 **不需要依赖** ，非常简便，这是一个不小的优势。

*Go* 语言官方镜像非常大，超过 *500MB* 。
镜像之所以如此庞大是因为它包含了构建 *Go* 程序所需的全部 **工具链** 。
然而，运行编译好的(静态)二进制程序并不需要这些工具。

本文介绍如何制作一个紧凑的 *Docker* 镜像用于部署 *Go* 应用，大小控制在 *10MB* 以内。

.. note::

    本文实验所有操作均在 *OSX* 下进行，在其他平台进行也是类似的。

实验以这个简单程序为例：

.. literalinclude:: /_src/practices/docker/demo.go
    :name: practices/docker/demo.go
    :language: go
    :lines: 13-
    :linenos:

程序先输出当前本地时间以及东京时间；接着请求百度首页并判断是否成功。

先对其进行编译并查看二进制程序大小：

.. code-block:: shell-session

    $ go build demo.go
    $ ls -lh demo
    -rwxr-xr-x 1 fasion staff 5.8M Nov 29 18:36 demo

注意到，二进制程序大小仅为 *5.8MB* 。

接下来我们准备一个 `Dockerfile` 来构建镜像：

.. literalinclude:: /_src/practices/docker/Dockerfile-without-zoneinfo
    :name: practices/docker/Dockerfile-without-zoneinfo
    :language: Dockerfile
    :lines: 1-
    :linenos:

执行镜像构建命令：

.. code-block:: shell-session

    $ docker build -t demo .
    Sending build context to Docker daemon  6.299MB
    Step 1/4 : FROM scratch
     --->
    Step 2/4 : ENV TZ=Asia/Shanghai
     ---> Using cache
     ---> f6653a3d26c7
    Step 3/4 : ADD demo /
     ---> 437b92b7460a
    Step 4/4 : CMD ["/demo"]
     ---> Running in 0bc0aef56fab
    Removing intermediate container 0bc0aef56fab
     ---> e8ff5745453e
    Successfully built e8ff5745453e
    Successfully tagged demo:latest

接着，尝试执行该镜像：

.. code-block:: shell-session

    $ docker run demo
    standard_init_linux.go:190: exec user process caused "exec format error"

执行过程中不幸出错，错误为 *exec format error* ，即 **可执行程序格式错误** 。
这个错误是由于我们在 `Docker` 容器内执行 `OSX` 二进制程序造成的：

.. code-block:: shell-session

    $ file demo
    demo: Mach-O 64-bit executable x86_64

.. note::

    *Docker* 容器技术是基于 `Linux` 内核的，因此只能执行 `Linux` 格式二进制程序。

接着需要通过 **交叉编译** 来构建 *Linux* 二进制程序：

.. code-block:: shell-session

    $ CGO_ENABLED=0 GOOS=linux go build demo.go
    $ file demo
    demo: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped

交叉编译生成的程序为 `ELF` 格式，正是 `Linux` 二进制程序格式！

.. note::

    注意到交叉编译生成的程序是 **静态链接** 的，无需依赖任何动态库文件。
    由于 `scratch` 母镜像不包含必要的动态库文件，运行动态链接的程序将报错。

有了 `Linux` 二进制后，便可重新构建镜像并执行：

.. code-block:: shell-session

    $ docker run demo
    panic: time: missing Location in call to Time.In

    goroutine 1 [running]:
    time.Time.In(...)
            /usr/local/go/src/time/time.go:1086
    main.main()
            /Users/fasion/workspace/works/go-book/docs/zh_CN/_src/practices/docker/demo.go:24 +0x311

镜像成功启动，但是遇到另一个错误。
原因在于 `scratch` 镜像并不包含任何 **时区信息** ，我们需要从本地系统中复制一份。

由于 `scratch` 镜像几乎不包含任何东西，甚至没有 `mkdir` 命令。
因此，我们需要对时区信息进行打包，然后再通过 *ADD* 指令进行添加，以此绕过目录创建：

.. code-block:: shell-session

    $ tar -chzf zoneinfo.tar.gz /usr/share/zoneinfo
    tar: Removing leading '/' from member names

修改后的 `Dockerfile` 如下：

.. literalinclude:: /_src/practices/docker/Dockerfile-without-cacert
    :name: practices/docker/Dockerfile-without-cacert
    :language: Dockerfile
    :lines: 1-
    :linenos:

再次构建并执行：

.. code-block:: shell-session

    $ docker run demo
    Local time: 2018-11-29 18:47:42.1671632 +0800 CST m=+0.002501501
    Tokyo time: 2018-11-29 19:47:42.1671632 +0900 JST
    Baidu website is DOWN
    Err: Get https://www.baidu.com/: x509: certificate signed by unknown authority

这时，时区信息问题已经解决了，但是发起 `HTTPS` 请求还存在问题，
原因是 `scratch` 镜像不包含任何 `SSL` `CA` 证书。

接下来，从 `https://curl.haxx.se/docs/caextract.html <https://curl.haxx.se/docs/caextract.html>`_ 下载 `CA` 证书：

.. code-block:: shell-session

    $ curl -o cacert.pem https://curl.haxx.se/ca/cacert.pem

更新 `Dockerfile` 将证书添加到镜像：

.. literalinclude:: /_src/practices/docker/Dockerfile
    :name: practices/docker/Dockerfile
    :language: Dockerfile
    :lines: 1-
    :linenos:

再次构建并执行：

.. code-block:: shell-session

    $ docker run demo
    Local time: 2018-11-29 19:01:19.5481835 +0800 CST m=+0.001815301
    Tokyo time: 2018-11-29 20:01:19.5481835 +0900 JST
    Baidu website is UP

至此，大功告成！

最后，我们分别检查二进制程序以及镜像的大小：

.. code-block:: shell-session

    $ ls -lh demo
    -rwxr-xr-x 1 fasion staff 5.8M Nov 29 18:47 demo
    $ docker images
    REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
    demo                latest              a9ae4b8e7f1e        8 minutes ago       7.06MB

正如你所见，二进制程序大小是 *5.8MB* ，而 `Docker` 镜像只有 *7MB* 左右！

总结
====

构建紧凑的 `Go` 应用部署镜像，需要注意以下要点：

- 以 `scratch` 为母版；
- 静态链接；
- 添加时区信息；
- 添加 `CA` 证书；

这个经验适用于任何二进制程序，不局限于 `Go` ，其他语言如 `C` 也是类似的。

下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

参考文献
========

#. `Create a small Docker image for a GoLang binary <https://sebest.github.io/post/create-a-small-docker-image-for-a-golang-binary/>`_

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

