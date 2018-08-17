一款简易的本地代理工具
使用方式:
{
 "port":":8080",//本地端口号
 "target":"http://xxx/api",//代理地址
 "proxy":"/api",//拦截url
 "static":"xxx/static" //本地资源路径
 "dev":true  //开发环境
}
目前仅支持配置方式启动，请将config.json与程序放在同目录
