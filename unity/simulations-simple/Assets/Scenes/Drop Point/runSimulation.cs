using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;

public class runSimulation : MonoBehaviour
{
    public GameObject[] pointObjects;
    public float drawSpeed = 0.5f;
    private float t = 0.0f;
    private bool opening;
    private bool drawComplete;
    public TMP_Text _dist;
    private float dist;
    // private LineRenderer lr;
    // public int lengthOfLineRenderer = 20;
    // LineRenderer lineRenderer = gameObject.AddComponent<LineRenderer>();
    // lineRenderer.material = new Material(Shader.Find("Sprites/Default"));
    // lineRenderer.widthMultiplier = 0.2f;
    // lineRenderer.positionCount = lengthOfLineRenderer;
    // Start is called before the first frame update

    public void OpenDoor() {
        if (!opening) {
            StartCoroutine(RunSimulation());
        }
    }
    IEnumerator RunSimulation()
    {
        opening = true;
        // lr = GetComponent<LineRenderer>();
        pointObjects = GameObject.FindGameObjectsWithTag("pointsTag");
        Debug.Log(pointObjects.Length);
        transform.Find("LineRenderer").gameObject.SetActive(true);
        LineRenderer lr = transform.Find("LineRenderer").gameObject.GetComponent<LineRenderer>();
        lr.positionCount = 1;

        // foreach (GameObject p in pointObjects)
        // {
        //     Debug.Log("Found 1.");
        // }
        lr.SetPosition(0, new Vector3(pointObjects[0].transform.position.x, pointObjects[0].transform.position.y, pointObjects[0].transform.position.z));
        for (int i = 1; i<pointObjects.Length; i++)
        {
            Debug.Log(i);
            lr.positionCount = i+1;
            t = 0.0f;
            drawComplete = false;

            while (!drawComplete)
            {
                // transform.position = Vector3.Lerp(positionA , positionB,LerpTime); 
                if (t > 1.0f & !drawComplete)
                {
                    drawComplete = true;
                    t = 1.0f;
                }
                lr.SetPosition(i, new Vector3(Mathf.Lerp(pointObjects[i - 1].transform.position.x, pointObjects[i].transform.position.x, t), Mathf.Lerp(pointObjects[i - 1].transform.position.y, pointObjects[i].transform.position.y, t), pointObjects[0].transform.position.z));
                yield return new WaitForEndOfFrame();
                t += drawSpeed * Time.deltaTime;
            }
            dist = dist+Vector3.Distance(pointObjects[i - 1].transform.position, pointObjects[i].transform.position);
        }
        _dist.text = dist.ToString("F2")+"km";
        opening = false;
    }
}
