import React from 'react'

import styles from './Card.module.css'
import { Card, Button } from 'react-bootstrap'
import Image from 'next/image'
export default function MenuItem(props) {
    const val = [
        {
            path: "/book",
            title: "Books",
            imgPath: "/books.png",
            desc: "Find the latest books here"
        }
        ,
        {
            path: "/3d",
            title: "3d world",
            imgPath: "/3d.png",
            desc: "sell of buy your 3D models"
        },
        {
            path: "/appliance.jpg",
            title: "Home appliance",
            imgPath: "/appliance.png",
            desc: "Find your necessary home appliance here"
        },
        {
            path: "/architecture",
            title: "Books",
            imgPath: "/architecture.png",
            desc: "Get the best architectact here hire now"
        },
        {
            path: "/bedding",
            title: "Bedding and Bedclothes",
            imgPath: "/decoration.png",
            desc: "Looking for bedding cloth you are in the right place"
        },
        {
            path: "/handicrafts",
            title: "Handicrafts",
            imgPath: "/handcraft.png",
            desc: "Find the latest books here"
        },
    
    ]




    return (
        <div className={styles.grid}>
            { val.map((res,key) => {
                return(
                <a key={key} href={res.path} className={styles.card}>
                    
                    <div key={key} className="card__image-container">
                        <h3>{res.title} &rarr;</h3>
                        <Image src={res.imgPath} quality="85" width={400} height={400}/>
                    </div>
                    <p>{res.desc}</p>
                </a>
            
                )
             
            })

            }
        </div>
    )


}

