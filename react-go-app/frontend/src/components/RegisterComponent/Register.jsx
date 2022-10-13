import { createUserWithEmailAndPassword, getAuth } from "firebase/auth";
import { useEffect, useState } from "react";
// import SocialLogins from "./SocialLogins";
import "../../firebase";
import { addDoc, collection, serverTimestamp } from "firebase/firestore";
import { db } from "../../firebase";
import { useHistory, Redirect, Router, useNavigate } from "react-router-dom";

function SignUp() {
  let history = useHistory();

  const initialData = {
    email: "",
    password: "",
    firstName: "",
    lastName: "",
  };
  const [data, setData] = useState(initialData);
  const auth = getAuth();

  // const [authData, setAuthData] = useState({
  //   email: "",
  //   // firebase_id: "",
  //   dummy_id: "",
  // });

  const [firebaseId, setFirebaseId] = useState("");

  const handleChange = (e) => {
    setData({ ...data, [e.target.name]: e.target.value });
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Inside Handle Submit : ", data);
    createUserWithEmailAndPassword(auth, data.email, data.password)
      .then((userCredential) => {
        // Signed in
        const user = userCredential.user;
        sessionStorage.setItem("Token", user.accessToken);

        // storeUser();
        // setTimeout(() => {
        history.push("/login");
        // }, 1000);
      })
      .catch((error) => {
        console.log("Error : ", error.message);
      });
  };

  // const storeUser = async () => {
  //   if (authData.email !== "" && authData.dummy_id !== "") {
  //     const email = localStorage.getItem("email");
  //     const dummy_id = localStorage.getItem("dummy_id");
  //     console.log("Email : ", email);
  //     console.log("Dummy Id : ", dummy_id);

  //     const docRef = await addDoc(collection(db, "users"), {
  //       ...authData,
  //       createdAt: serverTimestamp(),
  //     });
  //     console.log("Document written with ID: ", docRef.id);
  //   }
  // };

  // useEffect(() => {
  //   storeUser();
  // }, [authData]);

  // store data in Firestore DB named "users"

  return (
    <div>
      <div className="relative min-h-screen flex ">
        <div className="flex flex-col sm:flex-row items-center md:items-start sm:justify-center md:justify-start flex-auto min-w-0 bg-white">
          <div className="sm:w-1/2 xl:w-3/5 h-full hidden md:flex flex-auto items-end justify-end p-10 pr-20 overflow-hidden bg-purple-900 text-white bg-no-repeat bg-cover relative">
            <img
              src="https://images.unsplash.com/photo-1472841298542-80b08039a5ab?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=htmlFormat&fit=crop&w=2070&q=80"
              alt="Login"
              className="absolute inset-0 w-full h-full object-cover opacity-25"
            />
            <div className="absolute bg-gradient-to-b from-cyan-600 to-pink-500 opacity-75 inset-0 z-0"></div>
            <div className="w-full  max-w-lg z-10 mb-20">
              <div className="sm:text-4xl xl:text-5xl font-bold leading-tight mb-6">
                Reference site about Lorem Ipsum..
              </div>
              <div className="sm:text-sm xl:text-md text-gray-100 font-normal">
                {" "}
                What is Lorem Ipsum Lorem Ipsum is simply dummy text of the
                printing and typesetting industry Lorem Ipsum has been the
                industry's standard dummy text ever since the 1500s when an
                unknown printer took a galley of type and scrambled it to make a
                type specimen book it has?
              </div>
            </div>
            <ul className="circles">
              <li></li>
              <li></li>
              <li></li>
              <li></li>
              <li></li>
              <li></li>
              <li></li>
              <li></li>
              <li></li>
              <li></li>
            </ul>
          </div>
          <div className="md:flex md:items-center md:justify-center w-full sm:w-auto md:h-full w-2/5 xl:w-2/5 p-8  md:p-10 lg:p-14 sm:rounded-lg md:rounded-none bg-white">
            <div className="max-w-md w-full space-y-8">
              <div className="mb-10">
                <h3 className="font-semibold text-2xl text-gray-800">
                  REGISTRATION FORM{" "}
                </h3>
                <p className="text-gray-500">Please register your account.</p>
              </div>
              {/* <SocialLogins /> */}

              <div className="flex items-center justify-center space-x-2 my-5">
                <span className="h-px w-16 bg-gray-200"></span>
                <span className="text-gray-300 font-normal">
                  or continue with
                </span>
                <span className="h-px w-16 bg-gray-200"></span>
              </div>
              <form className="w-full max-w-lg">
                <div className="flex flex-wrap -mx-3 mb-4">
                  <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                    <label
                      className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                      htmlFor="grid-first-name"
                    >
                      First Name
                    </label>
                    <input
                      className="appearance-none block w-full bg-gray-100 text-gray-700 rounded py-3 px-4 mb-3 border border-gray-200 leading-tight focus:outline-none focus:bg-white focus:border-teal-300"
                      id="grid-first-name"
                      type="text"
                      placeholder=""
                      name="firstName"
                      onChange={(e) => {
                        handleChange(e);
                      }}
                    />
                  </div>
                  <div className="w-full md:w-1/2 px-3">
                    <label
                      className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                      htmlFor="grid-last-name"
                    >
                      Last Name
                    </label>
                    <input
                      className="appearance-none block w-full bg-gray-100 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-teal-300"
                      id="grid-last-name"
                      type="text"
                      placeholder=""
                      name="lastName"
                      onChange={(e) => {
                        handleChange(e);
                      }}
                    />
                  </div>
                </div>
                <div className="flex flex-wrap -mx-3 mb-4">
                  <div className="w-full px-3">
                    <label
                      className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                      htmlFor="grid-password"
                    >
                      Email
                    </label>
                    <input
                      className="appearance-none block w-full bg-gray-100 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-teal-300"
                      id="grid-email"
                      name="email"
                      onChange={(e) => {
                        handleChange(e);
                      }}
                      type="text"
                      placeholder=""
                    />
                  </div>
                </div>
                <div className="flex flex-wrap -mx-3 mb-4">
                  <div className="w-full px-3">
                    <label
                      className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                      htmlFor="grid-password"
                    >
                      Password
                    </label>
                    <input
                      className="appearance-none block w-full bg-gray-100 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-teal-300"
                      id="grid-password"
                      name="password"
                      onChange={(e) => {
                        handleChange(e);
                      }}
                      type="password"
                      placeholder="******************"
                    />
                  </div>
                </div>

                <div className="space-y-6 mt-6">
                  <div>
                    <button
                      type="submit"
                      onClick={handleSubmit}
                      className="w-full flex justify-center bg-purple-700 rounded-md hover:bg-purple-600 focus:outline-none focus:bg-purple-600 p-3 rounded-full text-white  shadow-lg cursor-pointer transition ease-in duration-500"
                    >
                      Sign Up
                    </button>
                  </div>
                </div>
              </form>
              <p className="mt-8 text-xs font-light text-center text-gray-700">
                {" "}
                Have an account already?{" "}
                <a
                  href="/"
                  className="font-medium text-purple-600 hover:underline"
                >
                  Login
                </a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default SignUp;
