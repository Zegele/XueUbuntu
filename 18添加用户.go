课程18 添加用户

1.添加用户
sudo useradd -m test1
其中，
sudo,表示以管理员身份执行。
-m, 表示在/home下添加用户目录

2.修改用户密码
sudo passwd test1

3. 删除用户
sudo userdel test1

然后
sudo rm -rf /home/test1
才会清理干净，整个test1用户才会完全删除。

要点：
1.在登录系统时，默认不允许以root用户登录。
2. 只有特殊的用户，才能执行sudo
比如，初始用户可以执行sudo，但test1不行。

Linux下，把能执行sudo命令的用户叫sudoer。

