import { useState } from "react";
// import { getData } from "../App";
import { useMutation, queryClient } from "react-query";
import { useContext } from "react";
import { UserContext } from "../ContextAPI";
import { useEffect } from "react";
import { postQuote, updateQuote } from "../api/request";

function Form() {
  const initialState = {
    author: "",
    quote: "",
    id: "",
  };

  const [data, setData] = useState(initialState);

  const { quote } = useContext(UserContext);

  useEffect(() => {
    setData(quote);
  }, [quote]);

  const [count, setCount] = useState(6);

  const handleChange = (e) => {
    setData({
      ...data,
      [e.target.name]: e.target.value,
    });
  };

  const addID = () => {
    setCount(count + 1);
    data.id = count;
  };

  const usePostQuote = () => {
    return useMutation(postQuote, {
      onSuccess: () => {
        queryClient.invalidateQueries("quotes");
      },
    });
  };

  const { mutate } = usePostQuote();

  const useEditQuote = () => {
    return useMutation(updateQuote, {
      onSuccess: () => {
        queryClient.invalidateQueries("quotes");
      },
    });
  };

  const { mutate: editMutate } = useEditQuote();

  const handleSubmit = (e) => {
    e.preventDefault();

    if (e.target.name === "Update") {
      console.log("Inside Update");
      editMutate(data);
    }

    if (e.target.name === "Submit") {
      console.log("Inside Submit");
      addID();
      mutate(data);
    }

    // fetch getData and clear the form
    setData(initialState);
  };

  return (
    <div className="flex justify-center">
      <form>
        <h1 className="text-xl font-semi-bold font-thin text-center mb-3">
          {data?.doc_id ? "Update" : "Add"} a Quote
        </h1>
        <div className="mb-6 ">
          <label
            htmlFor="author"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
          >
            Author
          </label>
          <input
            type="text"
            id="author"
            name="author"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder=""
            value={data?.author}
            onChange={(e) => handleChange(e)}
            required
          />
        </div>
        <div className="mb-6">
          <label
            htmlFor="quote"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
          >
            Quote
          </label>
          <textarea
            type="text"
            id="quote"
            name="quote"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            value={data?.quote}
            onChange={(e) => handleChange(e)}
            required
          />
        </div>

        <button
          type="submit"
          className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          name={data?.doc_id ? "Update" : "Submit"}
          onClick={(e) => handleSubmit(e)}
        >
          {/* {isLoading ? "Loading " : "Submit"} */}
          {data?.doc_id ? "Update" : "Submit"}
        </button>
      </form>
    </div>
  );
}

export default Form;
