import React from 'react'
import { Navbar,MenuItem } from '../src/Components'
import styles from '../styles/Home.module.css'
const Home =() => {

    const Menu = [{title:'Books',
                ImgPath:"/books.jpeg", 
                desc:"Here you can see the books form various author"},
                {title:'Books',
                ImgPath:"/books.jpeg", 
                desc:"Here you can see the books form various author"},
                {title:'Books',
                ImgPath:"/books.jpeg", 
                desc:"Here you can see the books form various author"},
                {title:'Books',
                ImgPath:"/books.jpeg", 
                desc:"Here you can see the books form various author"},
                {title:'Books',
                ImgPath:"/books.jpeg", 
                desc:"Here you can see the books form various author"},
                {title:'Books',
                ImgPath:"/books.jpeg", 
                desc:"Here you can see the books form various author"}]

    return (
        <div>
           <Navbar/>            
            <MenuItem/>
        </div>    
    )
}

export default Home