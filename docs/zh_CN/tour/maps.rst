.. 映射表
    FileName:   maps.rst
    Author:     Fasion Chan
    Created:    2018-05-30 13:31:57
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, map, 映射表, 字典, 集合

======
映射表
======

映射表( ``map`` )是一种将 **键** ( ``key`` )映射到 **值** ( ``value`` )的数据结构。

映射表的零值是 ``nil`` 。
一个 ``nil`` 映射表既不包含任何键，也不能添加。

映射表同样可以通过 `make`_ 函数来创建并初始化：

.. literalinclude:: /_src/tour/maps.go
    :caption:
    :name: tour/maps.go
    :language: go
    :lines: 13-
    :linenos:

字面量
======

映射表也可以用 **字面量** ( `literal` )定义。
写法与 :doc:`structs` 字面量类似，额外写上 **键** 而已。

.. literalinclude:: /_src/tour/map-literals.go
    :caption:
    :name: tour/map-literals.go
    :language: go
    :lines: 13-
    :linenos:

当然了，你也可以省略类型：

.. literalinclude:: /_src/tour/map-literals-continued.go
    :caption:
    :name: tour/map-literals-continued.go
    :language: go
    :lines: 13-
    :linenos:

操作
====

**插入** 或 **更新** 映射表 `m` 中的一个元素：

.. code-block:: go

    m[key] = elem

取出一个元素：

.. code-block:: go

    elem = m[key]

删除一个元素：

.. code-block:: go

    delete(m, key)

通过 **双赋值** ( `two-value assignment` )测试给定键是否存在：

.. code-block:: go

    elem, ok = m[key]

这样一来，如果 `k` 在 `m` 中存在， `ok` 便是 ``true`` ；否则， `ok` 则是 ``false`` ，而 `elem` 则是字典元素类型的 :ref:`zero-value` 。

.. literalinclude:: /_src/tour/mutating-maps.go
    :caption:
    :name: tour/mutating-maps.go
    :language: go
    :lines: 13-
    :linenos:

.. note::

    注意到，如果变量 `elem` 或者 `ok` 未声明，可以采用 :ref:`variables-short-declaration` 。即：

    ``elem, ok := m[key]``

下一步
======

:doc:`下一节 <function-values>` 我们一起来看看 `Go` 语言 :doc:`function-values` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. comments
    comment something out below

.. _make: https://golang.org/pkg/builtin/#make
