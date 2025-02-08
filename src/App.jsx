import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./index.css";
import Login from "./Login";
import Posts from "./Posts";
import Body from "./Body.jsx";
import Header from "./Header.jsx";

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
      <div className="body-container">
        {/* <div>
        {loggedIn ? (
          <div>
            <h1>Dashboard</h1>
          </div>
        ) : (
          <Login onLoginSuccess={handleLoginSuccess} />
        )}
      </div>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p> */}
        <Header />
        <Body />
      </div>
    </>
  );
}

export default App;
