// this file is generate by tools,do not modify it.
using UnityEngine;

public partial class DataTableManager
{
    private TextAsset textAsset;

    private void InitData()
    {
        textAsset = Resources.Load<TextAsset>("Json/TestDataTable");
        TTestDataTableHelper.InitData(textAsset.text);
    }
}