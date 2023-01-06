using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class runSimulation : MonoBehaviour
{
    public GameObject[] pointObjects;
    // private LineRenderer lr;
    // public int lengthOfLineRenderer = 20;
    // LineRenderer lineRenderer = gameObject.AddComponent<LineRenderer>();
    // lineRenderer.material = new Material(Shader.Find("Sprites/Default"));
    // lineRenderer.widthMultiplier = 0.2f;
    // lineRenderer.positionCount = lengthOfLineRenderer;
    // Start is called before the first frame update
    public void RunSimulation()
    {
        // lr = GetComponent<LineRenderer>();
        pointObjects = GameObject.FindGameObjectsWithTag("pointsTag");
        Debug.Log(pointObjects.Length);
        transform.Find("LineRenderer").gameObject.SetActive(true);
        LineRenderer lr = transform.Find("LineRenderer").gameObject.GetComponent<LineRenderer>();
        lr.positionCount = pointObjects.Length;

        // foreach (GameObject p in pointObjects)
        // {
        //     Debug.Log("Found 1.");
        // }
        for(int i = 1; i<pointObjects.Length; i++)
        {
            Debug.Log(i);
            lr.SetPosition(0, new Vector3(pointObjects[0].transform.position.x,pointObjects[0].transform.position.y,pointObjects[0].transform.position.z));
            lr.SetPosition(i, new Vector3(pointObjects[i].transform.position.x,pointObjects[i].transform.position.y,pointObjects[i].transform.position.z));
        }
    }
}
