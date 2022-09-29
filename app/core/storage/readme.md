# storage:

- 统一存储服务

## 服务列表：

- ✅ 文件存储：
    - 小文件存储： 图片， 文本
    - 大文件存储： 视频， 音频
- ✅ KV存储
- ✅ CDN

## 应用场景:

- 网盘 : 海量文件
- 社交网站：海量图片
- 电商网站：海量商品图片
- 视频网站：海量视频文件

## 参考:

### 存储方案:

> minio:

- https://github.com/minio/minio
- https://www.minio.org.cn/
- http://docs.minio.org.cn/docs/
- http://docs.minio.org.cn/minio/baremetal/
- 私有云存储, 兼容 Amazon S3 API
- 替代 S3, HDFS
- Minio对象存储系统适用于大文件场景，海量小文件的场景下并不适合。
- [基于 Go 开源项目 MIMIO 的对象存储方案在探探的实践](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487119&idx=1&sn=6e09abb32392e015911be3a1d7f066e5)
- https://xie.infoq.cn/article/66ffc331f851f5873a3e1b2d3
- https://devopsman.cn/archives/minio-de-ying-yong-chang-jing
- https://zhuanlan.zhihu.com/p/259594073

> 网盘:

- https://github.com/haiwen/seafile

> 其他:

- https://github.com/seaweedfs/seaweedfs