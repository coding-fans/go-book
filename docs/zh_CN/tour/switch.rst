.. switch 语句
    FileName:   switch.rst
    Author:     Fasion Chan
    Created:    2018-05-29 10:33:16
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, switch, go语言, switch语句

===========
switch 语句
===========

`switch` 语句也是一种经典控制语句，可以看做是 `if-else` 语句链的简写。

.. literalinclude:: /_src/tour/switch.go
    :caption:
    :name: tour/switch.go
    :language: go
    :lines: 13-
    :linenos:

`Go` 语言 `switch` 结构跟 `C` 、 `C++` 、 `Java` 、 `JavaScript` 以及 `PHP` 等类似，
不同的是， `Go` 只执行匹配的 `case` 代码体，不包括下面的。
对于其他语言，一般需要在每个 `case` 末尾处用 `break` 语句来结束。
实际上， `Go` 相当于自动在每个 `case` 末尾添加了 `break` 语句，
这避免了大量因为漏掉 `break` 而导致的错误。

另外， `Go` 语言 `switch` 更为灵活， `case` 条件不必是常量，也不必是整数。

检查顺序
========

`switch` 从上往下对 `case` 进行检查，直到匹配。

.. literalinclude:: /_src/tour/switch-evaluation-order.go
    :caption:
    :name: tour/switch-evaluation-order.go
    :language: go
    :lines: 13-
    :linenos:

举另一个例子，如果 ``i`` 的值为零，那么函数 ``f`` 就不会被调用了：

.. code-block:: go

    switch i {
        case 0:
        case f():
    }

省略条件
========

在 `Go` 语言， `switch` 条件可以省略，等价于 ``switch true`` 。
这种结构非常简洁，可以用来代替过长的 `if-else` 链。

.. literalinclude:: /_src/tour/switch-with-no-condition.go
    :caption:
    :name: tour/switch-with-no-condition.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <defer>` 我们一起来看看 `Go` 语言 :doc:`defer` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

