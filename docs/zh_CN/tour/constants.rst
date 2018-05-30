.. 常量
    FileName:   constants.rst
    Author:     Fasion Chan
    Created:    2018-05-29 09:19:03
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
    :keywords: golang, constants, go语言, 常量

====
常量
====

常量( ``constants`` )申明与变量一样，只不过换成 ``const`` 关键字。
常量可以是字符、字符串、布尔，或者数值类型。
另外，常量不能使用 ``:=`` 语法申明。

.. literalinclude:: /_src/tour/constants.go
    :caption:
    :name: tour/constants.go
    :language: go
    :lines: 13-27
    :linenos:

数值常量
========

数值常量是高精度数值。
常量虽然没有指定类型，却可以根据实际情况采用合适类型，保证精度够用。

试试输出 ``needInt(Big)`` ：

.. literalinclude:: /_src/tour/numeric-constants.go
    :caption:
    :name: tour/numeric-constants.go
    :language: go
    :lines: 13-39
    :linenos:

下一步
======

:doc:`下一节 <for>` 我们一起来看看 `Go` 语言 :doc:`for` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. comments
    comment something out below

