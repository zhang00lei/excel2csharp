// this file is generate by tools,do not modify it.
using System;
using System.Collections.Generic;
using UnityEngine;

public class TBossConfig
{

    /// <summary>
    /// id
    /// </summary>
    public int Id { get; set; }

    /// <summary>
    /// 整形
    /// </summary>
    public int IntInfo { get; set; }

    /// <summary>
    /// 整形集合
    /// </summary>
    public string ListInt { get; set; }

    /// <summary>
    /// 浮点型
    /// </summary>
    public float FloatInfo { get; set; }

    /// <summary>
    /// 浮点型集合
    /// </summary>
    public string ListFloat { get; set; }

    /// <summary>
    /// 布尔型
    /// </summary>
    public string BoolInfo { get; set; }

    /// <summary>
    /// 布尔集合
    /// </summary>
    public string ListBool { get; set; }
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