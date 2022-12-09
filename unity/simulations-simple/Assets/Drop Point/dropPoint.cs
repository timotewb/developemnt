using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class dropPoint : MonoBehaviour
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
                if (hit.transform.gameObject.layer == 3){
                    Instantiate(PointToPlace, hit.point, Quaternion.identity);
                }
            }
        }
    }
}
