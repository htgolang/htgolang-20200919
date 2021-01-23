1. 完成告警存储
    定义
        ID
        Fingerprint
        AlertName Labels["alertname"]
        Status: firing/resolved
        StartsAt
        EndsAt

        Labels text (json格式字符)
        Annotations text (json格式字符串)

        CreatedAt
        UpdatedAt
        DeletedAt

    同一个instance, job, 同一个规则产生的告警Fingerprint相同


    新增规则：
        产生告警 alertmanger 会连续发送多次 只存储一条
        fingerprint status=firing 查询 => 查询到数据，跳过
        无 => 插入

    更新规则：
        恢复时更新, fingerprint firing状态=>resolved
        fingerprint status=firing => update(resolved)

2. 查询页面:
    告警名称查询 展示所有满足条件的告警
    table
