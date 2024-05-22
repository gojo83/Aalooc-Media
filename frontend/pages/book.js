import React from 'react'
import axios from 'axios'
const Book=({d})=>{



    return (
        <div>
          {  d.map(r=>{
              return(
                <p>{r.name}</p>
              )
                 
            })
        }
        </div>    
    )
}


export async function getServerSideProps(ctx){
   var d =""
   await axios.get(
        "http://localhost:9090/products",
   ).then(r=>{
        console.log(r.data)
         d = r.data
       
    })
   
     return { props:{d}}
}

export default Book;