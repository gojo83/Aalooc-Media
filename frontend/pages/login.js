import React, { useState } from 'react'
import { TextField, colors, Button } from "@material-ui/core";
import styles from '../styles/Login.module.css'
import axios from 'axios'
import Router from 'next/router'
const Login = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState(false)
    const formSubmit = () => {
        const headers = {
            'Content-Type': 'text/plain'
        };
        axios.post(
            "http://localhost:9000/login",
            {
                email:email,
                password:password
            }
        ).then((r)=>{
            console.log(r.data)
            const result = r.data.msg
            if(result === "invalid user"){
              localStorage.setItem('authTrue' , "false")
              setError(true)
            }
            else{
              setError(false)
              localStorage.setItem('authTrue' , "true")
              if (localStorage.getItem('authTrue') == "true"){
                  Router.push('/home')
              }
            }
            
          })
    
     }
    
  

     const formSubmitNext = () => {
        const headers = {
            'Content-Type': 'text/plain'
        };
        axios.post(
            "http://localhost:3000/api/hello",
            {
                email:email,
                password:password
            }
        ).then((r)=>{
            console.log(r.data)
            const result = r.data.msg
            if(result === "invalid user"){
              localStorage.setItem('authTrue' , "false")
              setError(true)
            }
            else{
              setError(false)
              localStorage.setItem('authTrue' , "true")
              if (localStorage.getItem('authTrue') == "true"){
                  Router.push('/home')
              }
            }
            
          })
    
     }

    return ( 
        <div>
        
            <form className={styles.login} onSubmit={formSubmitNext}>
                <div
                    style={{
                        alignContent: "center",
                        color: "white",
                        fontStyle: "bold",
                    }}
                >
                    {" "}
          Welcome back!
        </div>
                {
                    error ?
                        <div className={styles.error}>
                            Invalid user cradentials
        </div> : <div></div>
                }

                <br></br>
                <TextField
                    className={styles.input}
                    id="outlined-secondary"
                    label="Outlined secondary"
                    variant="outlined"
                    label="Email"
                    type="text"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <br></br>
                <br></br>

                <TextField
                    className={styles.input}
                    id="outlined-secondary"
                    label="Outlined secondary"
                    variant="outlined"
                    label="Password"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <br></br>
                <br></br>
                <Button variant="contained" style={{ width: '400px', height: '50px', backgroundColor: '#3483eb', color: 'white' }}
                    onClick={formSubmitNext}>
                    Login
        </Button>
            </form>
        </div>
    )
}

export default Login;