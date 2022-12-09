using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//https://forum.unity.com/threads/placing-objects-with-a-mouse-click.66121/


public class placePoint : MonoBehaviour
{
    [SerializeField] Transform PointToPlace; 
    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        if(Input.GetButtonDown("Fire1"))
        {
            RaycastHit hit;
            Ray ray = GetComponent<Camera>().ScreenPointToRay(Input.mousePosition);
            
            if (Physics.Raycast(ray, out hit, Mathf.Infinity)) {
                Instantiate(PointToPlace, hit.point, Quaternion.identity);
            }
        }
    }
}