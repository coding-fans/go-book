.. if 语句
    FileName:   if.rst
    Author:     Fasion Chan
    Created:    2018-05-29 10:31:22
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, if, go语言, if语句

=======
if 语句
=======

跟 :doc:`for` 一样， `if` 语句条件表达式不需要括号包围，但是花括号却不能省略。

.. literalinclude:: /_src/tour/if.go
    :caption:
    :name: tour/if.go
    :language: go
    :lines: 13-
    :linenos:

初始语句
========

`if` 语句同样支持初始化语句，在条件表达式求值之前执行。
初始化语句一般都是 ``:=`` 赋值语句。

.. literalinclude:: /_src/tour/if-with-a-short-statement.go
    :caption:
    :name: tour/if-with-a-short-statement.go
    :language: go
    :lines: 13-
    :linenos:

在初始语句中申明的变量，只在 `if` 结构中可见。

if-else 语句
============

在初始化语句定义的变量，在 ``else`` 代码块也是可见的。

.. literalinclude:: /_src/tour/if-and-else.go
    :caption:
    :name: tour/if-and-else.go
    :language: go
    :lines: 13-
    :linenos:

.. note::

    `main` 函数中两次 `pow` 调用都在 `fmt.Println` 调用之前执行。
    因为，调用 `fmt.Println` 前，需要对参数进行求值。

下一步
======

:doc:`下一节 <switch>` 我们一起来看看 `Go` 语言 :doc:`switch` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

