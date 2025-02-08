import { Menubar } from "primereact/menubar";
import SparkLogo from "./assets/sparkhacks-logo.svg";
import { useEffect, useState } from "react";
import Insta from "./Insta.jsx";
import Twitter from "./Twitter.jsx";
import Stats from "./Stats.jsx";
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
        

function Dashboard() {

  const postInsta = () => {
  
  };

  const postTweet = () => {

  };
  
  const travelDestinations = [
    "TRENDING TAGS:",
    "#business",
    "#digitalmarketing",
    "#entrepreneur",
    "#onlinebusiness",
    "#marketing",
    "#store",
    "#ecommerce",
    "#smallbusiness",
    "#shopping",
    "#sales",
  ];
  return (
    <>
      <div className="navbar">
        <img id="logo" src={SparkLogo} />
        <h1>MyBusinessBoard : Social Media Manager</h1>
      </div>
      <div className="body-container">
        <div className="column-container">
          <div className="post-container">
            <div className="posts">
              {" "}
              <Insta />
            </div>
            <Button className="post-button" type="button">Create Instagram Post ⊕</Button>
          </div>
          <div className="post-container">
            <div className="posts">
              {" "}
              <Twitter />
            </div>
            <Button className="post-button" type="button">Create Twitter Post ⊕</Button>
          </div>
          <div className="stats">
            <Stats />
          </div>
        </div>
        <div className="ticker-bar">
          <div className="news-ticker">
            <div className="ticker-content">
              {travelDestinations.map((country, index) => (
                <span key={index} className="ticker-item">
                  {country} &nbsp;&nbsp;
                </span>
              ))}
            </div>
          </div>
        </div>{" "}
      </div>
    </>
  );
}

export default Dashboard;
