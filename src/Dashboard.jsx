// import { Menubar } from "primereact/menubar";
// import SparkLogo from "./assets/sparkhacks-logo.svg";
// import { useEffect, useState } from "react";
// import Insta from "./Insta.jsx";
// import Twitter from "./Twitter.jsx";
// import Stats from "./Stats.jsx";
// import { Button } from "primereact/button";
// import { Dialog } from "primereact/dialog";

// function Dashboard() {
//   const postInsta = () => {};
//   const [visible, setVisible] = useState(false);

//   const postTweet = () => {};

//   const travelDestinations = [
//     "TRENDING TAGS:",
//     "#business",
//     "#digitalmarketing",
//     "#entrepreneur",
//     "#onlinebusiness",
//     "#marketing",
//     "#store",
//     "#ecommerce",
//     "#smallbusiness",
//     "#shopping",
//     "#sales",
//   ];
//   return (
//     <>
//       <div className="navbar">
//         <img id="logo" src={SparkLogo} />
//         <h1>MyBusinessBoard : Social Media Manager</h1>
//       </div>
//       <div className="body-container">
//         <div className="column-container">
//           <div className="post-container">
//             <div className="posts">
//               {" "}
//               <Insta />
//             </div>
//             <Button
//               className="post-button"
//               type="button"
//               onClick={() => setVisible(true)}
//             >
//               Create Instagram Post ⊕
//             </Button>
//             <Dialog
//               header="Create an Instagram Post"
//               visible={visible}
//               draggable={false}
//               className="createInstaPost"
//               modal={false}
//               style={{ width: "50vw" }}
//               onHide={() => {
//                 if (!visible) return;
//                 setVisible(false);
//               }}
//             >
//               <p className="m-0">
//                 <form>

//                 </form>
//               </p>
//             </Dialog>
//           </div>
//           <div className="post-container">
//             <div className="posts">
//               {" "}
//               <Twitter />
//             </div>
//             <Button className="post-button" type="button">
//               Create Twitter Post ⊕
//             </Button>
//           </div>
//           <div className="stats">
//             <Stats />
//           </div>
//         </div>
//         <div className="ticker-bar">
//           <div className="news-ticker">
//             <div className="ticker-content">
//               {travelDestinations.map((country, index) => (
//                 <span key={index} className="ticker-item">
//                   {country} &nbsp;&nbsp;
//                 </span>
//               ))}
//             </div>
//           </div>
//         </div>{" "}
//       </div>
//     </>
//   );
// }

// export default Dashboard;

import { Menubar } from "primereact/menubar";
import SparkLogo from "./assets/sparkhacks-logo.svg";
import { useState } from "react";
import Insta from "./Insta.jsx";
import Twitter from "./Twitter.jsx";
import Stats from "./Stats.jsx";
import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";

function Dashboard() {
  const [instaVisible, setInstaVisible] = useState(false);
  const [twitterVisible, setTwitterVisible] = useState(false);

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
          {/* Instagram Post Section */}
          <div className="post-container">
            <div className="posts">
              <Insta />
            </div>
            <Button
              className="post-button"
              type="button"
              onClick={() => setInstaVisible(true)}
            >
              Create Instagram Post ⊕
            </Button>
            <Dialog
              header="Create an Instagram Post"
              visible={instaVisible}
              draggable={false}
              className="createPost"
              modal={true}
              style={{ width: "50vw" }}
              onHide={() => setInstaVisible(false)}
            >
              <form>
                <p>Instagram post form goes here</p>
              </form>
            </Dialog>
          </div>

          {/* Twitter Post Section */}
          <div className="post-container">
            <div className="posts">
              <Twitter />
            </div>
            <Button
              className="post-button"
              type="button"
              onClick={() => setTwitterVisible(true)}
            >
              Create Twitter Post ⊕
            </Button>
            <Dialog
              header="Create a Twitter Post"
              visible={twitterVisible}
              draggable={false}
              className="createPost"
              modal={true}
              style={{ width: "50vw" }}
              onHide={() => setTwitterVisible(false)}
            >
              <form>
                <p>Twitter post form goes here</p>
              </form>
            </Dialog>
          </div>

          <div className="stats">
            <Stats />
          </div>
        </div>

        <div className="ticker-bar">
          <div className="news-ticker">
            <div className="ticker-content">
              {travelDestinations.map((tag, index) => (
                <span key={index} className="ticker-item">
                  {tag} &nbsp;&nbsp;
                </span>
              ))}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default Dashboard;
