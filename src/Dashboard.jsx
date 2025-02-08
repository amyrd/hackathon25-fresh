import { Menubar } from "primereact/menubar";
import SparkLogo from "./assets/sparkhacks-logo.svg"
import { useEffect, useState } from "react";
import Insta from "./Insta.jsx";
import Twitter from "./Twitter.jsx";
import Stats from "./Stats.jsx";

function Dashboard() {
  return (
    <>
      <div className="navbar">
        <img
          id="logo"
          src={SparkLogo}
        />
        <h1>MyBusinessBoard : Social Media Manager</h1>
      </div>
      <div className="body-container">
        <div className="column-container">
          <div className="post-container">
            <div className="posts">
              {" "}
              <Insta />
            </div>
            <button className="post-button"></button>
          </div>
          <div className="post-container">
            <div className="posts">
              {" "}
              <Twitter />
            </div>
            <button className="post-button"></button>
          </div>
          <div className="stats">
            <Stats />
          </div>
        </div>
        <div className="row-container">Footer element</div>
      </div>
    </>
  );
}

export default Dashboard;
