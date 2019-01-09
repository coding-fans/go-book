.. 提供HTTP服务
    FileName:   http-server.rst
    Author:     Fasion Chan
    Created:    2019-01-02 18:04:24
    @contact:   fasionchan@gmail.com
    @version:   $Id$

    Description:

    Changelog:

.. meta::
    :description lang=zh:
        发起HTTP请求一文介绍了如何使用net/http包发起HTTP请求。
        net/http包同样提供了用于开发实现HTTP服务的基础类库，本文将进一步介绍如何使用这些类库。
    :keywords: golang, http, server, api, net/http, ServeHTTP

============
提供HTTP服务
============

:doc:`http-request` 一文介绍了如何使用 `net/http`_ 包发起 `HTTP` 请求。
`net/http`_ 包同样提供了用于开发实现 `HTTP` 服务的基础类库，本文将进一步介绍如何使用这些类库。

先来看看一个简单的例子——计数器服务接口。
这个接口非常简单，每次均返回接口调用总次数，以 `1` 开始：

.. literalinclude:: /_src/practices/http-server/counter-server.go
    :name: practices/http-server/counter-server.go
    :language: go
    :lines: 13-
    :linenos:

代码中，第 *9* 行声明一个用于记录次数的变量 *counter* ，初始值为 *0* ；
第 *13-16* 行定义请求处理器 *CounterHandler* 及其处理函数 *ServeHTTP* ，处理函数先将 *counter* 自增并返回响应；
*main* 函数中第 *19-22* 行，申明并初始化 *http.Server* ，指定监听端口以及请求处理器；
第 *24* 调用 *ListenAndServe* 方法开始监听并处理网络请求。

接着启动计数器服务：

.. code-block:: shell-session

    $ go run counter-server.go

通过 *8080* 端口即可访问该服务：

.. code-block:: shell-session

    $ curl http://localhost:8080
    1
    $ curl http://localhost:8080
    2

这是一个非常简单的程序，但不失为一个完整的 `HTTP` 服务。
在 `net/http`_ 包的协助下，若干行代码即可实现 `HTTP` 服务！

数据交换
========

开发 `HTTP` 服务，不可避免地要在客户端和服务端之间 **交换数据** 。
交换数据的形式非常多样，至少包括以下类型：

- `URL` 值；
- `URL` 参数；
- `HTTP` 头部；
- `Cookie` ；
- `POST` 数据；

下面是一个非常细致的示例程序，演示服务端如何 **获取请求信息** 以及如何往客户端 **响应数据** 。
代码结构与计数器服务非常类似：

.. literalinclude:: /_src/practices/http-server/echo-server.go
    :name: practices/http-server/echo-server.go
    :language: go
    :lines: 13-
    :linenos:

例子第 *14* 行向客户端返回响应头部 *X-Data* ，值为： *foo* ；
第 *17-22* 行为客户端设置 *Cookie* ；
第 *24* 行，设置返回状态码并开始响应头部(后续便不能再修改响应头部了)；
接着，以响应体的形式返回各种请求信息，以此演示其获取方式。

第 *28* 行，获取远端(对端)地址；
第 *32-38* 行，分别获取 **请求方法** 、 `URL` 、 **主机名** ( **域名** )、 **协议版本** 等信息；
第 *43-44* 行，获取 **请求头部** ；
第 *49* 行，读取 **请求体** 。

下一步
======

.. include:: /_fragments/next-step-to-wechat-mp.rst

参考文献
========

#. `http - The Go Programming Language <https://golang.org/pkg/net/http>`_

.. include:: /_fragments/wechat-reward.rst

.. include:: /_fragments/disqus.rst

.. _net/http: https://golang.org/pkg/net/http

.. comments
    comment something out below

