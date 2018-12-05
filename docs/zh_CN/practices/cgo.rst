.. 调用C库函数
    FileName:   cgo.rst
    Author:     Fasion Chan
    Created:    2018-03-04 00:25:56
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
        有些项目没有Go实现，需要在Go中直接调用C库函数。
        cgo就是Go语言调用C库函数的机制，本节以具体的例子演示使用方式。
        gcc -fPIC -c callee.c
        gcc -shared -o libcallee.so callee.o
        #cgo CFLAGS: -I.
        #cgo LDFLAGS: -L. -lcallee
        #include "callee.h"
    :keywords: Go, C, cgo, 库函数, Go语言调用C函数, Go调用C, Makefile

===========
调用C库函数
===========

相对而言，`Go语言 <https://golang.org/>`_ 还是一门非常年轻的语言。

虽然发展迅猛，局限性也有——类库不是特别丰富。
有些开源项目，例如 `x264` ，只有 `C` 语言版本， `Go` 重写一遍也不太现实。

好在， `Go` 提供了一种调用 `C` 函数的机制—— `cgo <https://golang.org/cmd/cgo/>`_ 。
本节以具体的例子演示，如何使用 `cgo` 调用 `C` 函数：

被调用库函数
============

为了演示需要，虚设一个名为 ``callee`` 的函数库。
函数库只提供一个名为 ``sayHello`` 的函数，接口如头文件：

.. literalinclude:: /_src/practices/cgo/callee.h
    :caption:
    :name: cgo/caller.h
    :language: c
    :linenos:

``sayHello`` 函数只是简单输出 ``Hello, world!`` ，实现如下：

.. literalinclude:: /_src/practices/cgo/callee.c
    :caption:
    :name: cgo/caller.c
    :language: c
    :linenos:

接下来，将上述代码编译成动态库：

.. code-block:: shell-session

    $ gcc -fPIC -c callee.c
    $ ls
    Makefile  callee.c  callee.h  callee.o  caller.go

这个命令将 ``callee.c`` 源码编译成 ``callee.o`` 目标文件。
接下来，将目标文件转成成 ``libcallee.so`` 动态库文件：

.. code-block:: shell-session

    $ gcc -shared -o libcallee.so callee.o
    $ ls
    Makefile  callee.c  callee.h  callee.o  caller.go  libcallee.so

调用方
======

接下来看看如何在 `Go` 中调用这个动态库。代码如下：

.. literalinclude:: /_src/practices/cgo/caller.go
    :caption:
    :name: cgo/caller.go
    :language: go
    :linenos:

第 ``16`` 行通过 ``-I`` 选项指定头文件搜索路径，编译器据此发现头文件 ``callee.h`` 。
这里 ``.`` 表示当前目录，可以根据项目情况设置为其他目录。

第 ``17`` 行通过 ``-L`` 选项指定动态库搜索路径，编译器据此发现 ``libcallee.so`` 。
接着，通过 ``-l`` 参数链接到该动态库。

第 ``18`` 行则是引入头文件，据此编译器知晓 `C` 函数接口。

第 ``20`` 通过 ``import`` 关键字引入一个特殊的模块 ``C`` ，之后便可访问到所用链接的 `C` 库函数。
第 ``27`` 行，调用 ``sayHello`` 函数。

.. warning::
    第19行与20行之间不能留空行，不然构建失败！

接下来，编译整个程序：

.. code-block:: shell-session

    $ go build caller.go
    $ ls
    Makefile  callee.c  callee.h  callee.o  caller  caller.go  libcallee.so
    $ ./caller
    Hello, world!
    Success!

看到没有，成功调用 ``sayHello`` 函数并输出 ``Hello, world!`` !

自动化构建
==========

以上例子包含多个编译步骤，需要执行多个命令。

在程序开发中，经常也是如此。
如果每次修改代码后，都需要手工执行这么多命令，那么效率和质量将深受拖累。

我们可以用更自动化的手段进行构建，以 ``Makefile`` 为例：

.. literalinclude:: /_src/practices/cgo/Makefile
    :caption:
    :name: cgo/Makefile
    :language: makefile
    :linenos:

在源码目录准备以上 ``Makefile`` ，之后运行：

.. code-block:: shell-session

    $ make
    gcc -fPIC -c callee.c
    gcc -shared -o libcallee.so callee.o
    go build caller.go
    ./caller
    Hello, world!
    Success!

可以看到， ``make`` 命令根据 `Makefile` 定义自动执行编译命令并执行目标程序。

`Makefile` 更深入的使用方法不在本节的讨论范畴，有兴趣的童鞋自行 `Google <https://www.google.com/>`_ 。

下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst
