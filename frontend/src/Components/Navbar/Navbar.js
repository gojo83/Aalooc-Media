import React from 'react'
import cn from 'classnames'

import styles from './Navbar.module.css'

const Navbarhome = () => {
    return (
        <div className={styles.navbar}>
            <div className={cn(styles.navcontainer,styles.flex)}>

                <h1>Alocmedia</h1>
                <nav>
                    <ul>
                
                        <a href="/login">Login</a>
                    
                
                        <a href="/signin">Sign in</a>    
                            
                    </ul>
                </nav>
            </div>
        </div>
                )
}

export default Navbarhome