# 微信开放平台

## 微信公众号相关接口

### 发布文章（综合）

    weixin/libs/officialaccount.PublishArticle

**文章发布流程**
1. 上传文章中的图片，获取图片url
2. 将文章内容中的图片替换为微信的图片链接
3. 上传文章封面图片，获取media_id
4. 新建草稿
5. 发布文章
6. 轮询监控发布状态（异步执行）

### 上传图文消息内的图片获取URL
[文档](https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html)

本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。

    libs/officialaccount.MediaUploadImg

### 上传永久性素材（用于上传文章封面图片）
[文档](https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html)

    libs/officialaccount.MediaAddMaterial


### 新建草稿
[文档](https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html)

    libs/officialaccount.AddDraft


### 发布接口
[文档](https://developers.weixin.qq.com/doc/offiaccount/Publish/Publish.html)

    libs/officialaccount.Publish

### 发布状态轮询接口
[文档](https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_status.html)
    
    libs/officialaccount.PublishStatus
