import React from "react";
import "./EndpointsTable.css";

import EndpointExampleModal from "./EndpointExampleModal";

function EndpointsTable() {
  var devEnv = false;
  var base_url = "";
  if (devEnv) {
    base_url = "http://localhost:8080";
  } else {
    base_url = "https://fakeiotapi.xyz";
  }

  const endpoints = [
    {
      endpoint: "/api/v1/farms",
      verb: "GET",
      description: "Get all farms",
      example: base_url + "/api/v1/farms",
    },
    {
      endpoint: "/api/v1/farms/{farm-id}",
      verb: "GET",
      description: "Get a single farm",
      example: base_url + "/api/v1/farms/1",
    },

    {
      endpoint: "/api/v1/employees",
      verb: "GET",
      description: "Get all employees",
      example: base_url + "/api/v1/employees",
    },
    {
      endpoint: "/api/v1/employees/{employee-id}",
      verb: "GET",
      description: "Get a single employee",
      example: base_url + "/api/v1/employees/1",
    },

    {
      endpoint: "/api/v1/fields",
      verb: "GET",
      description: "Get all fields",
      example: base_url + "/api/v1/fields",
    },
    {
      endpoint: "/api/v1/fields/{field-id}",
      verb: "GET",
      description: "Get a single field",
      example: base_url + "/api/v1/fields/1",
    },

    {
      endpoint: "/api/v1/sensors",
      verb: "GET",
      description: "Get all sensors",
      example: base_url + "/api/v1/sensors",
    },
    {
      endpoint: "/api/v1/sensors/{sensor-id}",
      verb: "GET",
      description: "Get a single sensor",
      example: base_url + "/api/v1/sensors/1",
    },

    {
      endpoint: "/api/v1/farms/{farm-id}/employees",
      verb: "GET",
      description: "Get all employees for a farm",
      example: base_url + "/api/v1/farms/1/employees",
    },
    {
      endpoint: "/api/v1/farms/{farm-id}/fields",
      verb: "GET",
      description: "Get all fields for a farm",
      example: base_url + "/api/v1/farms/1/fields",
    },
    {
      endpoint: "/api/v1/farms/{farm-id}/sensors",
      verb: "GET",
      description: "Get all sensors for a farm",
      example: base_url + "/api/v1/farms/1/sensors",
    },

    {
      endpoint: "/api/v1/fields/{field-id}/sensors",
      verb: "GET",
      description: "Get all sensors for a field",
      example: base_url + "/api/v1/fields/1/sensors",
    },

    {
      endpoint: "/api/v1/sensors/{sensor-id}/temperature/current",
      verb: "GET",
      description: "Get current temperature for a sensor",
      example: base_url + "/api/v1/sensors/1/temperature/current",
    },
    {
      endpoint: "/api/v1/sensors/{sensor-id}/humidity/current",
      verb: "GET",
      description: "Get current humidity for a sensor",
      example: base_url + "/api/v1/sensors/2/humidity/current",
    },
    {
      endpoint: "/api/v1/sensors/{sensor-id}/pressure/current",
      verb: "GET",
      description: "Get current pressure for a sensor",
      example: base_url + "/api/v1/sensors/3/pressure/current",
    },

    {
      endpoint: "/api/v1/sensors/{sensor-id}/temperature/history",
      verb: "GET",
      description: "Get temperature history for a sensor",
      example: base_url + "/api/v1/sensors/1/temperature/history",
    },
    {
      endpoint: "/api/v1/sensors/{sensor-id}/humidity/history",
      verb: "GET",
      description: "Get humidity history for a sensor",
      example: base_url + "/api/v1/sensors/2/humidity/history",
    },
    {
      endpoint: "/api/v1/sensors/{sensor-id}/pressure/history",
      verb: "GET",
      description: "Get pressure history for a sensor",
      example: base_url + "/api/v1/sensors/3/pressure/history",
    },
  ];

  return (
    <div className="table-container">
      <table className="table">
        <thead>
          <tr>
            <th>Description</th>
            <th>Verb</th>
            <th>Endpoint</th>
            <th>Example</th>
            <th>Response</th>
          </tr>
        </thead>
        <tbody>
          {endpoints.map((endpoint, index) => (
            <tr key={index}>
              <td>{endpoint.description}</td>
              <td>{endpoint.verb}</td>
              <td>{endpoint.endpoint}</td>
              <td>{endpoint.example}</td>
              <td>
                <EndpointExampleModal example={endpoint} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default EndpointsTable;
