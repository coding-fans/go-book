.. 函数值
    FileName:   function-values.rst
    Author:     Fasion Chan
    Created:    2018-05-30 20:34:34
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, 函数值, 闭包, function value, closure

======
函数值
======

在 `Go` 语言，函数也是一种值( `C++` 函数对象、 `Python <https://python-book.readthedocs.io/zh_CN/latest/>`_ 函数类似)，可以被传递。

跟其他普通值一样，函数也可以作为 **参数** 传递或作为 **返回值** 返回。

.. literalinclude:: /_src/tour/function-values.go
    :caption:
    :name: tour/function-values.go
    :language: go
    :lines: 13-
    :linenos:

闭包
====

`Go` :doc:`functions` 可以是 **闭包** 。
闭包是指一个引用外部变量的 :doc:`function-values` (函数对象)。
闭包函数可以访问外部变量，也可以为其赋值。
如此看来，闭包函数相当于与外部变量绑在一起。

.. literalinclude:: /_src/tour/function-closures.go
    :caption:
    :name: tour/function-closures.go
    :language: go
    :lines: 13-
    :linenos:

这个例子， ``adder`` 函数返回了一个闭包函数。
每个闭包函数都与自己的 ``sum`` 变量绑定，互相独立。

练习
====

1. 斐波那契数列
---------------

接下来做运用函数知识做一些有趣的事情：

实现一个函数 `fibonacci` ，返回一个闭包函数。
每次调用该闭包函数时，均返回下一个斐波那契数(从零开始)。

代码框架如下：

.. literalinclude:: /_src/tour/exercise-fibonacci-closure.go
    :caption:
    :name: tour/exercise-fibonacci-closure.go
    :language: go
    :lines: 13-
    :linenos:

答案
++++

.. literalinclude:: /_src/tour/solution-fibonacci-closure.go
    :caption:
    :name: tour/solution-fibonacci-closure.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <index>` 我们一起来看看 `Go` 语言 :doc:`methods` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

