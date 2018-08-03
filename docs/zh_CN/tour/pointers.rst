.. 指针
    FileName:   pointers.rst
    Author:     Fasion Chan
    Created:    2018-05-30 10:44:18
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, 指针, pointer

====
指针
====

`Go` 语言也有指针。
指针用来保存一个值的内存地址。

类型 `*T` 就是类型 `T` 的指针类型。
指针的 :ref:`zero-value` ( ``zero value`` )是： ``nil`` 。

.. code-block:: go

    var p *int

操作符( ``operator`` ) ``&`` 用来取被操作数( ``operand`` )的指针(内存地址)：

.. code-block:: go

    i := 42
    p = &i

操作符 ``*`` 来用取出指针指向的值：

.. code-block:: go

    fmt.Println(*p)
    *p = 21

这个操作简称 **取值** ( ``dereferencing`` )。

.. literalinclude:: /_src/tour/pointers.go
    :caption:
    :name: tour/pointers.go
    :language: go
    :lines: 13-
    :linenos:

.. note::

    与 `C` 语言不同， `Go` 语言指针 **没有指针算术** ( `pointer arithmetic` )。

下一步
======

:doc:`下一节 <structs>` 我们一起来看看 `Go` 语言 :doc:`structs` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. comments
    comment something out below

