import { Menubar } from "primereact/menubar";
import SparkLogo from "./assets/sparkhacks-logo.svg";
import { useState, useRef } from "react";
import Insta from "./Insta.jsx";
import Twitter from "./Twitter.jsx";
import Stats from "./Stats.jsx";
import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { FileUpload } from "primereact/fileupload";
import { Toast } from "primereact/toast";

function Dashboard() {
  const [instaVisible, setInstaVisible] = useState(false);
  const [twitterVisible, setTwitterVisible] = useState(false);
  const toast = useRef(null);
  const onUpload = () => {
    toast.current.show({
      severity: "info",
      summary: "Success",
      detail: "File Uploaded",
    });
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
                Enter a caption: <input></input>
                <div className="card flex justify-content-center">
            <Toast ref={toast}></Toast>
            <FileUpload mode="basic" name="demo[]" url="/api/upload" accept="image/*" maxFileSize={1000000} onUpload={onUpload} auto chooseLabel="Browse Files" />
        </div>  
        <Button type="Submit">Submit</Button>
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
                <div className="card flex justify-content-center">
                  <Toast ref={toast}></Toast>
                  <FileUpload
                    mode="basic"
                    name="demo[]"
                    url="/api/upload"
                    accept="image/*"
                    maxFileSize={1000000}
                    onUpload={onUpload}
                    auto
                    chooseLabel="Browse"
                  />
                </div>
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
