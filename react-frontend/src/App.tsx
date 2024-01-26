import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import ProfilePage from "./pages/ProfilePage";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import { useState, useEffect } from "react";
import { setPersistence, browserLocalPersistence } from "firebase/auth";
import { auth } from "./api/firebase";

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [userId, setUserId] = useState("");

  useEffect(() => {
    const setupAuthPersistence = async () => {
      try {
        await setPersistence(auth, browserLocalPersistence);
        console.log("Authentication persistence set successfully.");
      } catch (error) {
        console.error("Error setting authentication persistence:", error);
      }
    };

    setupAuthPersistence();

    const unsubscribe = auth.onAuthStateChanged((user) => {
      console.log(user);

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

    return () => {
      // Cleanup the auth state change listener when the component unmounts
      unsubscribe();
    };
  }, []);

  return (
    <Routes>
      <Route path="/" element={<HomePage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/profile/:id" element={<ProfilePage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/login" element={<Login loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/signup" element={<Signup loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="*" element={<HomePage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
    </Routes>
  );
}

export default App;
