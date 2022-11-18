// this file is generate by tools,do not modify it.
using System;
using System.Collections.Generic;
using UnityEngine;

public class TBossConfig
{

    /// <summary>
    /// 编号
    /// </summary>
    public int Id { get; set; }

    /// <summary>
    /// 鱼类型
    /// </summary>
    public float FishType { get; set; }

    /// <summary>
    /// Boss类型
    /// </summary>
    public int BossType { get; set; }

    /// <summary>
    /// 攻击内丹音效Id
    /// </summary>
    public string NeiDanAudioId { get; set; }

    /// <summary>
    /// 死亡金币动效UI路径
    /// </summary>
    public string DeadAnimImgPath { get; set; }

    /// <summary>
    /// 出场音效名称
    /// </summary>
    public string AppearAudio { get; set; }

    /// <summary>
    /// 普通音效名称
    /// </summary>
    public string NormalAudio { get; set; }

    /// <summary>
    /// 死亡音效名称
    /// </summary>
    public string DeadAudio { get; set; }

    /// <summary>
    /// 金币场boss死亡后撕咬内丹/连击倍率范围
    /// </summary>
    public string CoinHitBet { get; set; }

    /// <summary>
    /// 金币场boss死亡后撕咬内丹/连击次数
    /// </summary>
    public string CoinHitCount { get; set; }

    /// <summary>
    /// 魔晶场boss死亡后撕咬内丹/连击倍率范围
    /// </summary>
    public string MagicHitBet { get; set; }

    /// <summary>
    /// 魔晶场boss死亡后撕咬内丹/连击次数
    /// </summary>
    public string MagicHitCount { get; set; }

    /// <summary>
    /// 锁定位置偏移
    /// </summary>
    public string LockOffset { get; set; }

    /// <summary>
    /// 播放行走动作时角度偏移
    /// </summary>
    public string MoveAngle { get; set; }

    /// <summary>
    /// 移除场地boss旋转角度
    /// </summary>
    public string OutSceneAngle { get; set; }

    /// <summary>
    /// 移动速度
    /// </summary>
    public float MoveSpeed { get; set; }

    /// <summary>
    /// 播放表演动作路径节点
    /// </summary>
    public string StandbyIdx { get; set; }

    /// <summary>
    /// 行走路径名称
    /// </summary>
    public string MovePathName { get; set; }

    /// <summary>
    /// Boss缓存名称
    /// </summary>
    public string CacheBossName { get; set; }

    /// <summary>
    /// Boss预制体路径
    /// </summary>
    public string BossPrefabPath { get; set; }

    /// <summary>
    /// 攻击内丹timeline预制缓存名称
    /// </summary>
    public string CacheNeiDanName { get; set; }

    /// <summary>
    /// 攻击内丹timeline预制路径
    /// </summary>
    public string NeiDanTimeline { get; set; }

    /// <summary>
    /// 出场动作名称
    /// </summary>
    public string AppearAnimName { get; set; }

    /// <summary>
    /// 表演动作名称
    /// </summary>
    public string StandbyAnim { get; set; }

    /// <summary>
    /// 表演动作特效名称
    /// </summary>
    public string EffStandbyName { get; set; }

    /// <summary>
    /// 播放表演动作时长
    /// </summary>
    public float EffStandbyDuration { get; set; }

    /// <summary>
    /// 离场动作名称
    /// </summary>
    public string PullOutAnim { get; set; }

    /// <summary>
    /// 离场动作特效名称
    /// </summary>
    public string EffWalkoffName { get; set; }

    /// <summary>
    /// 是否为大boss(1大boss,0小boss,2盘古)(拥有出场timeline均为大boss)
    /// </summary>
    public int IsBigBoss { get; set; }

    /// <summary>
    /// 死亡后攻击间隔
    /// </summary>
    public float AttackInterval { get; set; }

    /// <summary>
    /// 死亡n秒后播放爆金UI
    /// </summary>
    public float DeadShowBombUI { get; set; }
}

public static class TBossConfigHelper
{
    private static List<TBossConfig> DataList;

    public static void InitData(string jsonStr)
    {
        DataList = LitJson.JsonMapper.ToObject<List<TBossConfig>>(jsonStr);
        if (DataList == null || DataList.Count == 0)
        {
            Debug.LogError("反序列化异常");
        }
    }

    public static List<TBossConfig> GetAll()
    {
        return DataList;
    }

    public static TBossConfig GetById(int id)
    {
        var info = GetByCondition(x => x.Id == id);
        if (info == null || info.Count == 0)
        {
            return null;
        }

        return info[0];
    }

    public static List<TBossConfig> GetByCondition(Predicate<TBossConfig> predicate)
    {
        return DataList.FindAll(predicate);
    }

    public static TBossConfig GetOneByCondition(Predicate<TBossConfig> predicate)
    {
        var temp = GetByCondition(predicate);
        if (temp == null || temp.Count == 0)
        {
            return null;
        }

        return temp[0];
    }
}