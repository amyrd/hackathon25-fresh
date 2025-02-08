import { useState } from "react";
import "./index.css";
import Login from "./Login";
import Posts from "./Insta.jsx";
import Dashboard from "./Dashboard.jsx";

function App() {
  const [count, setCount] = useState(0);
  const [loggedIn, setLoggedIn] = useState(false);
  const [token, setToken] = useState(null);

  const handleLoginSuccess = (token) => {
    setLoggedIn(true);
    setToken(token);
  };

  return (
    <>
      {/* <div>
        <div>
          {loggedIn ? (
            <div>
              
            </div>
          ) : (
            <Login onLoginSuccess={handleLoginSuccess} />
          )}
        </div>
      </div> */}
      <div className="body-container">
        <Dashboard />
      </div>
    </>
  );
}

export default App;
