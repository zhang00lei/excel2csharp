using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public partial class DataTableManager : MonoBehaviour
{
    private void Start()
    {
        InitData();
        List<TTestDataTable> info = TTestDataTableHelper.GetAll();
        for (var i = 0; i < info.Count; i++)
        {
            Debug.Log(info[i]);
        }
    }
}