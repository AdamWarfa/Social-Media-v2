import { signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "../api/firebase";
import { NavLink } from "react-router-dom";
import LoginAttempt from "../models/loginValues";

interface LoginProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function Login({ setLoggedIn, setUserId }: LoginProps) {
  const handleLogin = (loginAttempt: LoginAttempt) => {
    signInWithEmailAndPassword(auth, loginAttempt.email, loginAttempt.password)
      .then((response) => {
        const user = response.user;

        console.log(user.email, " is logged in");

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

  const handleLoginForm = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;

    const loginAttempt = {
      email: form.loginEmail.value,
      password: form.loginPassword.value,
      avatar: "",
      username: "",
    };

    handleLogin(loginAttempt);
  };

  return (
    <>
      <div style={{ padding: "1rem" }}>
        <h1 className="mt-24 mb-4 text-2xl text-center">Login</h1>
        <form action="" method="post" onSubmit={handleLoginForm} className="max-w-sm mx-auto">
          <div className="mb-5">
            <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Your email
            </label>
            <input
              type="email"
              name="loginEmail"
              id="loginEmail"
              placeholder="Type Email..."
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
          </div>
          <div className="mb-5">
            <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Your password
            </label>
            <input
              type="password"
              name="loginPassword"
              id="loginPassword"
              placeholder="Type Password..."
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
          </div>
          <div className="flex items-start mb-5">
            <div className="flex items-center h-5">
              <input
                id="remember"
                type="checkbox"
                value=""
                className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
                required
              />
            </div>
            <label htmlFor="remember" className="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300">
              Remember me
            </label>
          </div>
          <button
            type="submit"
            className="text-white bg-blue-600 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-500 dark:hover:bg-blue-600 dark:focus:ring-blue-800"
          >
            Login
          </button>
        </form>
        <p className="mt-10 ml-8">
          Don't have an account?
          <NavLink to={`/signup`} className="text-blue-400 hover:text-blue-500">
            <span> Sign up here</span>
          </NavLink>
        </p>
      </div>
    </>
  );
}
