// this file is generate by tools,do not modify it.
using System;
using System.Collections.Generic;
using UnityEngine;

public class TTestDataTable
{

    /// <summary>
    /// id
    /// </summary>
    public int Id { get; set; }

    /// <summary>
    /// 字符串
    /// </summary>
    public string StrInfo { get; set; }

    /// <summary>
    /// 整形
    /// </summary>
    public int IntInfo { get; set; }

    /// <summary>
    /// 浮点型
    /// </summary>
    public float FloatInfo { get; set; }

    /// <summary>
    /// 布尔型
    /// </summary>
    public bool BoolInfo { get; set; }
}

public static class TTestDataTableHelper
{
    private static List<TTestDataTable> DataList;

    public static void InitData(string jsonStr)
    {
        DataList = LitJson.JsonMapper.ToObject<List<TTestDataTable>>(jsonStr);
        if (DataList == null || DataList.Count == 0)
        {
            Debug.LogError("反序列化异常");
        }
    }

    public static List<TTestDataTable> GetAll()
    {
        return DataList;
    }

    public static TTestDataTable GetById(int id)
    {
        var info = GetByCondition(x => x.Id == id);
        if (info == null || info.Count == 0)
        {
            return null;
        }

        return info[0];
    }

    public static List<TTestDataTable> GetByCondition(Predicate<TTestDataTable> predicate)
    {
        return DataList.FindAll(predicate);
    }

    public static TTestDataTable GetOneByCondition(Predicate<TTestDataTable> predicate)
    {
        var temp = GetByCondition(predicate);
        if (temp == null || temp.Count == 0)
        {
            return null;
        }

        return temp[0];
    }
}