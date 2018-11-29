.. for 语句
    FileName:   for.rst
    Author:     Fasion Chan
    Created:    2018-05-29 09:55:48
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, for, go语言, for语句

========
for 语句
========

`Go` 只有一种循环结构—— `for` 循环。

最基本的 `for` 语句包含 3 部分，以分号 ``;`` 分隔：

- **初始语句** ：在第一次迭代判断之前；
- **条件语句(表达式)** ：在每次迭代前求值并判断；
- **迭代后语句** ：在每次迭代后执行；

.. code-block:: go

    for 初始语句; 条件语句; 迭代后语句 {
        代码体
    }

当条件表达式求值为 ``false`` 时，循环将停止迭代并退出。

.. literalinclude:: /_src/tour/for.go
    :caption:
    :name: tour/for.go
    :language: go
    :lines: 13-
    :linenos:

.. note::

    跟 `C` 、 `Java` 或者 `JavaScript` 等其他语言不同，
    `Go` 语言 `for` 语句 3 部分不需要用括号包住，
    但花括号是必要的，任何时候都 **不能省略** 。

另外，与其他语言类似，初始语句与迭代后语句也是 **可选** 的：

.. literalinclude:: /_src/tour/for-continued.go
    :caption:
    :name: tour/for-continued.go
    :language: go
    :lines: 13-
    :linenos:

实际上， `Go` 也是支持 `while` 语句的，只不过关键字还换成 `for` ：

.. literalinclude:: /_src/tour/for-is-gos-while.go
    :caption:
    :name: tour/for-is-gos-while.go
    :language: go
    :lines: 13-
    :linenos:

无限循环
========

如果省略循环条件，循环将 **永远执行** 。
这种循环就是众所周知的 **死循环** ，也叫做 **无限循环** 。
对我来说，我更愿意用 **无限循环** 。
因为， **死循环** 更应该用在程序有问题，循环行为不符合作者预期的场景。

无限循环是 `Go` 语言中最紧凑的循环结构：

.. literalinclude:: /_src/tour/forever.go
    :caption:
    :name: tour/forever.go
    :language: go
    :lines: 13-
    :linenos:

.. warning::

    使用无限循环时要特别小心！

下一步
======

:doc:`下一节 <if>` 我们一起来看看 `Go` 语言 :doc:`if` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

