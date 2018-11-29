.. 函数
    FileName:   functions.rst
    Author:     Fasion Chan
    Created:    2018-05-28 15:03:41
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, function, go语言, 函数

====
函数
====

`Go` 语言中，函数可以接受零或多个参数：

.. literalinclude:: /_src/tour/functions.go
    :caption:
    :name: tour/functions.go
    :language: go
    :lines: 13-
    :linenos:

在这个例子中， ``add`` 函数接受两个整型( ``int`` )参数。
注意到，类型申明紧跟在参数名之后，与其他语言有些区别。

如果参数类型相同，则在最后一个申明即可，前面的可以省略。

因此，可以将 ``x int, y int`` 简写成 ``x, y int`` 。

返回多个值
==========

`Go` 函数可以非常优雅地返回多个值，比起定义结构体返回指针之类的舒服多啦！
写个简单的程序交换两个字符串：

.. literalinclude:: /_src/tour/multiple-results.go
    :caption:
    :name: tour/multiple-results.go
    :language: go
    :lines: 13-
    :linenos:

命名返回值
==========

`Go` 函数返回值可以被命名( ``named`` )，命名后当做函数参数来对待。
命名的意义在于指明各个返回值含义。

.. literalinclude:: /_src/tour/named-results.go
    :caption:
    :name: tour/named-results.go
    :language: go
    :lines: 13-
    :linenos:

一个不带任何参数的 ``return`` 语句返回所有命名返回值，
这就是所谓的 **裸返回** ( ``naked return`` )。
裸返回只推荐在 **短函数** 中使用，如在 **长函数** 中滥用，则影响 **代码可读性** 。

下一步
======

:doc:`下一节 <variables>` 我们一起来看看 `Go` 语言 :doc:`variables` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

