import React from 'react'
import styles from './sidebar.module.css'
import cn from 'classnames'
import {FaBriefcase,FaAdn , FaCartPlus,FaAlignJustify} from 'react-icons/fa'
import {BsFillBriefcaseFill} from 'react-icons/bs'
import {CgFileAdd} from 'react-icons/cg'
import {AiOutlineSetting} from 'react-icons/ai'
export default function SideBar() {
    return (
        <div>
        <div className={styles.sidebar}>
            <div className={cn(styles.sidebar, styles.flex)}>
              
                <nav>
                    <ul>
                        <li>
                        <FaAdn className={styles.icon}/>    
                        <h1>Alocmedia</h1>
                        </li>    
                        <li>
                            <CgFileAdd className={styles.icon}/>
                            Create post
                        </li>
                        <li>
                            <FaBriefcase className={styles.icon}/>
                            Products
                        </li>
                        <li>
                            <FaCartPlus className={styles.icon}/>
                            Orders
                       </li>
                        <li>
                            <AiOutlineSetting className={styles.icon}/>
                            Settings
                    </li>
                    </ul>
                </nav>
            </div>
           
        </div> 
        <div >
        <FaAlignJustify className={styles.toogle}/>
        </div>
        </div>
    )
}