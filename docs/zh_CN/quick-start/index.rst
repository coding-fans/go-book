.. 快速开始
    FileName:   index.rst
    Author:     Fasion Chan
    Created:    2018-03-05 22:08:43
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

========
快速开始
========

本文以 `OSX` 为例，介绍如何快速搭建环境进行 `Go` 语言开发。

文章大部分内容对其他操作系统也是适用的，但安装部分因系统而异。
正文以 `OSX` 为例进行演示，其他操作系统可以参考以下链接：

.. toctree::
    :titlesonly:

    Linux <linux>

安装
====

开始开发之前，需要先搭建好开发环境。

首先，需要安装 `Go` 开发工具， 点击：`Go下载 <https://golang.org/dl/>`_ 。

`Go` 提供多种下载版本：

#.  **源码包** 需要先编译才能安装，操作相对比较复杂；
#.  **二进制包** 直接解压就能用，最多再设置一个环境变量即可；
#.  **安装器** 直接双击按照提示点下一步即可(与普通软件安装并无二致)。

推荐优先选择安装器，找到对应系统版本，并下载安装。以 `OSX` 为例：

.. figure:: /_images/quick-start/index/cad4cb4c6f014a928a1be03a449b04c6.png
    :width: 413px

打开终端，输入 ``go`` 按下回车。
如果看到该命令的使用帮助，说明 `Go` 开发环境已经搭建成功了：

.. figure:: /_images/quick-start/index/b0bfb6d2e0364e8827670fd6fc40d050.png
    :width: 285px

你好世界
========

`Go` 开发环境搭建完毕，可以着手写代码了。
以最经典的 `Hello world` 程序为例：

编辑代码
--------

用你熟悉的编辑器编辑代码文件：

.. literalinclude:: /_src/quick-start/hello/main.go
    :caption:
    :name: quick-start/hello/main.go
    :language: go
    :lines: 13-
    :linenos:

这个程序非常简单，先引入 ``fmt`` 包，然后在 ``main`` 函数中向屏幕输出 ``Hello, 世界`` 。

源代码可在 `Github` 上获取： `go-book <https://github.com/fasionchan/go-book>`_ 。

.. note::
    可以用任何编辑器编辑代码。

    Windows下的记事本，Linux及OSX下的文本编辑器都是可以的。
    当然了，功能强大的IDE软件则更好。

    如果喜欢在终端下开发，可以试试 Vim ，我一直是用这个的。

编译程序
--------

源代码需要编译成可执行文件，方能运行。
进入源码目录 ``hello`` ，运行 ``go build`` 命令：

.. code-block:: shell-session

    $ cd quick-start/hello
    $ ls
    main.go
    $ go build
    $ ls
    hello  main.go

看到目录新增一个名为 ``hello`` 的文件，这就是编译好的程序。

注意到，程序名字与源码目录相同。
因为，``go-build`` 以包的形式构建， ``quick-start/hello`` 目录就是代码包， ``hello`` 就是包名。

本文结尾处，以一独立小节解释这些 :ref:`go-build-term` 。

执行程序
--------

接下来，执行程序：

.. code-block:: shell-session

    $ ls
    hello  main.go
    $ ./hello
    Hello, 世界

看到没有，屏幕输出了 ``Hello, 世界`` ，第一个 `Go` 程序运行成功！

.. _go-build-term:

术语
----

最后，总结一下例子中涉及的术语：

.. csv-table:: 表格 术语参照
    :header: "术语", "含义"

    "包", "quick-start/hello"
    "源码", "hello.go"
    "可执行程序", "hello"

下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/disqus.rst

.. include:: /_fragments/wechat-reward.rst

.. comments

    - `OSX <#install-go-osx>`_
    .. _install-go-osx:
