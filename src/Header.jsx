import { Menubar } from "primereact/menubar";
import ReactImg from "./assets/react.svg";

function Header() {
  return (
    <>
      <div className="navbar">
        <img src={ReactImg} style={{ width: "40px", height: "40px" }} />
        Hello
      </div>
    </>
  );
}

export default Header;
