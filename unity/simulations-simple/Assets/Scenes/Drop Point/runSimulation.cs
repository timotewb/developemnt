using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class runSimulation : MonoBehaviour
{
    public GameObject[] pointObjects;
    // Start is called before the first frame update
    public void RunSimulation()
    {
        pointObjects = GameObject.FindGameObjectsWithTag("pointsTag");
        Debug.Log(pointObjects.Length);

        // foreach (GameObject p in pointObjects)
        // {
        //     Debug.Log("Found 1.");
        // }
    }
}
