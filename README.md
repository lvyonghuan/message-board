# message-board
    这算蒟蒻第一次写这么大的工程了。目前已实现的功能有：
    用户的注册，登陆，修改密码,密码的哈希加密，cookie的加密。
    留言板留言，查看我相关的留言，删除自己的留言，查看所有公开留言。
    未来计划实现的：私密留言，查看我的发言，回复，管理员权限（暂时定位为删除所有留言的权限）等等。
# 使用的技术
    该项目使用了gin框架与mysql驱动。
    顺便利用了ChatGPT这个工具帮助了部分代码的完善。我不知道大家怎么看这件事，反正ChatCPT给我的代码修改带来了很大灵感，让我知道了密码能用哈希加密和cookie的加密手段。以及，我不用苦苦对着百度到的一堆格式都没有的代码看半天了。
    在此单方面宣布ChatGPT是新手教程之王（逃。在这里就不长篇大论工具与人的关系了，这个问题建议丢到知乎去发愁。反正AI消灭不掉人的创作表达欲望，消灭的是简单重复的流水线工作罢了。
# 数据库表结构
    存放在数据库表结构文件中。
# 亮点（持续更新）
1.0

    能够给指定用户发送信息。
    检查用户的登录状态。

1.1

    游客在非登录状态下能够查看留言板中的所有公开留言。

1.2（算是一次安全方面的更新）

    实现了用户密码的哈希加密。
    实现了对cookie的加密。
    增加了一个计时器。故预计以后可以加入更多与时间相关的功能。