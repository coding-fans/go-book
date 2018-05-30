.. 基本类型
    FileName:   basic-types.rst
    Author:     Fasion Chan
    Created:    2018-05-29 08:48:52
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
    :keywords:

========
基本类型
========

`Go` 内置了以下基本类型：

- **布尔**
    - `bool`
- **字符串**
    - `string`
- **整数**
    - `int` `int8` `int16` `int32` `int64`
    - `uint` `uint8` `uint16` `uint32` `uint64`
- **字节**
    - `byte` ， `uint8` 的别名
- **Unicode**
    - `rune` ， `int32` 别名
- **浮点**
    - `float32` `float64`
- **复数**
    - `complex64` `complex128`

.. literalinclude:: /_src/tour/basic-types.go
    :caption:
    :name: tour/basic-types.go
    :language: go
    :lines: 13-31
    :linenos:

例子展示了几种不同类型变量的用法。
注意到，跟 :ref:`import-statement` 一样，变量申明可以批量写到一个代码块里。

一般来说，`int` 、 `uint` 以及 `uintptr` 类型在 ``32`` 位机器上是 ``32`` 位长；
在 ``64`` 位机器上则是 ``64`` 位长。
需要使用整数时， ``int`` 类型是首选，
除非你有特别的理由一定要用 **定长** 或者 **无符号** 类型。

零值
====

变量申明时没有显式赋初始值，则默认是“ **零** ”。

不同的类型有不同的“ **零** ”：

- 对于数值类型是 ``0`` ；
- 对于布尔类型是 ``false`` ；
- 对于字符串类型是 ``""`` (空字符串)；

.. literalinclude:: /_src/tour/zero.go
    :caption:
    :name: tour/zero.go
    :language: go
    :lines: 13-24
    :linenos:

类型转换
========

**表达式** ( ``expression`` ) ``T(v)`` 将值 ``v`` 转换成类型 ``T`` ，
这就是所谓的 **类型转换** ( ``type conversions`` )。

这是一些数值类型转换：

.. code-block:: go

    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)

或者简写成：

.. code-block:: go

    i := 42
    f := float(i)
    u := uint(f)

跟 `C` 语言有所不同， `Go` 在不同类型之间赋值，需要显式类型转换。
不信，将下面例子中 ``float64`` 和 ``unit`` 类型转换移除，看看发生什么？

.. literalinclude:: /_src/tour/type-conversions.go
    :caption:
    :name: tour/type-conversions.go
    :language: go
    :lines: 13-26
    :linenos:

类型推理
========


变量类型通过右边的值推理而来。

如果申明右边的值是有类型的，那么新变量也是一样的类型：

.. code-block:: go

    var i int
    j := i  // j is an int as well

如果右边只是一个数值常量，没有具体类型，那么新变量可能是 `int` 、
`float64` 以及 `complex128` 三种类型中的一种，取决于常量的精度。

.. code-block:: go

    i := 42             // int
    f := 3.142          // float64
    g := 0.867 + 0.5i   // complex128

接下来做个实验吧！
改变例子中 ``v`` 的初始值，观察它是如何影响变量类型的：

.. literalinclude:: /_src/tour/type-inference.go
    :caption:
    :name: tour/type-inference.go
    :language: go
    :lines: 13-21
    :linenos:

下一步
======

:doc:`下一节 <constants>` 我们一起来看看 `Go` 语言 :doc:`constants` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. comments
    comment something out below

