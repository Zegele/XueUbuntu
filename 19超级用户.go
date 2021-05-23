19 超级用户
超级用户root
类似于windows下的administrator用户



1. 首次使用时，需要给root设置密码
sudo passws root
先输入现用户密码，然后设置root用户密码

2. 切换到root用户
su root
需要输入root用户的密码

其中，su表示 switch user（切换用户）
切换到root用户，有全部权限，可以直接useradd等命令。


3. 退出
exit

要点
1. su root 仅仅对当前会话（终端）有效
不影响当前桌面环境

2. root 权力太大，需要小心使用。

