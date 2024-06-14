import React from "react";
import "./AppLayout.css";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
// import SearchIcon from "@mui/icons-material/Search";
import { Outlet } from "react-router-dom";

const AppLayout = () => {
  return (
    <div>
      <nav className="navbar">
        <div className="navbar-logo">
          <a href="/">
            <img src="img/logo1.webp" alt="Logo" />
          </a>
        </div>
        {/* <SearchIcon /> */}
        <ul className="navbar-list">
          <li className="navbar-item">
            <a href="favorite">
              <FavoriteBorderIcon />
            </a>
          </li>
          <li className="navbar-item special">
            <a href="profile">sign up</a>
          </li>
        </ul>
      </nav>
      <Outlet />
    </div>
  );
};

export default AppLayout;
