注意点:

Param
参数，表示需要传递到服务器端的参数，有五列参数，使用空格或者 tab 分割，五个分别表示的含义如下
参数名
参数类型，可以有的值是 formData、query、path、body、header
formData 表示是 post 请求的数据，
query 表示带在 url 之后的参数
path 表示请求路径上得参数，
body 表示是一个 raw 数据请求
header 表示带在 header 信息中得参数。
获取参数的方式与gin框架获取参数方式一样
参数类型
是否必须
注释
@Success
成功返回给客户端的信息，三个参数，
第一个是 status code。
第二个参数是返回的类型，必须使用 {} 包含，
第三个是返回的对象或者字符串信息
@router
路由信息，包含两个参数，使用空格分隔.
第一个是请求的路由地址，支持正则和自定义路由
第二个参数是支持的请求方法,放在 [] 之中，如果有多个方法，那么使用 , 分隔。
@Failure
失败返回的信息，包含两个参数，使用空格分隔，
第一个表示 status code
第二个表示错误信息
@Accept ----------------Content-type
代表的是http实体首部字段,接收的类型
application/x-www-form-urlencoded
1）浏览器的原生form表单
2） 提交的数据按照 key1=val1&key2=val2 的方式进行编码，key和val都进行了URL转码
multipart/form-data
常见的 POST 数据提交的方式。
application/json
消息主体是序列化后的 JSON 字符串,这个类型越来越多地被大家所使用
text/xml
是一种使用 HTTP 作为传输协议，XML 作为编码方式的远程调用规范
@Produce application/json
指定返回的内容类型
————————————————
