import React from "react";
import "../../Style.css";

import heroImg from "../../images/heroimg.jpg";

function HeroBanner() {
  return (
    <div className="hero-banner">
      <div className="hero-content">
        <div className="hero-text">
          <h1>IoT data with REST API</h1>
          <p>Pseudo-real IoT data provided by a free REST API.</p>
        </div>
        <div className="hero-buttons">
          <a href="/docs">
            <button className="button-primary">Show Examples</button>
          </a>
          <a
            href="https://github.com/defilippomattia/fakeiotapi/"
            target="_blank"
            rel="noopener noreferrer"
          >
            <button className="button-secondary">GitHub</button>
          </a>
        </div>
      </div>
      <div className="hero-image">
        <img src={heroImg} alt="Hero Banner" />
      </div>
    </div>
  );
}

export default HeroBanner;
