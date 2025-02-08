import { Menubar } from "primereact/menubar";
import ReactImg from "./assets/react.svg";

function Dashboard() {
  return (
    <>
      <div className="navbar">
        <img src={ReactImg} style={{ width: "40px", height: "40px" }} />
        Hello
      </div>
      <div className="body-container">
        <div className="column-container">
          <div className="pane">column 1</div>
          <div className="pane">column 2</div>
          <div className="pane">column 3</div>
        </div>
        <div className="row-container">Footer element</div>
      </div>
    </>
  );
}

export default Dashboard;
