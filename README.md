# 答题（模仿百万英雄）
用来在公司年会上，由主持人在舞台上展示ppt提问，并用手机web页面调用本系统的api接口控制问答以及生成最后的排行榜。 

没有使用数据库， 只把标准答案和主持人账号使用gob编码后放在文件里。

可以直接通过swagger查看api列表。 
http://localhost:8080/swagger/


# 用户角色分为主持人和游戏玩家
- 主持人
站在舞台上用ppt展示问题 以及用手机控制答题进度

- 游戏玩家 
用手机打开游戏答卷页面，提交每一道题的答案
  
- 答题全部完成后，主持人调用生成榜单的api后即可统计出所有参与游戏的榜单。
