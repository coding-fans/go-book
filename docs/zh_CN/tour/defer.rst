.. defer 语句
    FileName:   defer.rst
    Author:     Fasion Chan
    Created:    2018-05-29 10:36:05
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, defer, go语言, defer语句

==========
defer 语句
==========

``defer`` 语句将函数执行推迟到调用函数(包含函数)退出。
函数调用参数还是立马求值，只是执行推迟而已。

你可能已经猜到这个程序输出什么了：

.. literalinclude:: /_src/tour/defer.go
    :caption:
    :name: tour/defer.go
    :language: go
    :lines: 13-
    :linenos:

对，就是这么简单！

defer栈
=======

与函数调用类似，推迟执行的函数调用也被推到一个 **栈** 。
当函数返回时，这些被推迟执行的函数调用将被执行，以后进先出( ``last-in-first-out`` )的顺序。

.. literalinclude:: /_src/tour/defer-multi.go
    :caption:
    :name: tour/defer-multi.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <pointers>` 我们一起来看看 `Go` 语言 :doc:`pointers` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

