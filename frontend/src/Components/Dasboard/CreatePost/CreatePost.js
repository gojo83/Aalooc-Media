import { StylesProvider } from '@material-ui/core'
import React , {useState} from 'react'
import Styles from './CreatePost.module.css'
import axios from 'axios'

import { Divider, Form, Label,Button, Checkbox } from "semantic-ui-react";
import Notification from '../../Notification/Notification';

const CreatePost  = (props) => {

    const [title , setTitle] = useState('')
    const [description, setDescription] = useState('')
    const [price , setPrice] = useState('')
    const [storeName ,setStore] = useState('')
    const [location , setLocation] = useState('')
    const [thumbnail, setThumbnail] = useState('')
    const [images , setImage] = useState([])
    const [catagory,setCatagory] = useState('')
    const [files , setFiles] = useState('')
    const [showNotification , setShownotification] = useState(false)
    const imageList = []



    const submit =()=>{
      axios.post("http://localhost:9090/addproduct" , {
        
          "title":title,
          "description":description,
          "price":price,
          "userId":props.info.userId,
          "store_name":storeName,
          "location":location,
          "thumbnail":imageList[0],
          "images":imageList,
          "catagory":catagory
      }
      ).then(res=>{
          if(res){
           setShownotification(true)
          }
      })
    }

    const uploadFiles=()=>{
      for (let i =0 ; i < files.length;i++){
        if(imageList.indexOf(files[i].name)===-1){
          imageList.push(files[i].name)
        }
        console.log(images)
        upload(files[i])
      }
    }

    const upload =(file)=>{
      let id = 48
      const data  = new FormData()
           data.append('id' , id)
           data.append('file' , file) 
      axios.post("http://localhost:9001/",
      data).then(r=>{
        console.log(r.data)
      })
    }  
    return (
      <div className={Styles.container}>
       <h1>{props.info.userId}</h1>
       {showNotification?
       <Notification/>:'' 
       }
       
     <Form >
            <Form.Field inline>
              <input
                className={Styles.title}
  
                type="text"
                placeholder="Title"
                name="title"
                value={title}
                onChange={(e)=>setTitle(e.target.value)}
              />
            </Form.Field>
            <br></br>

            <Form.Field inline>
              <textarea
                className={Styles.desc}
                type="text"
                placeholder="Description"
                name="description"
                value={description}
                onChange={(e)=>setDescription(e.target.value)}
              />
             
            </Form.Field>
            <Form.Field inline>
              <input
                className={Styles.title}
  
                type="text"
                placeholder="price"
                name="price"
                value={price}
                onChange={(e)=>setPrice(e.target.value)}
              />
            </Form.Field>
            <Form.Field inline>
              <input
                className={Styles.title}
  
                type="text"
                placeholder="Location"
                name="location"
                value={location}
                onChange={(e)=>setLocation(e.target.value)}
              />
            </Form.Field>
            <Form.Field inline>
              <input
                className={Styles.title}
  
                type="text"
                placeholder="Catagory"
                name="catagory"
                value={catagory}
                onChange={(e)=>setCatagory(e.target.value)}
              />
            </Form.Field>  
            <label className={Styles.customfileupload}>
            <input style={{display:"none"}} type="file" name="file" multiple onChange={ e=> setFiles(e.target.files)}/>
            Choose Files
            </label>
            <br></br>
            <button className={Styles.upload} onClick={uploadFiles}>Upload</button>
            <br></br>
            <button onClick={submit}>submit</button>
          </Form>
      </div>    
    )
}

export default CreatePost