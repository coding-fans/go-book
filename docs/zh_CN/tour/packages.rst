.. 包
    FileName:   packages.rst
    Author:     Fasion Chan
    Created:    2018-05-28 14:36:54
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, package, go语言, 包

==
包
==

每个 `Go` 程序都是由一些包组成的。

程序从 `main` 包开始执行。

.. literalinclude:: /_src/tour/packages.go
    :caption:
    :name: tour/package.go
    :language: go
    :lines: 13-
    :linenos:

在这个程序，通过 import 导入两个包， ``fmt`` 和 ``math/rand`` (包路径)。

按照惯例，包名与包路径最后部分相同。
例如， ``math/rand`` 包中的源码文件都以 ``package rand`` 语句开头。

.. _import-statement:

import语句
==========

`Go` 通过 ``import`` 语句引入包并在代码中使用。

`import` 语句有两种不同的写法，上面例子是其中的一种写法—— **批量导入** ；
第二种则是分成多个语句：

.. code-block:: go

    import "fmt"
    import "math"

两种写法虽然没有实质区别，还是 **推荐采用批量写法** ，这是 **最佳风格** 。

名字导出
========

在 `Go` 语言，以大写字母开头的名字就会被 **导出** ( ``exported`` )。
举例， ``Pizza`` 就是一个导出名字， ``math`` 包中的 ``Pi`` 也是。

.. literalinclude:: /_src/tour/exported-names.go
    :caption:
    :name: tour/exported-names.go
    :language: go
    :lines: 13-
    :linenos:

相反， ``pizza`` 和 ``pi`` 由于不是大写字母开头，因此不会被导出。

一个包导入后，只能引用到导出名字。
其他任何非导出名字在包外是没有办法访问到的(不可见)。

下一步
======

:doc:`下一节 <functions>` 我们一起来看看 `Go` 语言 :doc:`functions` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. comments
    comment something out below

