import { useAuth } from "./AuthProvider";
import { NavLink, Link } from "react-router-dom";

export default function AuthStatus({ currentPage }: { currentPage: string }) {
  const auth = useAuth();

  const currentPageClass = "bg-gray-800 text-white rounded-md px-3 py-2 text-sm font-medium";
  const otherPageClass = "text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium";

  if (!auth.isLoggedIn()) {
    return (
      <NavLink to="/login" className={currentPage == "login" ? currentPageClass : otherPageClass}>
        Login
      </NavLink>
    );
  } else {
    return (
      <Link to="/logout" className={otherPageClass}>
        Logout (Logged in as {auth.username}){" "}
      </Link>
    );
  }
}
