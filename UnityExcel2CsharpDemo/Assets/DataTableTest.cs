using System.Collections;
using System.Collections.Generic;
using UnityEditorInternal.VersionControl;
using UnityEngine;

public class TestInfo
{
    public int IntInfo;
    public List<int> ListInt;

    public float FloatInfo;
    public List<float> ListFloat;

    public bool BoolInfo;
    public List<bool> ListBool;

    public Vector2 Vector2Info;
    public List<Vector2> ListVector2;

    public Vector3 Vector3Info;
    public List<Vector3> ListVector3;

    public Color32 Color32Info;

    public Dictionary<float, int> testIntDict;
}

public class DataTableTest : MonoBehaviour
{
    void Start()
    {
        TestInfo testInfo = new TestInfo();
        testInfo.IntInfo = 1;
        testInfo.ListInt = new List<int>() {1, 11, 111, 1111};

        testInfo.FloatInfo = 1.1f;
        testInfo.ListFloat = new List<float>() {1.1f, 1.11f, 1.111f, 1.1111f};

        testInfo.BoolInfo = false;
        testInfo.ListBool = new List<bool>() {false, true, false, true};

        testInfo.Vector2Info = new Vector2(1.2f, 1);
        testInfo.ListVector2 = new List<Vector2>()
        {
            new Vector2(1.1f, 1.2f),
            new Vector2(2.1f, 2.2f),
            new Vector2(3.1f, 3.2f),
        };

        testInfo.ListVector3 = new List<Vector3>()
        {
            new Vector3(1.1f, 1.2f, 1.3f),
            new Vector3(2.1f, 2.2f, 2.3f),
            new Vector3(3.1f, 3.2f, 3.3f),
        };

        // testInfo.Color32Info = new Color32(1, 1, 1, 1);

        testInfo.testIntDict = new Dictionary<float, int>()
        {
            {1.2f, 1}, {2, 2}
        };

        string str = LitJson.JsonMapper.ToJson(testInfo);
        Debug.Log(str);
    }
}