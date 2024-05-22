// Next.js API route support: https://nextjs.org/docs/api-routes/introduction


import axios  from 'axios'
import {setCookie} from 'nookies'
import { AuthToken } from '../../src/util/validator'

export default (req, res) => {
  const email = req.body.email
  const pass = req.body.password
  axios.post(
    "http://localhost:9000/login",
    {
        email:email,
        password:pass
    }
  ).then(r=>{
    console.log(r.data)
    const token = new AuthToken(r.data)
    setCookie({ res }, 'auth', r.data, { secure: process.env.NODE_ENV === 'devlopment', maxAge: token.expiresAt(), httpOnly: true, path: '/' }) 
    res.json(r.data)
  })



}

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