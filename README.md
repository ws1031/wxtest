# 微信开放平台

## 微信公众号相关接口

### 发布文章（综合）

    weixin/libs/officialaccount.SendNewsArticle

**文章发布流程**
1. 上传文章中的图片，获取图片url
2. 将文章内容中的图片替换为微信的图片链接
3. 上传文章封面图片，获取media_id
4. 新建草稿
5. 群发文章
6. 轮询监控群发状态（异步执行）

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


### 群发文章，消息推送
[文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html)

    libs/officialaccount.SendNews

### 获取群发消息推送状态
[文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html)
    
    libs/officialaccount.GetMassStatus



## 通知消息格式记录

### 群发消息通知
```json
{
    "ToUserName":"gh_fdefd501c79b",
    "FromUserName":"o7fbK55MPSGavlS-7E2fRC7olytc",
    "CreateTime":"1659512227",
    "MsgType":"event",
    "Event":"MASSSENDJOBFINISH",
    "MsgID":"1000000001",
    "Status":"send success",
    "TotalCount":"2",
    "FilterCount":"2",
    "SentCount":"2",
    "ErrorCount":"0",
    "CopyrightCheckResult":{
        "Count":"0",
        "CheckState":"1"
    },
    "ArticleUrlResult":{
        "Count":"2",
        "ResultList":{
            "item":[
                {
                    "ArticleIdx":"1",
                    "ArticleUrl":"http://mp.weixin.qq.com/s/VvouAj7fXB09HRBthjnElg"
                },
                {
                    "ArticleIdx":"2",
                    "ArticleUrl":"http://mp.weixin.qq.com/s/aaK-3vNnCt_Icc_g9nKXaA"
                }
            ]
        }
    }
}
```


### 发布文章通知
```json
{
    "ToUserName":"gh_fdefd501c79b",
    "FromUserName":"o7fbK5z3gnga3PEVeFOiKUr9aHrU",
    "CreateTime":"1659493195",
    "MsgType":"event",
    "Event":"PUBLISHJOBFINISH",
    "PublishEventInfo":{
        "publish_id":"2247483661",
        "publish_status":"0",
        "article_id":"LqjAvfE1cpNDvZE83nNDR4AFkX9So1JCV6Jea1KItCVwJZ_wBMo2Mc1_gG-3GkpK",
        "article_detail":{
            "count":"2",
            "item":[
                {
                    "idx":"1",
                    "article_url":"http://mp.weixin.qq.com/s?__biz=Mzg4ODgzODM4MQ==&mid=2247483661&idx=1&sn=f8ce5b67a3b4511d8addca1b6f2af648&chksm=cff4408ff883c999fd43a48358887d925f475fc802d02a4118ab8eb7c7c119107406328f3d8e#rd"
                },
                {
                    "idx":"2",
                    "article_url":"http://mp.weixin.qq.com/s?__biz=Mzg4ODgzODM4MQ==&mid=2247483661&idx=2&sn=dbb8d76a9c5ddb339440662de34ea05a&chksm=cff4408ff883c999eab958858898e46d5985c175be2dd4eb53c3f55a58eff91c95e5480655e4#rd"
                }
            ]
        }
    }
}
```

### 文字消息
```json
{
    "ToUserName":"gh_fdefd501c79b",
    "FromUserName":"o7fbK58Nq4POkwq_kMGzAheED9P4",
    "CreateTime":"1659421865",
    "MsgType":"text",
    "Content":"你好呀",
    "MsgId":"23757208460886931"
}
```

### 图片消息
```json
{
  "ToUserName":"gh_fdefd501c79b",
  "FromUserName":"o7fbK58Nq4POkwq_kMGzAheED9P4",
  "CreateTime":"1659421883",
  "MsgType":"image",
  "PicUrl":"http://mmbiz.qpic.cn/mmbiz_jpg/P5eXRa2DySVUbnWO6DbJRoMibRrFgIvffNvk8h6CCRVTY1ia3EiczjIlhoKnZbjdf7a2UK7MicDUSjgDGSVvmf8ib4A/0",
  "MsgId":"23757208157957746",
  "MediaId":"Wl_TTIXqibH_46V_m8sJbkvwL9aUx0Uxo0NllxbuQVN06ShH_QZ5_VBO83g3YqSR"
}
```

### 关注公众号
```json
{
    "ToUserName": "gh_fdefd501c79b",
    "FromUserName": "o7fbK59QFqPxKppu19ftpd0mAIXw",
    "CreateTime": "1659511270",
    "MsgType": "event",
    "Event": "subscribe"
}
```