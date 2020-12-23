# md-notify(开发中)
按照文件夹监控markdown文档变化，提供列表，检索，markdown文档内容接口
因为要写一个博客系统替换之前的wordpress博客，原因是wordpress太重。
目标是需要一个不依赖数据库,简单运行即可，前后端分离，复制即迁移的轻量级博客系统。wordpress显然无法满足需求。要实现markdown文件变更实时更新前台浏览必须有一个对markdown文档统一管理的组件。也就是本项目的最终目标。

### 预计支持的特性
- 提供文档列表接口(按序号)
- 提供文档列表接口(按类型)
- 提供分类列表
- 提供文档内容接口
- 按照标签筛选接口

### 预计支持的markdown标签
markdown标签和内容使用[content]进行分隔，类似section,比如下面的结构
```shell script
title=1
sub_title=副标题
image=1.jpg
other=3
[content]
这里开始是markdown文档开始部分
```
默认支持的标签如下,尽量简短
|  标签   | 说明  |
|  ----  | ----  |
| title  | 标题 |
| class  | 分类 |
| sort  | 排序号 |
| class  | 分类 |

默认标签是必填的，当然标签允许自定义,所有自定义标签都会使用倒排索引优化。标签数量不做限制
