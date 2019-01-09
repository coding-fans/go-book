.. Json序列化
    FileName:   jsonify.rst
    Author:     Fasion Chan
    Created:    2018-11-29 16:03:29
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
        JSON是目前最为流行的序列化手段，Go语言内置encoding/json包用于JSON序列化/反序列化操作。
        本文以详细代码示例，演示encoding/json包的使用方式。
    :keywords: golang, json

==========
Json序列化
==========

`JSON` 是目前最为流行的序列化手段， `Go` 语言内置 `encoding/json`_
包用于 `JSON` **序列化** / **反序列化** 操作。
本文以详细代码示例，演示 `encoding/json`_ 包的使用方式。

序列化
======

对于已有类型，序列化只需一行代码：

.. literalinclude:: /_src/practices/jsonify/quick-marshal.go
    :caption:
    :name: practices/jsonify/quick-marshal.go
    :language: go
    :lines: 13-
    :linenos:

例子第 `9-12` 行，定义了一个 `Pair` 类型，包含两个字段；
第 `15-18` 行，定义了一个 `Pair` 变量并初始化；
第 `20` 行对该变量进行 `JSON` 序列化操作，结果如下：

.. code-block:: json

    {"Name":"bar","Value":1}

静态序列化
----------

为了生成 `JSON` 数据，如果类型未定义，则需要先定义新的数据类型。

例如，为了临时生成以下两种不同格式的 `JSON` 数据：

.. code-block:: json

    {"Type":"map","Data":{"Bar":1,"Foo":2}}
    {"Type":"list","Data":[{"Name":"Bar","Value":1},{"Name":"Foo","Value":2}]}

可以先定义以下数据类型：

.. literalinclude:: /_src/practices/jsonify/static-marshal.go
    :caption:
    :name: practices/jsonify/static-marshal.go
    :language: go
    :lines: 23-36
    :linenos:

然后，初始化数据并序列化：

.. literalinclude:: /_src/practices/jsonify/static-marshal.go
    :caption:
    :name: practices/jsonify/static-marshal.go
    :language: go
    :lines: 40-46
    :linenos:

这便是 **静态序列化** 方式，为了生成 `JSON` 数据而定义新的数据类型，好比用牛刀杀鸡。
接着，在 :ref:`jsonify-free-marshal` 一节，介绍一种灵活构建 `JSON` 数据的方法，更加简便。

自定义字段名
------------

序列化结构体，字段名默认与结构体一致。
结构体字段名通过首字母大小写控制访问，因此有时并不满足需要。
例如，结构体中的字段名是驼峰风格，而 `JSON` 字段名却要求用小写字母以及下划线……
这时，只能为结构体字段打上辅助标签：

.. literalinclude:: /_src/practices/jsonify/static-marshal-with-tags.go
    :caption:
    :name: practices/jsonify/static-marshal-with-tags.go
    :language: go
    :lines: 23-36
    :linenos:

这样一来，序列化结果便满足要求了：

.. code-block:: json

    {"type":"map","data":{"bar":1,"value":2}}
    {"type":"list","data":[{"name":"Bar","value":1},{"name":"Foo","value":2}]}

.. _jsonify-free-marshal:

动态序列化
----------

像 `Python` 之类的动态类型语言，我们可以非常自由地组织数据，随时随地：

.. code-block:: python

    json.dumps({
        "type": "map",
        "data": {
            "bar": 1,
            "foo": 2,
        },
    })

然而， `Go` 不是动态类型语言，是不是就没有办法实现动态序列化了呢？

当然不是了，我们可以通过空接口实现。
首先定义几种基本类型：

.. literalinclude:: /_src/practices/jsonify/free-marshal.go
    :caption:
    :name: practices/jsonify/free-marshal.go
    :language: go
    :lines: 21-25
    :linenos:

其中， `Object` 可以是任意类型；
`Array` 可以是任意类型组成的 :doc:`/tour/arrays` ；
`JsonOject` 是一个 :doc:`/tour/maps` ，键为字符类型，值可以是任意类型；
`JsonArray` 与 `Array` 相同。

有了这些基本类型，我们也可以非常灵活的组装数据，同样随时随地：

.. literalinclude:: /_src/practices/jsonify/free-marshal.go
    :caption:
    :name: practices/jsonify/free-marshal.go
    :language: go
    :lines: 28-34
    :linenos:

组装另一种数据类型：

.. literalinclude:: /_src/practices/jsonify/free-marshal.go
    :caption:
    :name: practices/jsonify/free-marshal.go
    :language: go
    :lines: 41-53
    :linenos:

下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

参考文献
========

#. `Go语言中的动态JSON <https://studygolang.com/articles/12172>`_

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. _encoding/json: https://golang.org/pkg/encoding/json/

.. comments
    comment something out below

