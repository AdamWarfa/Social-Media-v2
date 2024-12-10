import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import ProfilePage from "./pages/ProfilePage";
import Login from "./security/Login";
import Signup from "./pages/Signup";
import Logout from "./security/Logout";
import Nav from "./components/Nav";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";

type Page = "/" | "/profile/:id" | "/nbagames" | "/login" | "/signup" | "/logout" | "*";

function App() {
  const [userId, setUserId] = useState(localStorage.getItem("userId") || "");
  const [currentPage, setCurrentPage] = useState<Page>(useLocation().pathname as Page);

  const location = useLocation().pathname;

  useEffect(() => {
    if (location.includes("profile")) {
      setCurrentPage("/profile/:id");
    } else {
      setCurrentPage(location as Page);
    }
    console.log("location:", location);
  }, [location]);

  return (
    <>
      <Nav setUserId={setUserId} userId={userId} currentPage={currentPage} />
      <Routes>
        <Route path="/" element={<HomePage setUserId={setUserId} userId={userId} />} />
        <Route path="/profile/:id" element={<ProfilePage />} />
        <Route path="/login" element={<Login setUserId={setUserId} userId={userId} />} />
        <Route path="/signup" element={<Signup setUserId={setUserId} userId={userId} />} />
        <Route path="/logout" element={<Logout />} />

        <Route path="*" element={<HomePage setUserId={setUserId} userId={userId} />} />
      </Routes>
    </>
  );
}

export default App;
