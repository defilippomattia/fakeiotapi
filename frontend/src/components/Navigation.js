import React from "react";
import "../Style.css";

import iotLogo from "../images/iotlogo.png";

function Navigation() {
  return (
    <nav className="navigation">
      <div className="logo-container">
        <a href="/">
          <img src={iotLogo} alt="Logo" className="logo" />
        </a>
      </div>
      <div className="links-container">
        <ul>
          <li>
            <a href="/">Home</a>
          </li>
          <li>
            <a href="/docs">Docs</a>
          </li>
          <li>
            <a
              href="https://github.com/defilippomattia/iot-fake-rest-api/"
              target="_blank"
              rel="noopener noreferrer"
            >
              GitHub
            </a>
          </li>
        </ul>
      </div>
    </nav>
  );
}

export default Navigation;
