import NbaGamesGrid from "../components/NbaGamesGrid";
import Nav from "../components/Nav";

interface ProfilePageProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function NbaGamesPage({ loggedIn, setLoggedIn, setUserId, userId }: ProfilePageProps) {

  return (
    <>
      <Nav loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} currentPage="nbagames" />
      <h1>Profile Page</h1>
      <NbaGamesGrid />
    </>
  );
}
