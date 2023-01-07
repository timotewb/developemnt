using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class lineRendererDrawLerp : MonoBehaviour
{

    public float posX = -9.0f;
    public float posY = -4.0f;
    public float posZ = 0.0f;

    public float posXEnd = 9.0f;
    public float posYEnd = 6.0f;
    public float posZEnd = 0.0f;

    public float drawSpeed = 0.5f;

    private LineRenderer lr;
    
    // starting value for the Lerp
    static float t = 0.0f;

    // Start is called before the first frame update
    void Start()
    {
        lr = gameObject.GetComponent<LineRenderer>();
        lr.positionCount = 2;
        lr.SetPosition(0, new Vector3(posX,posY,posZ));        
    }

    // Update is called once per frame
    void Update()
    {
        lr.SetPosition(1, new Vector3(Mathf.Lerp(posX, posXEnd, t),Mathf.Lerp(posY, posYEnd, t),0));
        
        t += drawSpeed * Time.deltaTime;

        // when line is completely drawn, undraw it.
        if (t > 1.0f)
        {
            float temp = posXEnd;
            posXEnd = posX;
            posX = temp;

            temp = posYEnd;
            posYEnd = posY;
            posY = temp;

            temp = posZEnd;
            posZEnd = posZ;
            posZ = temp;

            t = 0.0f;
        }
    }
}
