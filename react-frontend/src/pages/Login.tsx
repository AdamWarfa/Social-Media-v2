import { signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "../api/firebase";
import { NavLink } from "react-router-dom";
import LoginAttempt from "../models/loginValues";
import Nav from "../components/Nav";

interface LoginProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function Login({ loggedIn, setLoggedIn, userId, setUserId }: LoginProps) {
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
      <Nav loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} currentPage="login" />
      <div style={{ padding: "1rem" }}>
        <h1 style={{ marginTop: "70px" }}>Login</h1>
        <form action="" method="post" onSubmit={handleLoginForm}>
          <input type="email" name="loginEmail" id="loginEmail" placeholder="Type Email..." />
          <input type="password" name="loginPassword" id="loginPassword" placeholder="Type Password..." />
          <br />
          <br />
          <button>Login</button>
          <br />
          <br />
        </form>
        <p>
          Don't have an account?
          <NavLink to={`/signup`} className="text-blue-400 hover:text-blue-500">
            Sing up here
          </NavLink>
        </p>
      </div>
    </>
  );
}
