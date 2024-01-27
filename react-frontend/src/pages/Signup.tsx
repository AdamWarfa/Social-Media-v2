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
      <div style={{ padding: "1rem" }}>
        <h1 style={{ marginTop: "70px" }}>Create an account</h1>
        <form action="" method="post" onSubmit={handleSignupForm}>
          <input type="email" name="signupEmail" id="signupEmail" placeholder="Type Email..." required />
          <input type="Username" name="signupUsername" id="signupUsername" placeholder="Choose Username.." required />
          <input type="password" name="signupPassword" id="signupPassword" placeholder="Choose Password..." required />
          <input type="url" name="signupAvatar" id="signupAvatar" placeholder="Enter Profile Picture URL..." required />

          <br />
          <br />
          <button className="gradient-button py-2 px-4 rounded-xl">Sign Up</button>
          <br />
          <br />
        </form>
        <p>
          Already have an account?
          <NavLink to={`/login`} className="text-blue-400 hover:text-blue-500">
            Sign in here
          </NavLink>
        </p>
      </div>
    </>
  );
}
