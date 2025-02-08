import { Menubar } from "primereact/menubar";
import ReactImg from "./assets/react.svg";
import { useEffect, useState } from "react";
import Posts from "./Posts.jsx";

function Dashboard() {
  const [posts, setPosts] = useState([]);
  const [currentIndex, setCurrentIndex] = useState(0);

  const travelDestinations = [
    // "Japan #1",
    // "Italy #2",
    // "New Zealand #3",
    // "Canada #4",
    // "Switzerland #5",
    // "Greece #6",
    // "Australia #7",
    // "Thailand #8",
    // "France #9",
    // "Iceland #10",

    "#1 Japan",
    "#2 Italy",
    "#3 New Zealand",
    "#4 Canada",
    "#5 Switzerland",
    "#6 Greece",
    "#7 Australia",
    "#8 Thailand",
    "#9 France",
    "#10 Iceland",
  ]


  useEffect(() => {
    fetch("http://localhost:8080/twitter/posts") // Fetch data from Go backend
      .then((response) => response.json())
      .then((data) => setPosts(data))
      .catch((error) => console.error("Error fetching posts:", error));
  }, []);

  return (
    <>
      <div className="navbar">
        <img src={ReactImg} style={{ width: "40px", height: "40px" }} />
        <h1> MyBusinessBoard : Social Media Manager</h1>
      </div>
      <div className="body-container">
        <div className="column-container">
          <div className="post-container">
            <div className="posts">
              {" "}
              <Posts posts={posts} />
            </div>
            <button className="post-button"></button>
          </div>
          <div className="post-container">
            <div className="posts">
              <h3>Twitter Posts</h3>
              {posts.map((post) => (
                <p key={post.id}>
                  {post.id} {post.comments} {post.caption} {post.likes}
                </p>
              ))}
            </div>
            <button className="post-button"></button>
          </div>
          <div className="stats">
            <h3>Stats</h3>
          </div>
        </div>
      </div>

     {/* News Ticker */}
     <div className="ticker-bar">
        <div className="news-ticker">
          <div className="ticker-content">
            {travelDestinations.map((country, index) => (
              <span key={index} className="ticker-item">
                {country} &nbsp; | &nbsp;
              </span>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}
export default Dashboard;
