import React, { useEffect, useState } from "react";

function Stats() {
  const [analytics, setAnalytics] = useState([]);
  useEffect(() => {
    fetch("http://localhost:8080/instagram/analytics") // Adjust URL to match your backend endpoint
      .then((response) => response.json())
      .then((data) => setAnalytics(data))
      .catch((error) => console.error("Error fetching analytics:", error));
  }, []);

  return (
    <>
    <div id="instagram-analytics">
      <h2>Instagram Analytics</h2>
      {analytics.length > 0 ? (
        analytics.map((item, index) => (
          <div
            key={index}
            style={{
              border: "1px solid #ccc",
              margin: "10px",
              padding: "10px",
            }}
          >
            <p>
              <strong>Date:</strong> {new Date(item.date).toLocaleDateString()}
            </p>
            <p>
              <strong>Impressions:</strong> {item.impressions}
            </p>
            <p>
              <strong>Engagement:</strong> {item.engagement}
            </p>
            <p>
              <strong>Followers:</strong> {item.followers}
            </p>
            <p>
              <strong>Profile Views:</strong> {item.profile_views}
            </p>
          </div>
        ))
      ) : (
        <p>No analytics data available.</p>
      )}
    </div>
    <div id="twitter-analytics"></div>
      <h2>Twitter Analytics</h2>
    </>
  );
}

export default Stats;