import logo from "./logo.svg";
import "./App.css";
import { useState, usEffect } from "react";

function App() {
  const [data, setData] = useState(null);
  const [status, setStatus] = useState("");
  const getData = () => {
    fetch("http://localhost:8080/albums")
      .then((res) => res.json())
      .then((data) => {
        setData(data);
        console.log(data);
      })
      .catch((err) => console.log(err));
  };

  const getStatus = () => {
    fetch("http://localhost:8080/status")
      .then((res) => res.json())
      .then((data) => {
        setStatus(data.status);
        console.log(data);
      })
      .catch((err) => console.log(err));
  };

  return (
    <div className="App">
      <h1>Hello World</h1>
      <button onClick={getStatus}>Get Status</button>{" "}
      {status ? status : "Loading..."}
      <button onClick={getData}>Get Data</button>
      {data &&
        data.map((ablum) => {
          return (
            <div key={ablum.id}>
              <h3>{ablum.title}</h3>
              <p>{ablum.artist}</p>
            </div>
          );
        })}
    </div>
  );
}

export default App;
