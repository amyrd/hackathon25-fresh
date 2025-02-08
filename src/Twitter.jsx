import React, { useEffect, useState } from "react";

function Twitter() {
  const [posts, setPosts] = useState([]);
  const [analyticsData, setAnalyticsData] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/twitter/posts") // Fetch data from Go backend
      .then((response) => response.json())
      .then((data) => setPosts(data))
      .catch((error) => console.error("Error fetching posts:", error));
  }, []);

  return (
    <div>
      <h2 className="section-header">Twitter Posts</h2>
      {posts.map((post) => (
        <div
          key={post.id}
          style={{
            border: "1px solid #ccc",
            margin: "10px",
            padding: "10px",
          }}
        >
          {" "}
          <p>{post.caption}</p>
          <img
            src={post.image_url}
            alt="Instagram post"
            style={{ width: "100%" }}
          />
          <div className="icon-bar">
            <a>💬 {post.comments}</a>
            <a>👍 {post.likes}</a>
            <a>🔁 {post.shares}</a>
          </div>
        </div>
      ))}
    </div>
  );
}

export default Twitter;
