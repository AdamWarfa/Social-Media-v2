import { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useAuth } from "./AuthProvider";
import { User } from "../services/authFacade.ts";
import { NavLink } from "react-router-dom";

interface LoginProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

const Login = ({ setLoggedIn, setUserId }: LoginProps) => {
  const [user, setUser] = useState({ email: "", password: "" });

  const navigate = useNavigate();
  const location = useLocation();
  const auth = useAuth();

  const [err, setErr] = useState(null);

  const from = location.state?.from?.pathname || "/";

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    const user = Object.fromEntries(formData) as unknown as User;

    setErr(null);
    console.log(err);
    alert("Login: " + JSON.stringify(user));
    // return;
    auth
      .signIn(user)
      .then((res) => {
        setLoggedIn(true);
        setUserId(res.id);
        navigate(from, { replace: true });
      })
      .catch((err) => {
        setErr(err);
      });
  }
  return (
    <>
      <div className="login-wrapper" style={{ padding: "1rem" }}>
        <form className="login-form max-w-sm mx-auto" onSubmit={handleSubmit}>
          <h2 className="mt-24 mb-4 text-2xl text-center">Login</h2>
          <div className="login-form-group">
            <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Email
            </label>
            <input
              type="text"
              name="email"
              id="email"
              value={user.email}
              onChange={(e) => setUser((prev) => ({ ...prev, email: e.target.value }))}
              required
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              autoComplete="email"
            />
          </div>
          <br />
          <div className="login-form-group">
            <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
              Password
            </label>
            <input
              type="password"
              name="password"
              id="password"
              value={user.password}
              onChange={(e) => setUser((prev) => ({ ...prev, password: e.target.value }))}
              required
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-black-800 dark:border-black-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              autoComplete="current-password"
            />
          </div>
          <br />
          <div className="flex items-center h-5">
            <input
              id="remember"
              type="checkbox"
              value=""
              className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
            />
          </div>

          <button type="submit" className="login-btn">
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
};

export default Login;
