.. 数组
    FileName:   arrays.rst
    Author:     Fasion Chan
    Created:    2018-05-30 13:07:20
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

====
数组
====

`Go` 也有 :doc:`arrays` ， ``[n]T`` 就表示一个由 ``n`` 个类型 ``T`` 元素组成的数组类型。

下面这个表达式，申明了一个由 10 个整数组成的数组变量：

.. code-block:: go

    var a [10]int

数组的长度是类型的一部分(不同长度意味着不同类型)，所以数组没有办法调整尺寸。
这看上去很有局限性；然而并不用太担心， `Go` 提供的方式很方便。

.. literalinclude:: /_src/tour/arrays.go
    :caption:
    :name: tour/arrays.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <slices>` 我们一起来看看 `Go` 语言 :doc:`slices` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. comments
    comment something out below

    .. meta::
        :description lang=zh:
        :keywords:

