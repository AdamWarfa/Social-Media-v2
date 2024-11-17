import { NavLink, Link } from "react-router-dom";
import { AuthContextType } from "./AuthProvider";

interface AuthStatusProps {
  currentPage: string;
  auth: AuthContextType;
  currentPageClass: string;
  otherPageClass: string;
}

export default function AuthStatus({ currentPage, auth, currentPageClass, otherPageClass }: AuthStatusProps) {
  if (!auth.isLoggedIn()) {
    return (
      <NavLink to="/login" className={currentPage == "/login" ? currentPageClass : otherPageClass}>
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
