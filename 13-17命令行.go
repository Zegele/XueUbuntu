课程13-16 命令行使用


ls: 即list，列出目录下的所有项
ls /地址 可以查看该地址内的所有项
ls -l 以详细方式显示所有项，

cd：打开文件夹(change directory)切换目录
几个特殊的目录
~ 代表当前用户的主目录 
cd ~ 切换到主目录
cd ~/example 切换到主目录下的example目录

. 代表当前目录 		 


.. 代表上一级目录
cd ../www 切换到上级目录，再到www子目录。



pwd：显示当前位置（print working directory）

mkdir：新建文件 如：mkdir abc
mkdir -p abc/123/test
使用-p参数，可以将路径的层次目录全部创建。


rmdir：删除文件（remove directory）一般不使用这个
rmdir abc
如果目录非空，则删除失败。


cp：复制文件或者目录
cp -rf abc abc1 把abc复制为abc2

mv：移动文件或目录（重命名）
mv simple.txt simple2.txt  //将simple重命名为simple2



rm：remove删除文件或目录
rm -rf abc 删除abc目录，和子项一并删除。
其中，r表示recusive（递归删除）， f表示force（强制删除）
rm -rf /* 删库跑路
*/

对于文件， rm/cp/mv这三个命令同样适用



tar：归档
tape archive 档案打包
创建档案包
tar -crf example.tar example
(把example文件，打包成.tar档案包)
其中
c，表示create创建档案
v，表示verbose显示详情
f，表示file
也可以多个目录打包 
如： 
tar -cvf xxx.tar file1 file2 file3

还原档案包
归档
tar -xvf example.tar
默认解到当前目录下。
tar -xvf example.tar -C outdir
其中， -C参数制定目标目录

先前的tar格式并没有压缩，体积较大。
归档并压缩
tar -czvf example.tar.gz example
z, 代表压缩

解压缩
tar -xzvf example.tar.gz
tar -xzvf exampme.tar.gz -C outdir
z, 代表解压缩，这几个参数顺序无关。

通常我们所见的，都是*.tar.gz这种格式。


ln：软连接（相当于win的快捷方式）
ln -s source link
其中， -s表示soft软连接
source 是原始的文件名或目录名
link 是生成软链接的名字

比如
ln -s example example2
example2是example的软链接。



zip：

unzip：



输入命令和路径时，按Tab可以自动补全。

输入历史可以翻阅，按上，下键，即可。

宿主机和虚拟机之间可以拷贝粘贴。
一般情况下，文本和文件都可以拷贝。