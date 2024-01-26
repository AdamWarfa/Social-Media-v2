import { NavLink } from "react-router-dom";
import { auth } from "../api/firebase";
import { signOut } from "firebase/auth";

interface HeaderProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

// Opretter og eksporterer komponentet for navigationen
export default function Header({ loggedIn, setLoggedIn, setUserId, userId }: HeaderProps) {
  const handleSignOut = () => {
    signOut(auth)
      .then(() => {
        // Sign-out successful.
        setLoggedIn(false);
        setUserId("");

        console.log("Sign-out successful.");
      })
      .catch((error) => {
        // An error happened.
        console.log(error);
      });
  };
  const loginLink = loggedIn ? (
    <NavLink to="/login" onClick={handleSignOut}>
      Log-out
    </NavLink>
  ) : (
    <NavLink to="/login">Log-in</NavLink>
  );

  // Returnerer navigationen
  return (
    <>
      <nav className="flex text-center items-center justify-center space-y-10">
        <div className="grid grid-cols-4 sm:grid-cols-4 lg:grid-cols-8 gap-10 place-items-stretch  ">
          <NavLink to="/">Home</NavLink>
          {loginLink}
          {loggedIn ? <NavLink to={`/profile/${userId}`}>Profile</NavLink> : <NavLink to="/signup">Sign-up</NavLink>}
        </div>
      </nav>
    </>
  );
}
