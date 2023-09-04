import React, { useState, useEffect } from "react";
import "./EndpointExampleModal.css";

function EndpointExampleModal({ example }) {
  const [modal, setModal] = useState(false);
  const [response, setResponse] = useState(null);
  const [loading, setLoading] = useState(false);

  const fetchData = async () => {
    try {
      setLoading(true);
      const response = await fetch(example.example);
      const data = await response.json();
      console.log(data);
      setResponse(data);
    } catch (error) {
      console.error("Error fetching data:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (modal) {
      // Make the GET request when the modal is opened
      fetchData();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [modal]);

  const toggleModal = () => {
    setModal(!modal);
  };

  if (modal) {
    document.body.classList.add("active-modal");
  } else {
    document.body.classList.remove("active-modal");
  }

  return (
    <>
      <button onClick={toggleModal} className="btn-modal">
        Show Response
      </button>

      {modal && (
        <div className="modal">
          <div onClick={toggleModal} className="overlay"></div>
          <div className="modal-content">
            {loading ? (
              <p>Loading...</p>
            ) : (
              <>
                <p>
                  <b>Verb:</b> {example.verb}
                </p>
                <p>
                  <b>Endpoint:</b> {example.example}
                </p>
                <p>
                  <b>Description:</b> {example.description}
                </p>
                <b>
                  <p>Response:</p>
                </b>
                <pre>{JSON.stringify(response, null, 2)}</pre>
              </>
            )}
            <button className="close-modal" onClick={toggleModal}>
              CLOSE
            </button>
          </div>
        </div>
      )}
    </>
  );
}

export default EndpointExampleModal;
