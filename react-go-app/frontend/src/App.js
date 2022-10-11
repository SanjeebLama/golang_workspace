import "./App.css";
import { useState, usEffect } from "react";
import Card from "./components/Card";
import Form from "./components/Form";
import { useEffect } from "react";

function App() {
  const [data, setData] = useState(null);
  const [status, setStatus] = useState("");

  const getData = () => {
    fetch("http://localhost:8080/quotes")
      .then((res) => res.json())
      .then((result) => {
        result.sort((a, b) => {
          return b.id - a.id;
        });

        setData(result);
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

  useEffect(() => {
    getData();
  }, []);

  return (
    <div className="App">
      <div className="flex justify-center my-5">
        <h1 className="text-2xl font-bold">React-Go-Firebase CRUD App</h1>
      </div>

      <div className=" flex justify-center h-10 w-30 ">
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold px-3 mx-4 rounded"
          onClick={getStatus}
        >
          Get Status
        </button>
        {/* <button
          className="bg-green-500 hover:bg-green-700 text-white font-bold px-3 rounded"
          onClick={getData}
        >
          Get Data
        </button> */}
      </div>

      <div className="my-auto p-7">{status ? status : ""}</div>

      <div className="flex justify-content">
        <div className="flex flex-col w-1/2">
          <Form getData={getData} />
        </div>
        <div className="flex flex-col overflow-y-auto h-96 w-1/2">
          {data &&
            data.map((item) => {
              return (
                <Card
                  id={item.id}
                  author={item.author}
                  quote={item.quote}
                  key={item.id}
                  doc_id={item.doc_id}
                  getData={getData}
                />
              );
            })}
        </div>
      </div>
    </div>
  );
}

export default App;
