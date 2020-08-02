import React from 'react'

export const Navbar = () => {
    return (

        <header className="navbar">
            <ul className="navbar-items">
                <li className="home"><a href="/">Home</a></li>
                <li className="history"><a href="#history">History</a></li>
                <li className="login"><a href="#login">login</a></li>
                <li className="signup"><a href="#signup">signup</a></li>

            </ul>
        </header>        
      )
}