using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class clearSimulation : MonoBehaviour
{
    public GameObject[] pointObjects;
    public void ClearSimulation()
    {
        pointObjects = GameObject.FindGameObjectsWithTag("pointsTag");
        LineRenderer lr = transform.Find("LineRenderer").gameObject.GetComponent<LineRenderer>();
        lr.positionCount = 0;
        transform.Find("LineRenderer").gameObject.SetActive(false);

        foreach (GameObject p in pointObjects)
        {
            GameObject.Destroy(p);
        }
    }
}
