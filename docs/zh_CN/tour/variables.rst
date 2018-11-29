.. 变量
    FileName:   variables.rst
    Author:     Fasion Chan
    Created:    2018-05-29 08:31:48
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, variables, go语言, 变量

====
变量
====

``var`` 语句 **申明** ( ``declare`` )变量列表；
跟 :doc:`functions` 参数列表一样，类型在最后指定。

``var`` 语句的作用域(可见范围)可以是 **包级别** 或者 **函数级别** 。
下面这个例子同时包含这两种级别：

.. literalinclude:: /_src/tour/variables.go
    :caption:
    :name: tour/variables.go
    :language: go
    :lines: 13-
    :linenos:

作用域是啥意思呢？

以上述代码为例，变量 ``c`` 、 ``python`` 、 ``java`` 的作用域是 **包级别** ，
意味着包内任何函数都可以访问这些变量；
定义在函数内部的 ``i`` 则是 **函数级别** ，
只有在 ``main`` 函数内部才能访问。

初始值
======

变量申明可以带初始值，一个变量一个。
在初始值存在的情况下，类型可以忽略；变量则继承初始值的类型。

.. literalinclude:: /_src/tour/variables-with-initializers.go
    :caption:
    :name: tour/variables-with-initializers.go
    :language: go
    :lines: 13-
    :linenos:

.. _variables-short-declaration:

简式申明
========

在函数内部，可以用 ``:=`` 赋值语句代替 ``var`` 变量申明语句，
变量类型也可以省略，这就是 **简式申明** 。

.. literalinclude:: /_src/tour/short-variable-declarations.go
    :caption:
    :name: tour/short-variable-declarations.go
    :language: go
    :lines: 13-
    :linenos:

在函数外部，每个语句都必须由一个关键字开始(如 ``var`` 、 ``func`` 等)，
``:=`` 语句不可用。

下一步
======

:doc:`下一节 <basic-types>` 我们一起来看看 `Go` 语言 :doc:`basic-types` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

