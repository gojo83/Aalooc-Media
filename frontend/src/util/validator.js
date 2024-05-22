import jwt_decode from "jwt-decode"

export class AuthToken {
    

    constructor(token){
        if(token){
           this.token = jwt_decode(token) 
        }
        
    }


    expiresAt() {
        console.log(this.token)
        return new Date(this.token.exp *1000)
    }    

    isExpired() {
       let time =new Date().getTime()
       console.log(time/1000)
     
       let expire = time/1000 > this.token.exp
       return expire
    }

    static async storeToken(token)   {
        
    }
    getToken () {
        return this.token
    }
        
}