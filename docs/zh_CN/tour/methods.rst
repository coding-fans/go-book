.. 方法
    FileName:   methods.rst
    Author:     Fasion Chan
    Created:    2018-08-02 20:15:56
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, 方法, method

====
方法
====

`Go` 语言没有类的概念。
但是，你可以为某个类型定义 **方法** ( `method` )。

**方法** 是一个带 **接收者参数** 的特殊函数。
接收者参数位于 `func` 关键字与方法名之间，以括号包围。

下面这个例子中， `Abs` 方法有一个 `Vertex` 类型的接收者参数 `v` ：

.. literalinclude:: /_src/tour/methods.go
    :caption:
    :name: tour/methods.go
    :language: go
    :lines: 13-
    :linenos:

再次强调： **方法只是一个带有接收者参数的函数而已** 。

你可以重写 `Abs` ，将其实现成一个普通函数，功能上并没有任何区别：

.. literalinclude:: /_src/tour/methods-funcs.go
    :caption:
    :name: tour/methods-funcs.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <index>` 我们一起来看看 `Go` 语言 :doc:`index` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

