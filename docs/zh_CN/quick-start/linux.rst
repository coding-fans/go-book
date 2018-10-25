.. Linux
    FileName:   linux.rst
    Author:     Fasion Chan
    Created:    2018-08-16 14:53:17
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

=====
Linux
=====

二进制包
========

`Linux` 下，推荐通过二进制包部署 `Go` 语言开发环境。

首先，从 `Go` 官网 `下载页面 <https://golang.org/dl/>`_ 下载 `Go` 语言二进制包。

当前最新版本为： `go1.10.3.linux-amd64.tar.gz <https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz>`_ 。

.. code-block:: shell-session

    $ ls go1.10.3.linux-amd64.tar.gz
    go1.10.3.linux-amd64.tar.gz

下载完毕后，将压缩包解压到家目录的 `opt` 目录中：

.. code-block:: shell-session

    $ mkdir -p ~/opt
    $ tar -xf go1.10.3.linux-amd64.tar.gz -C ~/opt

解压完毕后，我们可以看到这些目录和文件：

.. code-block:: shell-session

    $ ls ~/opt/go
    api      bin   CONTRIBUTING.md  doc          lib      misc     pkg        robots.txt  test
    AUTHORS  blog  CONTRIBUTORS     favicon.ico  LICENSE  PATENTS  README.md  src         VERSION

其中， `Go` 语言开发工具就位于 ``bin`` 目录下。
尝试运行 `go` 命令：

.. code-block:: shell-session

    $ ~/opt/go/bin/go version
    go version go1.10.3 linux/amd64

至此， `Go` 已经安装完毕！

等等！ ``~/opt/go/bin/go`` 这是什么鬼？

哈哈，可以将 ``~/opt/go/bin`` 加到 ``PATH`` 环境变量中：

.. code-block:: shell-session

    $ export PATH=$PATH:~/opt/go/bin
    $ go
    go version go1.10.3 linux/amd64

另外，建议将 ``export PATH=$PATH:~/opt/go/bin`` 这句写到 `shell` 配置，一劳永逸！

.. comments
    comment something out below

    .. meta::
        :description lang=zh:
        :keywords:


下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst
