import React from "react";
import "./Navbar.css";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";

const Navbar = () => {
  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <img src="img/logo1.webp" alt="Logo" />
      </div>
      <ul className="navbar-list">
        <li className="navbar-item">
          <a href="#favorite">
            <FavoriteBorderIcon />
          </a>
        </li>
        <li className="navbar-item special">
          <a href="#profile">sign up</a>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;
