import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import ProfilePage from "./pages/ProfilePage";
import NbaGamesPage from "./pages/NbaGamesPage";
import Login from "./security/Login";
import Signup from "./pages/Signup";
import Logout from "./security/Logout";
import { useState } from "react";

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [userId, setUserId] = useState(localStorage.getItem("userId") || "");
  console.log(userId);

  return (
    <Routes>
      <Route path="/" element={<HomePage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/profile/:id" element={<ProfilePage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/nbagames" element={<NbaGamesPage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/login" element={<Login loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/signup" element={<Signup loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
      <Route path="/logout" element={<Logout />} />

      <Route path="*" element={<HomePage loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />} />
    </Routes>
  );
}

export default App;
