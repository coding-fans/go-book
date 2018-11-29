.. 切片
    FileName:   slices.rst
    Author:     Fasion Chan
    Created:    2018-05-30 13:27:25
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, slice, 切片, 数组引用, len, cap, 长度, 容量, 遍历, 追加

====
切片
====

:doc:`arrays` 长度是固定的。
相反， **切片** ( `slice` )是一个数组元素的弹性视图，长度可以是动态的。
实际上，切片比 :doc:`arrays` 更为常用。

类型 ``[]T`` 就是一个由类型 ``T`` 元素组成的切片。

切片通过两个索引下标定义，一个表示 **下边界** ( `low bound` )，一个表示 **上边界** ( `high bound` )，以冒号( ``:`` )分隔：

.. code-block:: go

    a[low : high]

这表示一个 **半开半闭** 区间，包括第一个元素，但不包括最后一个。

以下表达式创建一个包含元素 `1` 到元素 `3` 的切片(不包括元素 `4`)：

.. code-block:: go

    a[1:4]

完整例子如下：

.. literalinclude:: /_src/tour/slices.go
    :caption:
    :name: tour/slices.go
    :language: go
    :lines: 13-
    :linenos:

数组引用
========

切片实际上并 **不存储任何数据** ，只是用来描述关联数组的一部分。
因此，修改切片元素等价于修改对应的数组元素，其他共用该数组的切片对此也可见。
换句话讲， **切片就像是数组的引用** 。

.. literalinclude:: /_src/tour/slice-pointers.go
    :caption:
    :name: tour/slice-pointers.go
    :language: go
    :lines: 13-
    :linenos:

下标默认值
==========

定义切片，可以忽略上边界或者下边界，而使用 **默认值** 。
对于下边界，默认值为 `0` ；下边界，默认值则是 **切片长度** 。

因此，对于数组：

.. code-block:: go

    var a [10]int

以下切片表达式均是等价的：

.. code-block:: go

    a[0:10]
    a[:10]
    a[0:]
    a[:]

.. literalinclude:: /_src/tour/slice-bounds.go
    :caption:
    :name: tour/slice-bounds.go
    :language: go
    :lines: 13-
    :linenos:

长度和容量
==========

切片有两个重要属性， **长度** ( `length` )和 **容量** ( `capacity` )。

切片的长度等于切片包含的元素个数。

切片的容量等于底下数组的元素个数，从切片第一个元素算起。

对于切片 `s` ，长度和容量分别可以通过表达式 ``len(s)`` 以及 ``cap(s)`` 获得。

通过 **重切** ( `re-slicing` )，你可以扩张一个切片的长度，只要容量足够。
你可以试试修改下面这个例子，将切片扩张到超出容量，看看会发生什么事情：

.. literalinclude:: /_src/tour/slice-len-cap.go
    :caption:
    :name: tour/slice-len-cap.go
    :language: go
    :lines: 13-
    :linenos:

空切片
======

切片的 :ref:`zero-value` 是 `nil` ，即 **空切片** 。
空切片长度和容量均为 `0` ，当然也不需要底层数组。

.. literalinclude:: /_src/tour/nil-slices.go
    :caption:
    :name: tour/nil-slices.go
    :language: go
    :lines: 13-
    :linenos:

make
====

切片可以由内置函数 `make`_ 来创建，相当于你可以创建动态长度的数组。

`make` 函数分配一个由 :ref:`zero-value` 填充的数组，并返回一个引用该数组的切片：

.. code-block:: go

    a := make([]int, 5) // len(a)=5

要指定容量，可以通过 `make` 函数第三个参数指定：

.. code-block:: go

    b := make([]int, 0, 5) // len(b)=0, cap(b)=5

    b = b[:cap(b)]  // len(b)=5, cap(b)=5
    b = b[1:]       // len(b)=4, cap(b)=4

完整例子如下：

.. literalinclude:: /_src/tour/making-slices.go
    :caption:
    :name: tour/making-slices.go
    :language: go
    :lines: 13-
    :linenos:

切片的切片
==========

切片可以包含任何类型，当然包括其他切片。

.. literalinclude:: /_src/tour/slices-of-slices.go
    :caption:
    :name: tour/slices-of-slices.go
    :language: go
    :lines: 13-
    :linenos:

元素追加
========

向切片追加元素是一个很常用的操作，为此 `Go` 提供了一个 **内置函数** 。
`内置包文档 <https://golang.org/pkg/builtin/>`_ 详细描述了这个内置函数 `append`_ ：

.. code-block:: go

    func append(s []T, vs ...T) []T

`append` 函数第一个参数 `s` 是一个类型为 `T` 的切片，其余参数均为追加至 `s` 的 `T` 元素。

`append` 函数返回一个新切片，包含原切片以及所有追加元素。

如果底层数组太小， `append` 函数会分配一个更大的数组，新切片则指向新数组。

.. literalinclude:: /_src/tour/append.go
    :caption:
    :name: tour/append.go
    :language: go
    :lines: 13-
    :linenos:

遍历
====

配合 `range` 关键字， `for` 循环可对切片进行 **遍历** 。
下节，我们将看到，这种做法也适用于 :doc:`maps` ( `map` )。

每次迭代都返回两个值，第一个是 **下标** ( `index` )，第二个是与该下标对应的 **元素拷贝** 。

.. literalinclude:: /_src/tour/range.go
    :caption:
    :name: tour/range.go
    :language: go
    :lines: 13-
    :linenos:

如果无须下标，可以直接赋值给下划线 ``_`` 。
如果只要下标，则不写后半部分即可。

.. literalinclude:: /_src/tour/range-continued.go
    :caption:
    :name: tour/range-continued.go
    :language: go
    :lines: 13-
    :linenos:

练习
====

1. 图片像素阵列
---------------

尝试实现一个函数 `Pic` ：返回一个长度为 `dy` 的切片，切片元素为长度为 `dx` 的 `8` 位无符号整数切片。
这个切片，其实就是一个二维阵列，可以用来表示一张图片的像素。

当你运行这个程序时，它将展示一张图片，每个数值被解释为图片像素的 **灰度值** ( `grayscale` )。

代码框架如下：

.. literalinclude:: /_src/tour/exercise-slices-pic.go
    :caption:
    :name: tour/exercise-slices-pic.go
    :language: go
    :lines: 13-
    :linenos:

你可以在 `Go官网 <https://tour.golang.org/moretypes/18>`_ 进行练习。
图片长啥样取决于你返回的切片。
试试以下以下公式来生成切片，结果将非常有趣：

.. math::
    \frac{x+y}{2}

.. math::
    x * y

.. math::
    x ^ y

这三个公式生成的图片分别是：

.. figure:: /_images/tour/slices/1fb9b097a97fde32fcd3c8292605be52.png
    :align: center

    (x+y) / 2

.. figure:: /_images/tour/slices/7261fd8ba2ae27e30e9fcf93a1ed900e.png
    :align: center

    x * y

.. figure:: /_images/tour/slices/0625993b39e53718ec2437d14d01f389.png
    :align: center

    x ^ y

答案
++++

相信你已经写出自己的程序，看到各种有趣的图片了。

如果练习过程中遇到什么问题，可以参考下面这个例子并尝试解决：

.. literalinclude:: /_src/tour/solution-slices-pic.go
    :caption:
    :name: tour/solution-slices-pic.go
    :language: go
    :lines: 13-
    :linenos:

下一步
======

:doc:`下一节 <maps>` 我们一起来看看 `Go` 语言 :doc:`maps` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. _append: https://golang.org/pkg/builtin/#append
.. _make: https://golang.org/pkg/builtin/#make

.. comments
    comment something out below
