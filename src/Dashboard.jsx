import { Menubar } from "primereact/menubar";
import ReactImg from "./assets/react.svg";
import { useEffect, useState } from "react";
import Posts from "./Posts.jsx";

function Dashboard() {
  const [posts, setPosts] = useState([]);

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
            <div className="post-button"></div>
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
            <div className="post-button"></div>
          </div>
          <div className="stats">
            <h3>Stats</h3>
          </div>
        </div>
        <div className="row-container">Footer element</div>
      </div>
    </>
  );
}

export default Dashboard;
