.. 方法
    FileName:   methods.rst
    Author:     Fasion Chan
    Created:    2018-08-02 20:15:56
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :keywords: golang, 方法, method

====
方法
====

`Go` 语言没有类的概念。
但是，你可以为某个类型定义 **方法** ( `method` )。

**方法** 是一个带 **接收者参数** 的特殊函数。
接收者参数位于 `func` 关键字与方法名之间，以括号包围。

下面这个例子中， ``Abs`` 方法有一个 ``Vertex`` 类型的接收者参数 ``v`` ：

.. literalinclude:: /_src/tour/methods.go
    :caption:
    :name: tour/methods.go
    :language: go
    :lines: 13-
    :linenos:

再次强调： **方法只是一个带有接收者参数的函数而已** 。

你可以重写 ``Abs`` ，将其实现成一个普通函数，功能上并没有任何区别：

.. literalinclude:: /_src/tour/methods-funcs.go
    :caption:
    :name: tour/methods-funcs.go
    :language: go
    :lines: 13-
    :linenos:

非结构体方法
============

不仅 :doc:`structs` 可以定义方法，其他任何自定义类型均可。

以下就是一例，为数值类型 ``MyFloat`` 定义方法 ``Abs`` ：

.. literalinclude:: /_src/tour/methods-continued.go
    :caption:
    :name: tour/methods-continued.go
    :language: go
    :lines: 13-
    :linenos:

方法和对应类型定义必须在同一个 :doc:`packages` 定义。

指针接收者
==========

方法接收者可以定义成 :doc:`pointers` 。

这样一来，对于类型 ``T`` 来说，接收者参数的类型就是 ``*T`` 。
需要注意的是， ``T`` 本身不能是指针，比如 ``*int`` 。

例子中， ``Scale`` 方法就定义在 ``*Vertex`` 上：

.. literalinclude:: /_src/tour/methods-pointers.go
    :caption:
    :name: tour/methods-pointers.go
    :language: go
    :lines: 13-
    :linenos:

接收者参数定义成指针的好处是，方法代码可以修改指针指向的值。
由于方法经常需要修改对应的值，因此指针接收者参数相对来说更常用。

读者可以自行修改程序，将 ``*`` 号从 ``Scale`` 方法移除，并观察程序行为。
不出意外，你将看到程序输出 ``5`` 。换句话讲，并没有修改到目标值。
这是为啥呢？

如果定义值接收者( `value receiver` )， `Scale` 方法相当于在原 ``Vertex`` 值的一个拷贝上操作(适用于其他参数)。
因此，为了修改 ``Vertex`` 值，接收者参数必须定义成指针。

传值与传引用
============

接下来，我们将 ``Abs`` 和 ``Scale`` 方法重写成普通函数。

.. literalinclude:: /_src/tour/methods-pointers-explained.go
    :caption:
    :name: tour/methods-pointers-explained.go
    :language: go
    :lines: 13-
    :linenos:

同样，将 ``*`` 号从 ``Scale`` 函数移除会怎样？
不出意外，结果是类似的。

这其实是编程里最经典的 **传值** 、 **传引用** 问题， **传指针相当于传引用** 。

间接传指针
==========

对比上面两个程序，你可能已经注意到了——带指针参数的函数只能传指针：

.. code-block:: go

    var v Vertex
    ScaleFunc(v, 5)     // Compile error!
    ScaleFunc(&v, 5)    // OK

然而，对于方法，不管接收者是一个值还是指针，均可调用：

.. code-block:: go

    var v Vertex
    v.Scale(5)      // OK

    p := &v
    p.Scale(10)     // OK

对于语句 ``v.Scale(5)`` ，尽管 ``v`` 是一个值而不是指针，还是自动调用了带指针接收者参数的方法。
这是因为，``Scale`` 方法需要指针接收者参数， `Go` 按照惯例将 ``v.Scale(5)`` 解释成： ``(&v).Scale(5)`` 。
这就是 **间接传指针** ，或者叫做 **隐式传指针** 。

.. literalinclude:: /_src/tour/indirection.go
    :caption:
    :name: tour/indirection.go
    :language: go
    :lines: 13-
    :linenos:

间接传值
========

对普通 :doc:`functions` 来说，值参数只能传对应类型的值，传指针则导致编译错误：

.. code-block:: go

    var v Vertex
    fmt.Println(AbsFunc(v))     // OK
    fmt.Println(AbsFunc(&v))    // Compile error!

相反，就算方法定义了值接收者参数，用指针调用也是可以的：

.. code-block:: go

    var v Vertex
    fmt.Println(v.Abs())    // OK

    p := &v
    fmt.Println(p.Abs())    // OK

在这，方法调用语句 ``p.Abs()`` 则被解释成： ``(*p).Abs()`` 。
这就是 **间接传值** ，或者叫做 **隐式传值** 。

.. literalinclude:: /_src/tour/indirection-values.go
    :caption:
    :name: tour/indirection-values.go
    :language: go
    :lines: 13-
    :linenos:

传值还是传指针
==============

那么，接收者参数到底是实现成值还是指针呢？
如何选择？

使用指针接收者参数主要有两方面考虑：

首先，只有这种方式能够对指向的值进行修改。

其次，从性能方面考虑，使用指针可以避免在每次调用方法时拷贝值。
这种方式相对来说更高效，特别是当接收者 :doc:`structs` 很大很复杂时。

在这个例子， ``Scale`` 方法和 ``Abs`` 方法接收者参数类型均为 ``*Vertex`` ，尽管 ``Abs`` 方法并不修改其接收者：

.. literalinclude:: /_src/tour/methods-with-pointer-receivers.go
    :caption:
    :name: tour/methods-with-pointer-receivers.go
    :language: go
    :lines: 13-
    :linenos:

通常，不管为何种类型编写方法，均需要定义 **值接收者** 或者 **指针接收者** ，不能混用。

下一步
======

:doc:`下一节 <index>` 我们一起来看看 `Go` 语言 :doc:`interfaces` 。

.. include:: /_fragments/next-step-to-wechat-mp.rst

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. comments
    comment something out below

