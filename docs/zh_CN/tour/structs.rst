.. 结构体
    FileName:   structs.rst
    Author:     Fasion Chan
    Created:    2018-05-30 12:57:12
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, structs, go语言, 结构体

======
结构体
======

`Go` 语言也有结构体——由若干 :ref:`tour.structs-fields` ( ``field`` )组成的集合。

.. literalinclude:: /_src/tour/structs.go
    :caption:
    :name: tour/structs.go
    :language: go
    :lines: 13-
    :linenos:

例子定义了一个名为 ``Vertex`` 的结构体，用于表示一个顶点。
结构体包含两个字段，分别是 ``X`` 和 ``Y`` ，类型都是 ``int`` 。

.. _tour.structs-fields:

字段
====

结构体字段可以通过点号 ``.`` 操作符访问。

.. literalinclude:: /_src/tour/struct-fields.go
    :caption:
    :name: tour/struct-fields.go
    :language: go
    :lines: 13-
    :linenos:

结构体指针
==========

结构体字段可以通过 :doc:`/tour/pointers` 来访问。

假设我们有一个结构体指针 ``p`` ，访问字段 ``X`` 理论上可以这么写 ``(*p).X`` 。
然而，这看上去很累赘不是？
作为很人性化的语言， `Go` 允许直接写成 ``p.X`` 。
这两种写法没有什么实质性的区别，但是后者显然更为清晰。

.. literalinclude:: /_src/tour/struct-pointers.go
    :caption:
    :name: tour/struct-pointers.go
    :language: go
    :lines: 13-
    :linenos:

字面量
======

结构体字面量( ``literal`` )，即通过列举字段值来表示一个新分配的结构体。

可以通过 ``字段名:`` 语法指定部分字段值，其他字段则默认为 :ref:`zero-value` 。
由于指定了字段名，字段列举顺序也就无关紧要了。

在结构体之前加上 ``&`` 操作符，则返回对应的结构体指针。

.. literalinclude:: /_src/tour/struct-literals.go
    :caption:
    :name: tour/struct-literals.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <arrays>` 我们一起来看看 `Go` 语言 :doc:`arrays` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

