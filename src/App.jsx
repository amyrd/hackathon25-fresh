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
      <div>
        {/* <div>
        {loggedIn ? (
          <div>
            <h1>Dashboard</h1>
<<<<<<< HEAD
=======
            {/*render stuff here */}
            <Posts />
>>>>>>> f0af1b5b1339872f2369213d1ecd96d59a5ab538
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
        <div className="body-container">
          <Header />
          <Body />
        </div>
      </div>
    </>
  );
}

export default App;
