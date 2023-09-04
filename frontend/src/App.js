import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navigation from "./components/Navigation";
import HomePage from "./components/HomePage/HomePage";
import DocsPage from "./components/DocsPage/DocsPage";
import Footer from "./components/Footer";

function App() {
  return (
    <Router>
      <div>
        <Navigation />

        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/docs" element={<DocsPage />} />
        </Routes>
        <Footer />
      </div>
    </Router>
  );
}

export default App;
