import React, { useEffect, useState } from 'react';

function Posts() {
  const [instagramPosts, setInstagramPosts] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/instagram-posts')
      .then((res) => res.json())
      .then((data) => setInstagramPosts(data))
      .catch((err) => console.error('Error fetching posts:', err));
  }, []);

  return (
    <div>
      <h2>Instagram Posts</h2>
      <div>
        {instagramPosts.map((post) => (
          <div key={post.id} style={{ border: '1px solid #ccc', margin: '10px', padding: '10px' }}>
            <img src={post.image} alt="Instagram post" style={{ width: '100%' }} />
            <p>{post.caption}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Posts;

