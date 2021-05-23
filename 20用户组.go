linux下可以创建多个用户，可以用组进行管理用户。

创建组 （在root用户下操作）
groupadd boys

创建用户
useradd -m -g boys ming
其中，-g表示在添加用户，同时将用户加到boys组。

修改现有用户到新的组
usermod -g boys a
其中， usermod表示修改用户信息。

如何查看用户和组？（不常用）
cat /etc/group
每一行表示一个group的信息，名称+ID

如何查看用户列表？
cat /etc/passwd
每一行表示一个用户的信息

要点：
一般不常用，如果有用group,就2个，两个人维护。

默认地，会给a1用户建立一个同名的组a1，也就是说这个组里只有他一个人。
