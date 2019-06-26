#### 程序简介:
	服务端： beego写的restapi服务，简称 restapi
	db：mysql 数据库 简称：restapi-mariadb


#### 准备工作：
	1. restapi 通过deployment部署；
	2. restapi-mariadb 通过有状态的StatefulSet 部署；
	3. db 的用户名密码通过secret配置；
	4. restapi-mariadb 使用本地存储持久化；
	5. restapi 根据 service 与 restapi-mariadb 通信，并通过service实现内部的负载均衡；
	
	这些准备工作前面都有教程，不熟悉的可以去看下；

#### 准备yml 文件
	1. restapi 部署文件 https://github.com/xiliangMa/restapi/blob/master/deployment.yml)
	2. restapi-mariadb 部署文件 (https://github.com/xiliangMa/restapi/blob/master/statefulset.yml)
	3. restapi-mariadb 密文 (https://github.com/xiliangMa/restapi/blob/master/restapi-mariadb-secret.yml)
	4. restapi service 部署文件 (https://github.com/xiliangMa/restapi/blob/master/service.yml)
	5. restapi-mariadb service 部署文件 (https://github.com/xiliangMa/restapi/blob/master/service-mariadb.yml)


#### 部署:
	1. 部署restapi 
	2. 部署restapi-mariadb
	3. 部署restapi-mariadb 密文
	4. 部署 restapi service
	5. 部署 restapi-mariadb  service
	
	

 ##### 通过rancher 部署： 喜欢图形化操作、不熟悉k8s 命令的就去搭建一个rancher部署； 	
 ##### 通过kubectl 部署： kubectl  apply -f 相关的文件

##### restapi:
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190626160652599.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MTgwNjI0NQ==,size_16,color_FFFFFF,t_70)
	
##### restapi-mariadb:

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190626160958667.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MTgwNjI0NQ==,size_16,color_FFFFFF,t_70)
##### service:

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190626160908165.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MTgwNjI0NQ==,size_16,color_FFFFFF,t_70)

##### 测试结果
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190626161137687.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MTgwNjI0NQ==,size_16,color_FFFFFF,t_70)


##### 总结：

自研发应用程序如何部署在k8s集群：
1. Dockerfile 编写将程序打包成docker 镜像；
2. 熟悉k8s 各个控制器 deployment、Statefulset， DaemonSet应用场景；
3. 熟悉service、ingress 作用应用场景；
4. 熟悉了第三点后才能通过k8s实现基础的负载均衡；后面会将高大上的**服务网格**；

好了最近有点忙，先到这里；