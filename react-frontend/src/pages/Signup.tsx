import { useEffect } from "react";
import { createUserWithEmailAndPassword } from "firebase/auth";
import { auth } from "../api/firebase";
import { NavLink } from "react-router-dom";
import LoginAttempt from "../models/loginValues";
import { postUser } from "../api/createUser";

import Nav from "../components/Nav";
import AuthorType from "../models/author";

interface LoginProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function Signup({ loggedIn, setLoggedIn, userId, setUserId }: LoginProps) {
  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged((user) => {
      if (user) {
        setLoggedIn(true);
        setUserId(user.uid);
        console.log(user.email, " is logged in");
        console.log(user.uid);
      } else {
        setLoggedIn(false);
        setUserId("");
        console.log("User is logged out");
      }
    });
    return unsubscribe;
  }, []);

  const handleSignUp = (loginAttempt: LoginAttempt) => {
    createUserWithEmailAndPassword(auth, loginAttempt.email, loginAttempt.password)
      .then((response) => {
        const user = response.user;
        console.log(user.email, " is registered");

        // Add user to database
        const newUser: AuthorType = {
          id: user.uid,
          email: loginAttempt.email,
          username: loginAttempt.username,
          avatar: loginAttempt.avatar,
          password: "secret",
          followers: 0,
        };

        postUser(newUser);

        return user.uid;
      })
      .then((userId) => {
        setLoggedIn(true);
        setUserId(userId);
      })
      .catch((error) => {
        alert(error);
      });
  };

  const handleSignupForm = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;

    const loginAttempt = {
      email: form.signupEmail.value,
      username: form.signupUsername.value,
      password: form.signupPassword.value,
      avatar: form.signupAvatar.value,
    };

    handleSignUp(loginAttempt);
  };

  return (
    <>
      <Nav loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} currentPage="signup/profile" />
      <div className="p-4">
        <h1 className="mt-24 mb-4 text-2xl text-center">Create an account</h1>
        <form action="" method="post" onSubmit={handleSignupForm} className="max-w-sm mx-auto">
          <div className="mb-5">
            <label htmlFor="signupEmail" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Your email
            </label>
            <input
              type="email"
              name="signupEmail"
              id="signupEmail"
              placeholder="Type Email..."
              required
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-black-300 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
          </div>
          <div className="mb-5">
            <label htmlFor="signupUsername" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Your username
            </label>
            <input
              type="Username"
              name="signupUsername"
              id="signupUsername"
              placeholder="Choose Username.."
              required
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-black-300 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
          </div>
          <div className="mb-5">
            <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Your password
            </label>
            <input
              type="password"
              name="signupPassword"
              id="signupPassword"
              placeholder="Choose Password..."
              required
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-black-300 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
          </div>
          <div className="mb-5">
            <label htmlFor="signUpAvatar" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Your profile picture
            </label>
            <input
              type="url"
              name="signupAvatar"
              id="signupAvatar"
              placeholder="Enter Profile Picture URL..."
              required
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-black-300 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
          </div>
          <br />
          <button
            type="submit"
            className="text-white bg-blue-600 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-500 dark:hover:bg-blue-600 dark:focus:ring-blue-800"
          >
            Sign Up
          </button>
          <br />
          <br />
        </form>
        <p className="mt-10 ml-8">
          Already have an account?
          <NavLink to={`/login`} className="text-blue-400 hover:text-blue-500">
            <span> Sign in here</span>
          </NavLink>
        </p>
      </div>
    </>
  );
}
