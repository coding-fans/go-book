.. 发起HTTP请求
    FileName:   http-request.rst
    Author:     Fasion Chan
    Created:    2018-12-04 18:47:31
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
        Go语言提供了net/http包，协助用户实施与HTTP协议相关的开发任务。
        本文介绍如何使用 net/http 包发起 HTTP 请求，覆盖大部分应用场景。
    :keywords: golang, http, get, post, http header, cookie, json

============
发起HTTP请求
============

在日常开发中，发起 `HTTP` 请求是一个非常频繁的任务。

`Go` 语言提供了 `net/http`_ 包，协助用户实施与 `HTTP` 协议相关的开发任务。
该包既提供了 `HTTP` **服务器** 实现，又提供了 **客户端** 实现。

本文介绍如何使用 `net/http`_ 包发起 `HTTP` 请求，覆盖大部分应用场景：

Get请求
=======

发起 `Get` 请求是最常见的场景之一，使用 `http.Get`_ 函数即可。代码示例如下:

.. literalinclude:: /_src/practices/http-request/get.go
    :name: practices/http-request/get.go
    :language: go
    :lines: 13-
    :linenos:

第 *10* 行调用 `http.Get`_ 函数发起 `Get` 请求；
第 *15* 行在 *main* 函数结束时关闭 *Body* 对象；
第 *17* 行读取整个响应体；
第 *23* 行输出响应内容。

.. warning::

    `http.Get`_ 函数使用默认的客户端设置，默认不超时。
    挂起的请求将劫持 `goroutine` ，严重时可导致程序停止响应。
    因此，禁止在生产业务中使用默认客户端设置。
    :ref:`http-client-timeout` 一节介绍如何设置请求超时时间。

Post请求
========

发起 `Post` 请求的方式也是类似的，代码如下：

.. literalinclude:: /_src/practices/http-request/post.go
    :name: practices/http-request/get.go
    :language: go
    :lines: 13-
    :linenos:

注意到， `http.Post`_ 函数接口与 `http.Get`_ 稍有区别，多了两个参数：

- `Content Type` 头部，这里我们填 ``plain/text`` ；
- 请求体，为 `io.Reader` 类型，没有请求体则填 ``nil`` ；

.. _http-client-timeout:

超时设置
========

超时设置需先准备一个 `http.Client` 结构体， `Timeout` 字段即是超时设置，单位为秒。
之后调用结构体相关请求方法发起请求，示例代码片段如下：

.. literalinclude:: /_src/practices/http-request/get-with-timeout.go
    :name: practices/http-request/get.go
    :language: go
    :lines: 23-27
    :linenos:

.. tip::

    生产业务请务必设置超时时间，以免 `goroutine` 被超时请求劫持而停止响应。

设置头部
========

设置头部须直接操作 `http.Request`_ 结构体，以添加头部 ``X-My-Auth`` 为例：

.. literalinclude:: /_src/practices/http-request/get-with-header.go
    :name: practices/http-request/get.go
    :language: go
    :lines: 13-
    :linenos:

第 `15` 行初始化 `http.Request`_ 结构体，指定请求方法， `URL` 以及请求体；
第 `21` 行调用 `Header` 字段 `Add` 方法添加给定头部；
第 `23` 行调用客户端 `Do` 方法发起请求。

设置Cookie
==========

添加 `Cookie` 也是通过调用 `http.Request` 结构体方法完成，代码片段如下：

.. literalinclude:: /_src/practices/http-request/get-with-cookie.go
    :name: practices/http-request/get-with-cookie.go
    :language: go
    :lines: 27-38
    :linenos:

第 `1` 行初始化 `http.Request`_ 结构体；
第 `7` 行调用 `AddCookie`_ 方法设置设置 `Cookie` ，参数为 `Cookie`_ 结构体，只填充键值字段；

Basic认证
=========

同样，进行 `Basic` 认证也需要操作 `http.Request`_ 结构体。
代码片段如下：

.. literalinclude:: /_src/practices/http-request/get-with-basic-auth.go
    :name: practices/http-request/get-with-basic-auth.go
    :language: go
    :lines: 27-35
    :linenos:

第 `1` 行初始化 `http.Request`_ 结构体；
第 `7` 行调用 `SetBasicAuth`_ 方法设置认证信息。

下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

参考文献
========

#. `http - The Go Programming Language <https://golang.org/pkg/net/http>`_
#. `Don’t use Go’s default HTTP client (in production) – Nathan Smith – Medium <https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779>`_

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. _http.Get: https://golang.org/pkg/net/http/#Client.Get
.. _http.Post: https://golang.org/pkg/net/http/#Client.Post
.. _http.Request: https://golang.org/pkg/net/http/#Request
.. _io.Reader: https://golang.org/pkg/io/#Reader
.. _net/http: https://golang.org/pkg/net/http
.. _AddCookie: https://golang.org/pkg/net/http/#Request.AddCookie
.. _Cookie: https://golang.org/pkg/net/http/#Cookie
.. _SetBasicAuth: https://golang.org/pkg/net/http/#Request.SetBasicAuth

.. comments
    comment something out below

