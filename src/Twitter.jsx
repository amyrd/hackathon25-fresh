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
      <h2>Twitter Posts</h2>
      {posts.map((post) => (
        <p key={post.id}>
          {post.id} {post.comments} {post.caption} {post.likes}{" "}
          <img
            src={post.image_url}
            style={{
              width: "100px",
              height: "auto",
            }}
          />
        </p>
      ))}
    </div>
  );
}

export default Twitter;